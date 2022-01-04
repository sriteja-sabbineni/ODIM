// Code generated by protoc-gen-go. DO NOT EDIT.
// source: update/update.proto

package update

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UpdateRequest struct {
	SessionToken         string   `protobuf:"bytes,1,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	URL                  string   `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	RequestBody          []byte   `protobuf:"bytes,3,opt,name=RequestBody,proto3" json:"RequestBody,omitempty"`
	ResourceID           string   `protobuf:"bytes,4,opt,name=resourceID,proto3" json:"resourceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5b0e1a82401048, []int{0}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *UpdateRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *UpdateRequest) GetRequestBody() []byte {
	if m != nil {
		return m.RequestBody
	}
	return nil
}

func (m *UpdateRequest) GetResourceID() string {
	if m != nil {
		return m.ResourceID
	}
	return ""
}

type UpdateResponse struct {
	StatusCode           int32             `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage        string            `protobuf:"bytes,2,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	Header               map[string]string `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5b0e1a82401048, []int{1}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *UpdateResponse) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

func (m *UpdateResponse) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *UpdateResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterMapType((map[string]string)(nil), "UpdateResponse.HeaderEntry")
}

func init() { proto.RegisterFile("update/update.proto", fileDescriptor_fc5b0e1a82401048) }

var fileDescriptor_fc5b0e1a82401048 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xd1, 0x4e, 0xea, 0x40,
	0x10, 0xbd, 0xa5, 0x40, 0x72, 0xa7, 0xc0, 0x25, 0x7b, 0x79, 0x68, 0xb8, 0x09, 0x69, 0x9a, 0xfb,
	0xc0, 0x53, 0x35, 0x10, 0x13, 0xf1, 0x51, 0x54, 0x24, 0xd1, 0x97, 0x56, 0x3e, 0xa0, 0xd0, 0x51,
	0x1b, 0xca, 0x6e, 0xdd, 0xdd, 0x62, 0xfa, 0x05, 0x7e, 0x89, 0x3f, 0xe5, 0xd7, 0x98, 0x6e, 0x4b,
	0xd2, 0xaa, 0x0f, 0xd5, 0xa7, 0x9d, 0x39, 0x33, 0x67, 0xcf, 0x99, 0xd9, 0x2c, 0xfc, 0x4d, 0xe2,
	0xc0, 0x97, 0x78, 0x94, 0x1f, 0x4e, 0xcc, 0x99, 0x64, 0xf6, 0x8b, 0x06, 0xdd, 0x95, 0x02, 0x5c,
	0x7c, 0x4a, 0x50, 0x48, 0x62, 0x43, 0xc7, 0x43, 0x21, 0x42, 0x46, 0xef, 0xd8, 0x16, 0xa9, 0xa9,
	0x59, 0xda, 0xf8, 0xb7, 0x5b, 0xc1, 0x48, 0x1f, 0xf4, 0x95, 0x7b, 0x63, 0x36, 0x54, 0x29, 0x0b,
	0x89, 0x05, 0x46, 0x71, 0xc1, 0x39, 0x0b, 0x52, 0x53, 0xb7, 0xb4, 0x71, 0xc7, 0x2d, 0x43, 0x64,
	0x04, 0xc0, 0x51, 0xb0, 0x84, 0x6f, 0x70, 0x79, 0x61, 0x36, 0x15, 0xb5, 0x84, 0xd8, 0x6f, 0x1a,
	0xf4, 0x0e, 0x4e, 0x44, 0xcc, 0xa8, 0xc0, 0x8c, 0x22, 0xa4, 0x2f, 0x13, 0x31, 0x67, 0x01, 0x2a,
	0x23, 0x2d, 0xb7, 0x84, 0x90, 0xff, 0xd0, 0xcd, 0xb3, 0x5b, 0x14, 0xc2, 0x7f, 0xc0, 0xc2, 0x50,
	0x15, 0x24, 0x53, 0x68, 0x3f, 0xa2, 0x1f, 0x20, 0x37, 0x75, 0x4b, 0x1f, 0x1b, 0x93, 0x7f, 0x4e,
	0x55, 0xc6, 0xb9, 0x56, 0xd5, 0x4b, 0x2a, 0x79, 0xea, 0x16, 0xad, 0x84, 0x40, 0x73, 0x9d, 0x0d,
	0xd2, 0x54, 0x83, 0xa8, 0x78, 0x38, 0x03, 0xa3, 0xd4, 0x9a, 0x2d, 0x61, 0x8b, 0x69, 0xb1, 0x9f,
	0x2c, 0x24, 0x03, 0x68, 0xed, 0xfd, 0x28, 0x39, 0xf8, 0xc8, 0x93, 0xb3, 0xc6, 0xa9, 0x36, 0x79,
	0xd5, 0xa1, 0x9d, 0xab, 0x92, 0x13, 0xe8, 0x2f, 0x50, 0xe6, 0x89, 0x87, 0x7c, 0x1f, 0x6e, 0x90,
	0xf4, 0x9c, 0xca, 0x1b, 0x0c, 0xff, 0x7c, 0xb0, 0x68, 0xff, 0x22, 0x33, 0x18, 0x2c, 0x50, 0x5e,
	0x85, 0x7c, 0xf7, 0xec, 0x73, 0x5c, 0xd2, 0x3d, 0x52, 0xc9, 0x78, 0x5a, 0x87, 0x3a, 0x87, 0xd1,
	0x57, 0xd4, 0x39, 0x8b, 0x22, 0xdc, 0xc8, 0x90, 0xd1, 0xfa, 0xfa, 0x1e, 0xbb, 0x97, 0x3f, 0xd4,
	0xff, 0x44, 0xfd, 0x9e, 0xfe, 0x04, 0xba, 0x5e, 0xb8, 0xc3, 0x38, 0xc2, 0x62, 0x8f, 0x35, 0x38,
	0xc7, 0x60, 0x78, 0xd2, 0xe7, 0xb2, 0x36, 0x63, 0xdd, 0x56, 0xbf, 0x62, 0xfa, 0x1e, 0x00, 0x00,
	0xff, 0xff, 0xd4, 0x97, 0x96, 0xac, 0x2c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UpdateClient is the client API for Update service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpdateClient interface {
	GetUpdateService(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetFirmwareInventory(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetFirmwareInventoryCollection(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetSoftwareInventory(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetSoftwareInventoryCollection(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	SimepleUpdate(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	StartUpdate(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type updateClient struct {
	cc *grpc.ClientConn
}

func NewUpdateClient(cc *grpc.ClientConn) UpdateClient {
	return &updateClient{cc}
}

func (c *updateClient) GetUpdateService(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/Update/GetUpdateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) GetFirmwareInventory(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/Update/GetFirmwareInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) GetFirmwareInventoryCollection(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/Update/GetFirmwareInventoryCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) GetSoftwareInventory(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/Update/GetSoftwareInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) GetSoftwareInventoryCollection(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/Update/GetSoftwareInventoryCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) SimepleUpdate(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/Update/SimepleUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) StartUpdate(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/Update/StartUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateServer is the server API for Update service.
type UpdateServer interface {
	GetUpdateService(context.Context, *UpdateRequest) (*UpdateResponse, error)
	GetFirmwareInventory(context.Context, *UpdateRequest) (*UpdateResponse, error)
	GetFirmwareInventoryCollection(context.Context, *UpdateRequest) (*UpdateResponse, error)
	GetSoftwareInventory(context.Context, *UpdateRequest) (*UpdateResponse, error)
	GetSoftwareInventoryCollection(context.Context, *UpdateRequest) (*UpdateResponse, error)
	SimepleUpdate(context.Context, *UpdateRequest) (*UpdateResponse, error)
	StartUpdate(context.Context, *UpdateRequest) (*UpdateResponse, error)
}

// UnimplementedUpdateServer can be embedded to have forward compatible implementations.
type UnimplementedUpdateServer struct {
}

func (*UnimplementedUpdateServer) GetUpdateService(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpdateService not implemented")
}
func (*UnimplementedUpdateServer) GetFirmwareInventory(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirmwareInventory not implemented")
}
func (*UnimplementedUpdateServer) GetFirmwareInventoryCollection(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirmwareInventoryCollection not implemented")
}
func (*UnimplementedUpdateServer) GetSoftwareInventory(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSoftwareInventory not implemented")
}
func (*UnimplementedUpdateServer) GetSoftwareInventoryCollection(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSoftwareInventoryCollection not implemented")
}
func (*UnimplementedUpdateServer) SimepleUpdate(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimepleUpdate not implemented")
}
func (*UnimplementedUpdateServer) StartUpdate(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartUpdate not implemented")
}

func RegisterUpdateServer(s *grpc.Server, srv UpdateServer) {
	s.RegisterService(&_Update_serviceDesc, srv)
}

func _Update_GetUpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).GetUpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Update/GetUpdateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).GetUpdateService(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_GetFirmwareInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).GetFirmwareInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Update/GetFirmwareInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).GetFirmwareInventory(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_GetFirmwareInventoryCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).GetFirmwareInventoryCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Update/GetFirmwareInventoryCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).GetFirmwareInventoryCollection(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_GetSoftwareInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).GetSoftwareInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Update/GetSoftwareInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).GetSoftwareInventory(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_GetSoftwareInventoryCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).GetSoftwareInventoryCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Update/GetSoftwareInventoryCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).GetSoftwareInventoryCollection(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_SimepleUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).SimepleUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Update/SimepleUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).SimepleUpdate(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_StartUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).StartUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Update/StartUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).StartUpdate(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Update_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Update",
	HandlerType: (*UpdateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpdateService",
			Handler:    _Update_GetUpdateService_Handler,
		},
		{
			MethodName: "GetFirmwareInventory",
			Handler:    _Update_GetFirmwareInventory_Handler,
		},
		{
			MethodName: "GetFirmwareInventoryCollection",
			Handler:    _Update_GetFirmwareInventoryCollection_Handler,
		},
		{
			MethodName: "GetSoftwareInventory",
			Handler:    _Update_GetSoftwareInventory_Handler,
		},
		{
			MethodName: "GetSoftwareInventoryCollection",
			Handler:    _Update_GetSoftwareInventoryCollection_Handler,
		},
		{
			MethodName: "SimepleUpdate",
			Handler:    _Update_SimepleUpdate_Handler,
		},
		{
			MethodName: "StartUpdate",
			Handler:    _Update_StartUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "update/update.proto",
}
