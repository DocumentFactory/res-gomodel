package fileshare

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type DaprInterceptor struct {
}

func NewDaprInterceptor() (*DaprInterceptor, error) {

	interceptor := &DaprInterceptor{}

	return interceptor, nil

}

func (i *DaprInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		return invoker(i.attachHeaders(ctx), method, req, reply, cc, opts...)

	}
}

func (i *DaprInterceptor) attachHeaders(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "fileshare", "dapr-stream", "true")
}

func (i *DaprInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

		return streamer(i.attachHeaders(ctx), desc, cc, method, opts...)

	}
}
