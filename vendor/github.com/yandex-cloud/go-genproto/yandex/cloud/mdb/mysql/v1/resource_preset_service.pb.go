// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1/resource_preset_service.proto

package mysql

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type GetResourcePresetRequest struct {
	// ID of the resource preset to return.
	// To get the resource preset ID, use a [ResourcePresetService.List] request.
	ResourcePresetId     string   `protobuf:"bytes,1,opt,name=resource_preset_id,json=resourcePresetId,proto3" json:"resource_preset_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResourcePresetRequest) Reset()         { *m = GetResourcePresetRequest{} }
func (m *GetResourcePresetRequest) String() string { return proto.CompactTextString(m) }
func (*GetResourcePresetRequest) ProtoMessage()    {}
func (*GetResourcePresetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca3c22c0aa16c49e, []int{0}
}

func (m *GetResourcePresetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResourcePresetRequest.Unmarshal(m, b)
}
func (m *GetResourcePresetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResourcePresetRequest.Marshal(b, m, deterministic)
}
func (m *GetResourcePresetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourcePresetRequest.Merge(m, src)
}
func (m *GetResourcePresetRequest) XXX_Size() int {
	return xxx_messageInfo_GetResourcePresetRequest.Size(m)
}
func (m *GetResourcePresetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourcePresetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourcePresetRequest proto.InternalMessageInfo

func (m *GetResourcePresetRequest) GetResourcePresetId() string {
	if m != nil {
		return m.ResourcePresetId
	}
	return ""
}

type ListResourcePresetsRequest struct {
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListResourcePresetsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the [ListResourcePresetsResponse.next_page_token]
	// returned by a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResourcePresetsRequest) Reset()         { *m = ListResourcePresetsRequest{} }
func (m *ListResourcePresetsRequest) String() string { return proto.CompactTextString(m) }
func (*ListResourcePresetsRequest) ProtoMessage()    {}
func (*ListResourcePresetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca3c22c0aa16c49e, []int{1}
}

func (m *ListResourcePresetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResourcePresetsRequest.Unmarshal(m, b)
}
func (m *ListResourcePresetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResourcePresetsRequest.Marshal(b, m, deterministic)
}
func (m *ListResourcePresetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourcePresetsRequest.Merge(m, src)
}
func (m *ListResourcePresetsRequest) XXX_Size() int {
	return xxx_messageInfo_ListResourcePresetsRequest.Size(m)
}
func (m *ListResourcePresetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourcePresetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourcePresetsRequest proto.InternalMessageInfo

func (m *ListResourcePresetsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListResourcePresetsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListResourcePresetsResponse struct {
	// List of resource presets.
	ResourcePresets []*ResourcePreset `protobuf:"bytes,1,rep,name=resource_presets,json=resourcePresets,proto3" json:"resource_presets,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListResourcePresetsRequest.page_size], use the [next_page_token] as the value
	// for the [ListResourcePresetsRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResourcePresetsResponse) Reset()         { *m = ListResourcePresetsResponse{} }
func (m *ListResourcePresetsResponse) String() string { return proto.CompactTextString(m) }
func (*ListResourcePresetsResponse) ProtoMessage()    {}
func (*ListResourcePresetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca3c22c0aa16c49e, []int{2}
}

func (m *ListResourcePresetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResourcePresetsResponse.Unmarshal(m, b)
}
func (m *ListResourcePresetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResourcePresetsResponse.Marshal(b, m, deterministic)
}
func (m *ListResourcePresetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourcePresetsResponse.Merge(m, src)
}
func (m *ListResourcePresetsResponse) XXX_Size() int {
	return xxx_messageInfo_ListResourcePresetsResponse.Size(m)
}
func (m *ListResourcePresetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourcePresetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourcePresetsResponse proto.InternalMessageInfo

func (m *ListResourcePresetsResponse) GetResourcePresets() []*ResourcePreset {
	if m != nil {
		return m.ResourcePresets
	}
	return nil
}

func (m *ListResourcePresetsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetResourcePresetRequest)(nil), "yandex.cloud.mdb.mysql.v1.GetResourcePresetRequest")
	proto.RegisterType((*ListResourcePresetsRequest)(nil), "yandex.cloud.mdb.mysql.v1.ListResourcePresetsRequest")
	proto.RegisterType((*ListResourcePresetsResponse)(nil), "yandex.cloud.mdb.mysql.v1.ListResourcePresetsResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1/resource_preset_service.proto", fileDescriptor_ca3c22c0aa16c49e)
}

