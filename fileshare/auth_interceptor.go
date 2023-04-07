package fileshare

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	authClient  *AuthClient
	authMethods map[string]bool
	accessToken string
}

func NewAuthInterceptor(authClient *AuthClient, authMethods map[string]bool, refreshDuration time.Duration) (*AuthInterceptor, error) {

	interceptor := &AuthInterceptor{authClient: authClient, authMethods: authMethods}
	err := interceptor.scheduleRefreshToken(refreshDuration)

	if err != nil {
		return nil, err
	}

	return interceptor, nil

}

func (i *AuthInterceptor) scheduleRefreshToken(refreshDuration time.Duration) error {

	err := i.RefreshToken()

	if err != nil {
		return err
	}

	go func() {
		wait := refreshDuration
		for {
			time.Sleep(wait)
			err := i.RefreshToken()
			if err != nil {
				wait = time.Second
			} else {
				wait = refreshDuration
			}
		}
	}()

	return nil

}

func (i *AuthInterceptor) RefreshToken() error {
	accessToken, err := i.authClient.Login()

	if err != nil {
		return err
	}

	i.accessToken = accessToken

	return nil
}

func (i *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Printf("Unary interceptor: %s", method)
		if i.authMethods[method] {
			return invoker(i.attachToken(ctx), method, req, reply, cc, opts...)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (i *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", i.accessToken, "dapr-app-id", "fileshare", "dapr-stream", "true")
}

func (i *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		log.Printf("Stream interceptor: %s", method)
		if i.authMethods[method] {
			return streamer(i.attachToken(ctx), desc, cc, method, opts...)
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}
