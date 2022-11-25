// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fileshare.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UploadFileRequest struct {
	// Types that are valid to be assigned to Data:
	//	*UploadFileRequest_Info
	//	*UploadFileRequest_ChunkData
	Data                 isUploadFileRequest_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *UploadFileRequest) Reset()         { *m = UploadFileRequest{} }
func (m *UploadFileRequest) String() string { return proto.CompactTextString(m) }
func (*UploadFileRequest) ProtoMessage()    {}
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e4ab3f4550d917, []int{0}
}

func (m *UploadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileRequest.Unmarshal(m, b)
}
func (m *UploadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileRequest.Marshal(b, m, deterministic)
}
func (m *UploadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileRequest.Merge(m, src)
}
func (m *UploadFileRequest) XXX_Size() int {
	return xxx_messageInfo_UploadFileRequest.Size(m)
}
func (m *UploadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileRequest proto.InternalMessageInfo

type isUploadFileRequest_Data interface {
	isUploadFileRequest_Data()
}

type UploadFileRequest_Info struct {
	Info *UploadFileInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadFileRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadFileRequest_Info) isUploadFileRequest_Data() {}

func (*UploadFileRequest_ChunkData) isUploadFileRequest_Data() {}

func (m *UploadFileRequest) GetData() isUploadFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UploadFileRequest) GetInfo() *UploadFileInfo {
	if x, ok := m.GetData().(*UploadFileRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (m *UploadFileRequest) GetChunkData() []byte {
	if x, ok := m.GetData().(*UploadFileRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UploadFileRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UploadFileRequest_Info)(nil),
		(*UploadFileRequest_ChunkData)(nil),
	}
}

type UploadFileInfo struct {
	FileId               string   `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	MasterKey            []byte   `protobuf:"bytes,2,opt,name=master_key,json=masterKey,proto3" json:"master_key,omitempty"`
	RunId                string   `protobuf:"bytes,3,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileInfo) Reset()         { *m = UploadFileInfo{} }
func (m *UploadFileInfo) String() string { return proto.CompactTextString(m) }
func (*UploadFileInfo) ProtoMessage()    {}
func (*UploadFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e4ab3f4550d917, []int{1}
}

func (m *UploadFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileInfo.Unmarshal(m, b)
}
func (m *UploadFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileInfo.Marshal(b, m, deterministic)
}
func (m *UploadFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileInfo.Merge(m, src)
}
func (m *UploadFileInfo) XXX_Size() int {
	return xxx_messageInfo_UploadFileInfo.Size(m)
}
func (m *UploadFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileInfo proto.InternalMessageInfo

func (m *UploadFileInfo) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

func (m *UploadFileInfo) GetMasterKey() []byte {
	if m != nil {
		return m.MasterKey
	}
	return nil
}

func (m *UploadFileInfo) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

type UploadFileResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size                 uint32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Nonce                []byte   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileResponse) Reset()         { *m = UploadFileResponse{} }
func (m *UploadFileResponse) String() string { return proto.CompactTextString(m) }
func (*UploadFileResponse) ProtoMessage()    {}
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e4ab3f4550d917, []int{2}
}

func (m *UploadFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileResponse.Unmarshal(m, b)
}
func (m *UploadFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileResponse.Marshal(b, m, deterministic)
}
func (m *UploadFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileResponse.Merge(m, src)
}
func (m *UploadFileResponse) XXX_Size() int {
	return xxx_messageInfo_UploadFileResponse.Size(m)
}
func (m *UploadFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileResponse proto.InternalMessageInfo

func (m *UploadFileResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UploadFileResponse) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *UploadFileResponse) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

type DownloadFileRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nonce                []byte   `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	MasterKey            []byte   `protobuf:"bytes,3,opt,name=master_key,json=masterKey,proto3" json:"master_key,omitempty"`
	RunId                string   `protobuf:"bytes,4,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadFileRequest) Reset()         { *m = DownloadFileRequest{} }
func (m *DownloadFileRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadFileRequest) ProtoMessage()    {}
func (*DownloadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e4ab3f4550d917, []int{3}
}

func (m *DownloadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadFileRequest.Unmarshal(m, b)
}
func (m *DownloadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadFileRequest.Marshal(b, m, deterministic)
}
func (m *DownloadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadFileRequest.Merge(m, src)
}
func (m *DownloadFileRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadFileRequest.Size(m)
}
func (m *DownloadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadFileRequest proto.InternalMessageInfo

func (m *DownloadFileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DownloadFileRequest) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *DownloadFileRequest) GetMasterKey() []byte {
	if m != nil {
		return m.MasterKey
	}
	return nil
}

func (m *DownloadFileRequest) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

type DownloadFileResponse struct {
	ChunkData            []byte   `protobuf:"bytes,1,opt,name=chunk_data,json=chunkData,proto3" json:"chunk_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadFileResponse) Reset()         { *m = DownloadFileResponse{} }
func (m *DownloadFileResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadFileResponse) ProtoMessage()    {}
func (*DownloadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e4ab3f4550d917, []int{4}
}

func (m *DownloadFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadFileResponse.Unmarshal(m, b)
}
func (m *DownloadFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadFileResponse.Marshal(b, m, deterministic)
}
func (m *DownloadFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadFileResponse.Merge(m, src)
}
func (m *DownloadFileResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadFileResponse.Size(m)
}
func (m *DownloadFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadFileResponse proto.InternalMessageInfo

func (m *DownloadFileResponse) GetChunkData() []byte {
	if m != nil {
		return m.ChunkData
	}
	return nil
}

type DeleteFolderRequest struct {
	RunId                string   `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFolderRequest) Reset()         { *m = DeleteFolderRequest{} }
func (m *DeleteFolderRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFolderRequest) ProtoMessage()    {}
func (*DeleteFolderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e4ab3f4550d917, []int{5}
}

func (m *DeleteFolderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFolderRequest.Unmarshal(m, b)
}
func (m *DeleteFolderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFolderRequest.Marshal(b, m, deterministic)
}
func (m *DeleteFolderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFolderRequest.Merge(m, src)
}
func (m *DeleteFolderRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFolderRequest.Size(m)
}
func (m *DeleteFolderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFolderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFolderRequest proto.InternalMessageInfo

func (m *DeleteFolderRequest) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

type DeleteFolderResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFolderResponse) Reset()         { *m = DeleteFolderResponse{} }
func (m *DeleteFolderResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteFolderResponse) ProtoMessage()    {}
func (*DeleteFolderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e4ab3f4550d917, []int{6}
}

func (m *DeleteFolderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFolderResponse.Unmarshal(m, b)
}
func (m *DeleteFolderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFolderResponse.Marshal(b, m, deterministic)
}
func (m *DeleteFolderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFolderResponse.Merge(m, src)
}
func (m *DeleteFolderResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteFolderResponse.Size(m)
}
func (m *DeleteFolderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFolderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFolderResponse proto.InternalMessageInfo

func (m *DeleteFolderResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UploadFileRequest)(nil), "pnocera.fileshare.UploadFileRequest")
	proto.RegisterType((*UploadFileInfo)(nil), "pnocera.fileshare.UploadFileInfo")
	proto.RegisterType((*UploadFileResponse)(nil), "pnocera.fileshare.UploadFileResponse")
	proto.RegisterType((*DownloadFileRequest)(nil), "pnocera.fileshare.DownloadFileRequest")
	proto.RegisterType((*DownloadFileResponse)(nil), "pnocera.fileshare.DownloadFileResponse")
	proto.RegisterType((*DeleteFolderRequest)(nil), "pnocera.fileshare.DeleteFolderRequest")
	proto.RegisterType((*DeleteFolderResponse)(nil), "pnocera.fileshare.DeleteFolderResponse")
}