var fileDescriptor_ca3c22c0aa16c49e = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0x99, 0xa4, 0x16, 0x33, 0x22, 0x2d, 0x03, 0xc2, 0xba, 0x5a, 0x88, 0xeb, 0xc1, 0xf4,
	0x90, 0x99, 0x9d, 0x14, 0xab, 0xe0, 0x9f, 0x43, 0x3c, 0x14, 0x41, 0xa4, 0x6c, 0x7b, 0xf2, 0x12,
	0x26, 0x99, 0x97, 0x75, 0x30, 0x3b, 0xb3, 0xdd, 0x99, 0x84, 0xb6, 0xe2, 0xc5, 0xa3, 0x57, 0xbf,
	0x80, 0x17, 0xfd, 0x2a, 0xf5, 0xee, 0x37, 0x10, 0x0f, 0x7e, 0x06, 0x4f, 0x92, 0xd9, 0x14, 0xdc,
	0xad, 0x5b, 0xe3, 0x6d, 0xe1, 0x7d, 0x7f, 0xef, 0xf3, 0xbc, 0xfb, 0xcc, 0x8b, 0x1f, 0x9c, 0x08,
	0x2d, 0xe1, 0x98, 0x4d, 0xa6, 0x66, 0x26, 0x59, 0x26, 0xc7, 0x2c, 0x3b, 0xb1, 0x47, 0x53, 0x36,
	0xe7, 0xac, 0x00, 0x6b, 0x66, 0xc5, 0x04, 0x46, 0x79, 0x01, 0x16, 0xdc, 0xc8, 0x42, 0x31, 0x57,
	0x13, 0xa0, 0x79, 0x61, 0x9c, 0x21, 0x37, 0x4b, 0x90, 0x7a, 0x90, 0x66, 0x72, 0x4c, 0x3d, 0x48,
	0xe7, 0x3c, 0xbc, 0x9d, 0x1a, 0x93, 0x4e, 0x81, 0x89, 0x5c, 0x31, 0xa1, 0xb5, 0x71, 0xc2, 0x29,
	0xa3, 0x6d, 0x09, 0x86, 0x6c, 0x65, 0xc5, 0x25, 0xb0, 0x55, 0x01, 0xe6, 0x62, 0xaa, 0xa4, 0x1f,
	0x58, 0x96, 0xa3, 0x97, 0x38, 0xd8, 0x03, 0x97, 0x2c, 0xd1, 0x7d, 0x4f, 0x26, 0x70, 0x34, 0x03,
	0xeb, 0xc8, 0x00, 0x93, 0xfa, 0x16, 0x4a, 0x06, 0xa8, 0x8b, 0x7a, 0x9d, 0xe1, 0xda, 0xcf, 0x33,
	0x8e, 0x92, 0xcd, 0xa2, 0x02, 0x3e, 0x97, 0x91, 0xc1, 0xe1, 0x0b, 0x65, 0x6b, 0x03, 0xed, 0xf9,
	0xc4, 0x7b, 0xb8, 0x93, 0x8b, 0x14, 0x46, 0x56, 0x9d, 0x42, 0xd0, 0xea, 0xa2, 0x5e, 0x7b, 0x88,
	0x7f, 0x9d, 0xf1, 0xf5, 0xb8, 0xcf, 0xe3, 0x38, 0x4e, 0xae, 0x2e, 0x8a, 0x07, 0xea, 0x14, 0x48,
	0x0f, 0x63, 0xdf, 0xe8, 0xcc, 0x1b, 0xd0, 0x41, 0xdb, 0x4b, 0x76, 0x3e, 0x7c, 0xe5, 0x57, 0x1e,
	0x3f, 0xe1, 0x71, 0x9c, 0xf8, 0x29, 0x87, 0x8b, 0x5a, 0xf4, 0x05, 0xe1, 0x5b, 0x7f, 0x55, 0xb4,
	0xb9, 0xd1, 0x16, 0xc8, 0x21, 0xde, 0xac, 0x2d, 0x61, 0x03, 0xd4, 0x6d, 0xf7, 0xae, 0x0d, 0xb6,
	0x69, 0x63, 0x08, 0xb4, 0xf6, 0x43, 0x36, 0xaa, 0x7b, 0x5a, 0xc2, 0xf1, 0x86, 0x86, 0x63, 0x37,
	0xfa, 0xc3, 0x64, 0xab, 0x6e, 0xf2, 0xfa, 0xa2, 0x63, 0xff, 0xdc, 0xe8, 0xe0, 0x7b, 0x0b, 0xdf,
	0xa8, 0x8e, 0x3d, 0x28, 0x9f, 0x04, 0xf9, 0x8c, 0x70, 0x7b, 0x0f, 0x1c, 0xd9, 0xb9, 0xc4, 0x50,
	0x53, 0x48, 0xe1, 0xea, 0x5b, 0x44, 0x4f, 0xdf, 0x7f, 0xfb, 0xf1, 0xb1, 0xf5, 0x90, 0xec, 0xb2,
	0x4c, 0x68, 0x91, 0x82, 0xec, 0x5f, 0x78, 0x3b, 0xcb, 0xfd, 0xd8, 0xdb, 0x8b, 0xc1, 0xbf, 0x23,
	0x9f, 0x10, 0x5e, 0x5b, 0xfc, 0x6a, 0x72, 0xff, 0x12, 0xcd, 0xe6, 0xf4, 0xc3, 0xdd, 0xff, 0xc5,
	0xca, 0x08, 0xa3, 0x6d, 0xef, 0xfb, 0x2e, 0xb9, 0xf3, 0x4f, 0xdf, 0x43, 0x89, 0xb7, 0x2a, 0x1a,
	0x22, 0x57, 0x15, 0x9d, 0x57, 0xcf, 0x52, 0xe5, 0x5e, 0xcf, 0xc6, 0x74, 0x62, 0xb2, 0xe5, 0x29,
	0xf5, 0xcb, 0xcb, 0x48, 0x4d, 0x3f, 0x05, 0xed, 0x8f, 0xa2, 0xf9, 0xc6, 0x1e, 0xf9, 0x8f, 0xf1,
	0xba, 0x6f, 0xdb, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x81, 0xf6, 0x7d, 0xa0, 0xff, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourcePresetServiceClient is the client API for ResourcePresetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourcePresetServiceClient interface {
	// Returns the specified resource preset.
	//
	// To get the list of available resource presets, make a [List] request.
	Get(ctx context.Context, in *GetResourcePresetRequest, opts ...grpc.CallOption) (*ResourcePreset, error)
	// Retrieves the list of available resource presets.
	List(ctx context.Context, in *ListResourcePresetsRequest, opts ...grpc.CallOption) (*ListResourcePresetsResponse, error)
}

