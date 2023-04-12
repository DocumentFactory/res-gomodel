// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileshareServiceClient is the client API for FileshareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileshareServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileshareService_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (FileshareService_DownloadFileClient, error)
	DeleteFolder(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*DeleteFolderResponse, error)
	DuplicateFile(ctx context.Context, in *DuplicateFileRequest, opts ...grpc.CallOption) (*DuplicateFileResponse, error)
}

type fileshareServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileshareServiceClient(cc grpc.ClientConnInterface) FileshareServiceClient {
	return &fileshareServiceClient{cc}
}

func (c *fileshareServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileshareService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileshareService_ServiceDesc.Streams[0], "/pnocera.fileshare.FileshareService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileshareServiceUploadFileClient{stream}
	return x, nil
}

type FileshareService_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type fileshareServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileshareServiceUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileshareServiceUploadFileClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileshareServiceClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (FileshareService_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileshareService_ServiceDesc.Streams[1], "/pnocera.fileshare.FileshareService/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileshareServiceDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileshareService_DownloadFileClient interface {
	Recv() (*DownloadFileResponse, error)
	grpc.ClientStream
}

type fileshareServiceDownloadFileClient struct {
	grpc.ClientStream
}

func (x *fileshareServiceDownloadFileClient) Recv() (*DownloadFileResponse, error) {
	m := new(DownloadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileshareServiceClient) DeleteFolder(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*DeleteFolderResponse, error) {
	out := new(DeleteFolderResponse)
	err := c.cc.Invoke(ctx, "/pnocera.fileshare.FileshareService/DeleteFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileshareServiceClient) DuplicateFile(ctx context.Context, in *DuplicateFileRequest, opts ...grpc.CallOption) (*DuplicateFileResponse, error) {
	out := new(DuplicateFileResponse)
	err := c.cc.Invoke(ctx, "/pnocera.fileshare.FileshareService/DuplicateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileshareServiceServer is the server API for FileshareService service.
// All implementations should embed UnimplementedFileshareServiceServer
// for forward compatibility
type FileshareServiceServer interface {
	UploadFile(FileshareService_UploadFileServer) error
	DownloadFile(*DownloadFileRequest, FileshareService_DownloadFileServer) error
	DeleteFolder(context.Context, *DeleteFolderRequest) (*DeleteFolderResponse, error)
	DuplicateFile(context.Context, *DuplicateFileRequest) (*DuplicateFileResponse, error)
}

// UnimplementedFileshareServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFileshareServiceServer struct {
}

func (UnimplementedFileshareServiceServer) UploadFile(FileshareService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileshareServiceServer) DownloadFile(*DownloadFileRequest, FileshareService_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileshareServiceServer) DeleteFolder(context.Context, *DeleteFolderRequest) (*DeleteFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFolder not implemented")
}
func (UnimplementedFileshareServiceServer) DuplicateFile(context.Context, *DuplicateFileRequest) (*DuplicateFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DuplicateFile not implemented")
}

// UnsafeFileshareServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileshareServiceServer will
// result in compilation errors.
type UnsafeFileshareServiceServer interface {
	mustEmbedUnimplementedFileshareServiceServer()
}

func RegisterFileshareServiceServer(s grpc.ServiceRegistrar, srv FileshareServiceServer) {
	s.RegisterService(&FileshareService_ServiceDesc, srv)
}

func _FileshareService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileshareServiceServer).UploadFile(&fileshareServiceUploadFileServer{stream})
}

type FileshareService_UploadFileServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type fileshareServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileshareServiceUploadFileServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileshareServiceUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileshareService_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileshareServiceServer).DownloadFile(m, &fileshareServiceDownloadFileServer{stream})
}

type FileshareService_DownloadFileServer interface {
	Send(*DownloadFileResponse) error
	grpc.ServerStream
}

type fileshareServiceDownloadFileServer struct {
	grpc.ServerStream
}

func (x *fileshareServiceDownloadFileServer) Send(m *DownloadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileshareService_DeleteFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileshareServiceServer).DeleteFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pnocera.fileshare.FileshareService/DeleteFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileshareServiceServer).DeleteFolder(ctx, req.(*DeleteFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileshareService_DuplicateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DuplicateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileshareServiceServer).DuplicateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pnocera.fileshare.FileshareService/DuplicateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileshareServiceServer).DuplicateFile(ctx, req.(*DuplicateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileshareService_ServiceDesc is the grpc.ServiceDesc for FileshareService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileshareService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pnocera.fileshare.FileshareService",
	HandlerType: (*FileshareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteFolder",
			Handler:    _FileshareService_DeleteFolder_Handler,
		},
		{
			MethodName: "DuplicateFile",
			Handler:    _FileshareService_DuplicateFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileshareService_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _FileshareService_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fileshare.proto",
}