func init() {
	proto.RegisterFile("fileshare.proto", fileDescriptor_e6e4ab3f4550d917)
}

var fileDescriptor_e6e4ab3f4550d917 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0x3a, 0xae, 0xab, 0x0c, 0xa1, 0xd0, 0x6d, 0x10, 0x21, 0xa2, 0x10, 0xac, 0x50, 0xa2,
	0xaa, 0xb2, 0x4b, 0x11, 0x42, 0x82, 0x5b, 0x15, 0x45, 0x8d, 0x90, 0x38, 0x18, 0x71, 0xe1, 0x62,
	0x6d, 0xe3, 0x49, 0xba, 0xaa, 0xbb, 0xeb, 0xfa, 0xa3, 0xa8, 0x54, 0x3d, 0xc0, 0x15, 0x71, 0xe2,
	0xaf, 0xf0, 0x4f, 0xf8, 0x0b, 0xfc, 0x10, 0xe4, 0xb5, 0xe3, 0xd8, 0xa9, 0x69, 0x6f, 0xde, 0xd9,
	0xf7, 0xde, 0xbc, 0x79, 0xb3, 0x86, 0x7b, 0x53, 0xee, 0x63, 0x74, 0xcc, 0x42, 0xb4, 0x82, 0x50,
	0xc6, 0x92, 0x6e, 0x04, 0x42, 0x4e, 0x30, 0x64, 0x56, 0x71, 0xd1, 0x7d, 0x3c, 0x93, 0x72, 0xe6,
	0xa3, 0xcd, 0x02, 0x6e, 0x33, 0x21, 0x64, 0xcc, 0x62, 0x2e, 0x45, 0x94, 0x11, 0xcc, 0x04, 0x36,
	0x3e, 0x05, 0xbe, 0x64, 0xde, 0x88, 0xfb, 0xe8, 0xe0, 0x59, 0x82, 0x51, 0x4c, 0xdf, 0x80, 0xce,
	0xc5, 0x54, 0x76, 0x48, 0x8f, 0x0c, 0xee, 0xec, 0x3f, 0xb3, 0xae, 0x89, 0x5a, 0x0b, 0xce, 0x58,
	0x4c, 0xe5, 0xe1, 0x8a, 0xa3, 0x08, 0xf4, 0x29, 0xc0, 0xe4, 0x38, 0x11, 0x27, 0xae, 0xc7, 0x62,
	0xd6, 0xd1, 0x7a, 0x64, 0xd0, 0x3a, 0x5c, 0x71, 0x9a, 0xaa, 0x36, 0x64, 0x31, 0x3b, 0x30, 0x40,
	0x4f, 0xaf, 0x4c, 0x17, 0xd6, 0xab, 0x12, 0xf4, 0x21, 0xac, 0xa5, 0xf2, 0x2e, 0xf7, 0x54, 0xdb,
	0xa6, 0x63, 0xa4, 0xc7, 0xb1, 0x47, 0xb7, 0x00, 0x4e, 0x59, 0x14, 0x63, 0xe8, 0x9e, 0xe0, 0x45,
	0xa6, 0xe9, 0x34, 0xb3, 0xca, 0x7b, 0xbc, 0xa0, 0x0f, 0xc0, 0x08, 0x13, 0x91, 0xd2, 0x1a, 0x8a,
	0xb6, 0x1a, 0x26, 0x62, 0xec, 0x99, 0x1f, 0x80, 0x96, 0xe7, 0x8a, 0x02, 0x29, 0x22, 0xa4, 0xeb,
	0xa0, 0x15, 0xfa, 0x1a, 0xf7, 0x28, 0x05, 0x3d, 0xe2, 0x5f, 0x51, 0xa9, 0xde, 0x75, 0xd4, 0x37,
	0x6d, 0xc3, 0xaa, 0x90, 0x62, 0x82, 0x4a, 0xaf, 0xe5, 0x64, 0x07, 0xf3, 0x0c, 0x36, 0x87, 0xf2,
	0x8b, 0x58, 0x4e, 0x6a, 0x59, 0xb0, 0x20, 0x6b, 0x25, 0xf2, 0xd2, 0x08, 0x8d, 0xff, 0x8f, 0xa0,
	0x97, 0x47, 0x78, 0x0d, 0xed, 0x6a, 0xcb, 0x7c, 0x88, 0xad, 0x4a, 0xc8, 0x24, 0x53, 0x2b, 0x22,
	0x36, 0x77, 0x61, 0x73, 0x88, 0x3e, 0xc6, 0x38, 0x92, 0xbe, 0x87, 0xe1, 0xdc, 0xe9, 0xa2, 0x09,
	0x29, 0x37, 0xd9, 0x83, 0x76, 0x15, 0x9d, 0x37, 0xe9, 0xc0, 0xda, 0x29, 0x46, 0x11, 0x9b, 0x61,
	0x8e, 0x9f, 0x1f, 0xf7, 0x7f, 0x37, 0xe0, 0xfe, 0x68, 0xfe, 0x10, 0x3e, 0x62, 0x78, 0xce, 0x27,
	0x48, 0xbf, 0x11, 0x80, 0x45, 0xde, 0xb4, 0x7f, 0xe3, 0x93, 0xc9, 0x2d, 0x75, 0x9f, 0xdf, 0x82,
	0xca, 0xac, 0x98, 0xfd, 0xef, 0x7f, 0xfe, 0xfe, 0xd2, 0x9e, 0x98, 0x8f, 0xec, 0xf3, 0x97, 0x76,
	0x81, 0xb4, 0x13, 0x85, 0x74, 0xd3, 0xc2, 0x5b, 0xb2, 0x33, 0x20, 0xf4, 0x07, 0x81, 0x56, 0x39,
	0x30, 0xba, 0x5d, 0xa3, 0x5f, 0xb3, 0xc4, 0xee, 0x8b, 0x5b, 0x71, 0xb9, 0x93, 0x81, 0x72, 0x62,
	0xd2, 0x5e, 0xd5, 0x89, 0x97, 0x63, 0x95, 0x17, 0xfb, 0x92, 0x7b, 0x57, 0x7b, 0x84, 0xfe, 0x4c,
	0xdd, 0x94, 0x92, 0xad, 0x77, 0x73, 0x7d, 0x51, 0xf5, 0x6e, 0x6a, 0x56, 0x64, 0xee, 0x2a, 0x37,
	0xdb, 0x3b, 0xfd, 0x25, 0x37, 0x0a, 0xeb, 0x4e, 0x15, 0xd8, 0xbe, 0xcc, 0x96, 0x7e, 0x75, 0x60,
	0x7c, 0xd6, 0xad, 0x77, 0xc1, 0xd1, 0x91, 0xa1, 0xfe, 0xfb, 0x57, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0xbe, 0x09, 0x3f, 0x3b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileshareServiceClient is the client API for FileshareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileshareServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileshareService_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (FileshareService_DownloadFileClient, error)
	DeleteFolder(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*DeleteFolderResponse, error)
}

type fileshareServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileshareServiceClient(cc grpc.ClientConnInterface) FileshareServiceClient {
	return &fileshareServiceClient{cc}
}

func (c *fileshareServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileshareService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileshareService_serviceDesc.Streams[0], "/pnocera.fileshare.FileshareService/UploadFile", opts...)
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
	stream, err := c.cc.NewStream(ctx, &_FileshareService_serviceDesc.Streams[1], "/pnocera.fileshare.FileshareService/DownloadFile", opts...)
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

// FileshareServiceServer is the server API for FileshareService service.
type FileshareServiceServer interface {
	UploadFile(FileshareService_UploadFileServer) error
	DownloadFile(*DownloadFileRequest, FileshareService_DownloadFileServer) error
	DeleteFolder(context.Context, *DeleteFolderRequest) (*DeleteFolderResponse, error)
}

// UnimplementedFileshareServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileshareServiceServer struct {
}

func (*UnimplementedFileshareServiceServer) UploadFile(srv FileshareService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedFileshareServiceServer) DownloadFile(req *DownloadFileRequest, srv FileshareService_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (*UnimplementedFileshareServiceServer) DeleteFolder(ctx context.Context, req *DeleteFolderRequest) (*DeleteFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFolder not implemented")
}

func RegisterFileshareServiceServer(s *grpc.Server, srv FileshareServiceServer) {
	s.RegisterService(&_FileshareService_serviceDesc, srv)
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

var _FileshareService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pnocera.fileshare.FileshareService",
	HandlerType: (*FileshareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteFolder",
			Handler:    _FileshareService_DeleteFolder_Handler,
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