type resourcePresetServiceClient struct {
	cc *grpc.ClientConn
}

func NewResourcePresetServiceClient(cc *grpc.ClientConn) ResourcePresetServiceClient {
	return &resourcePresetServiceClient{cc}
}

func (c *resourcePresetServiceClient) Get(ctx context.Context, in *GetResourcePresetRequest, opts ...grpc.CallOption) (*ResourcePreset, error) {
	out := new(ResourcePreset)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1.ResourcePresetService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcePresetServiceClient) List(ctx context.Context, in *ListResourcePresetsRequest, opts ...grpc.CallOption) (*ListResourcePresetsResponse, error) {
	out := new(ListResourcePresetsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1.ResourcePresetService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourcePresetServiceServer is the server API for ResourcePresetService service.
type ResourcePresetServiceServer interface {
	// Returns the specified resource preset.
	//
	// To get the list of available resource presets, make a [List] request.
	Get(context.Context, *GetResourcePresetRequest) (*ResourcePreset, error)
	// Retrieves the list of available resource presets.
	List(context.Context, *ListResourcePresetsRequest) (*ListResourcePresetsResponse, error)
}

// UnimplementedResourcePresetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedResourcePresetServiceServer struct {
}

func (*UnimplementedResourcePresetServiceServer) Get(ctx context.Context, req *GetResourcePresetRequest) (*ResourcePreset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedResourcePresetServiceServer) List(ctx context.Context, req *ListResourcePresetsRequest) (*ListResourcePresetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterResourcePresetServiceServer(s *grpc.Server, srv ResourcePresetServiceServer) {
	s.RegisterService(&_ResourcePresetService_serviceDesc, srv)
}

func _ResourcePresetService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourcePresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcePresetServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1.ResourcePresetService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcePresetServiceServer).Get(ctx, req.(*GetResourcePresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourcePresetService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcePresetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcePresetServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1.ResourcePresetService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcePresetServiceServer).List(ctx, req.(*ListResourcePresetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourcePresetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.mysql.v1.ResourcePresetService",
	HandlerType: (*ResourcePresetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ResourcePresetService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ResourcePresetService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/mysql/v1/resource_preset_service.proto",
}