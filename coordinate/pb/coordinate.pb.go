// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coordinate.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Coordinate struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ItemIds              []int32  `protobuf:"varint,3,rep,packed,name=itemIds,proto3" json:"itemIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinate) Reset()         { *m = Coordinate{} }
func (m *Coordinate) String() string { return proto.CompactTextString(m) }
func (*Coordinate) ProtoMessage()    {}
func (*Coordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinate_9946b7a78021c270, []int{0}
}
func (m *Coordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate.Unmarshal(m, b)
}
func (m *Coordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate.Marshal(b, m, deterministic)
}
func (dst *Coordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate.Merge(dst, src)
}
func (m *Coordinate) XXX_Size() int {
	return xxx_messageInfo_Coordinate.Size(m)
}
func (m *Coordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate proto.InternalMessageInfo

func (m *Coordinate) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Coordinate) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Coordinate) GetItemIds() []int32 {
	if m != nil {
		return m.ItemIds
	}
	return nil
}

type GetCoordinateRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCoordinateRequest) Reset()         { *m = GetCoordinateRequest{} }
func (m *GetCoordinateRequest) String() string { return proto.CompactTextString(m) }
func (*GetCoordinateRequest) ProtoMessage()    {}
func (*GetCoordinateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinate_9946b7a78021c270, []int{1}
}
func (m *GetCoordinateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCoordinateRequest.Unmarshal(m, b)
}
func (m *GetCoordinateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCoordinateRequest.Marshal(b, m, deterministic)
}
func (dst *GetCoordinateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCoordinateRequest.Merge(dst, src)
}
func (m *GetCoordinateRequest) XXX_Size() int {
	return xxx_messageInfo_GetCoordinateRequest.Size(m)
}
func (m *GetCoordinateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCoordinateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCoordinateRequest proto.InternalMessageInfo

func (m *GetCoordinateRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetCoordinateResponse struct {
	Coordinate           *Coordinate `protobuf:"bytes,1,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetCoordinateResponse) Reset()         { *m = GetCoordinateResponse{} }
func (m *GetCoordinateResponse) String() string { return proto.CompactTextString(m) }
func (*GetCoordinateResponse) ProtoMessage()    {}
func (*GetCoordinateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinate_9946b7a78021c270, []int{2}
}
func (m *GetCoordinateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCoordinateResponse.Unmarshal(m, b)
}
func (m *GetCoordinateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCoordinateResponse.Marshal(b, m, deterministic)
}
func (dst *GetCoordinateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCoordinateResponse.Merge(dst, src)
}
func (m *GetCoordinateResponse) XXX_Size() int {
	return xxx_messageInfo_GetCoordinateResponse.Size(m)
}
func (m *GetCoordinateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCoordinateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCoordinateResponse proto.InternalMessageInfo

func (m *GetCoordinateResponse) GetCoordinate() *Coordinate {
	if m != nil {
		return m.Coordinate
	}
	return nil
}

type GetCoordinatesByUserRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCoordinatesByUserRequest) Reset()         { *m = GetCoordinatesByUserRequest{} }
func (m *GetCoordinatesByUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetCoordinatesByUserRequest) ProtoMessage()    {}
func (*GetCoordinatesByUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinate_9946b7a78021c270, []int{3}
}
func (m *GetCoordinatesByUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCoordinatesByUserRequest.Unmarshal(m, b)
}
func (m *GetCoordinatesByUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCoordinatesByUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetCoordinatesByUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCoordinatesByUserRequest.Merge(dst, src)
}
func (m *GetCoordinatesByUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetCoordinatesByUserRequest.Size(m)
}
func (m *GetCoordinatesByUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCoordinatesByUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCoordinatesByUserRequest proto.InternalMessageInfo

func (m *GetCoordinatesByUserRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetCoordinatesByUserResponse struct {
	Coordinates          []*Coordinate `protobuf:"bytes,1,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetCoordinatesByUserResponse) Reset()         { *m = GetCoordinatesByUserResponse{} }
func (m *GetCoordinatesByUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetCoordinatesByUserResponse) ProtoMessage()    {}
func (*GetCoordinatesByUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinate_9946b7a78021c270, []int{4}
}
func (m *GetCoordinatesByUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCoordinatesByUserResponse.Unmarshal(m, b)
}
func (m *GetCoordinatesByUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCoordinatesByUserResponse.Marshal(b, m, deterministic)
}
func (dst *GetCoordinatesByUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCoordinatesByUserResponse.Merge(dst, src)
}
func (m *GetCoordinatesByUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetCoordinatesByUserResponse.Size(m)
}
func (m *GetCoordinatesByUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCoordinatesByUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCoordinatesByUserResponse proto.InternalMessageInfo

func (m *GetCoordinatesByUserResponse) GetCoordinates() []*Coordinate {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

type PostCoordinateRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ItemIds              []int32  `protobuf:"varint,2,rep,packed,name=itemIds,proto3" json:"itemIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostCoordinateRequest) Reset()         { *m = PostCoordinateRequest{} }
func (m *PostCoordinateRequest) String() string { return proto.CompactTextString(m) }
func (*PostCoordinateRequest) ProtoMessage()    {}
func (*PostCoordinateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinate_9946b7a78021c270, []int{5}
}
func (m *PostCoordinateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostCoordinateRequest.Unmarshal(m, b)
}
func (m *PostCoordinateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostCoordinateRequest.Marshal(b, m, deterministic)
}
func (dst *PostCoordinateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostCoordinateRequest.Merge(dst, src)
}
func (m *PostCoordinateRequest) XXX_Size() int {
	return xxx_messageInfo_PostCoordinateRequest.Size(m)
}
func (m *PostCoordinateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostCoordinateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostCoordinateRequest proto.InternalMessageInfo

func (m *PostCoordinateRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PostCoordinateRequest) GetItemIds() []int32 {
	if m != nil {
		return m.ItemIds
	}
	return nil
}

type PostCoordinateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostCoordinateResponse) Reset()         { *m = PostCoordinateResponse{} }
func (m *PostCoordinateResponse) String() string { return proto.CompactTextString(m) }
func (*PostCoordinateResponse) ProtoMessage()    {}
func (*PostCoordinateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinate_9946b7a78021c270, []int{6}
}
func (m *PostCoordinateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostCoordinateResponse.Unmarshal(m, b)
}
func (m *PostCoordinateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostCoordinateResponse.Marshal(b, m, deterministic)
}
func (dst *PostCoordinateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostCoordinateResponse.Merge(dst, src)
}
func (m *PostCoordinateResponse) XXX_Size() int {
	return xxx_messageInfo_PostCoordinateResponse.Size(m)
}
func (m *PostCoordinateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostCoordinateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostCoordinateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Coordinate)(nil), "pb.Coordinate")
	proto.RegisterType((*GetCoordinateRequest)(nil), "pb.getCoordinateRequest")
	proto.RegisterType((*GetCoordinateResponse)(nil), "pb.getCoordinateResponse")
	proto.RegisterType((*GetCoordinatesByUserRequest)(nil), "pb.getCoordinatesByUserRequest")
	proto.RegisterType((*GetCoordinatesByUserResponse)(nil), "pb.getCoordinatesByUserResponse")
	proto.RegisterType((*PostCoordinateRequest)(nil), "pb.postCoordinateRequest")
	proto.RegisterType((*PostCoordinateResponse)(nil), "pb.postCoordinateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoordinateServiceClient is the client API for CoordinateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoordinateServiceClient interface {
	GetCoordinate(ctx context.Context, in *GetCoordinateRequest, opts ...grpc.CallOption) (*GetCoordinateResponse, error)
	GetCoordinatesByUser(ctx context.Context, in *GetCoordinatesByUserRequest, opts ...grpc.CallOption) (*GetCoordinatesByUserResponse, error)
	PostCoordinate(ctx context.Context, in *PostCoordinateRequest, opts ...grpc.CallOption) (*PostCoordinateResponse, error)
}

type coordinateServiceClient struct {
	cc *grpc.ClientConn
}

func NewCoordinateServiceClient(cc *grpc.ClientConn) CoordinateServiceClient {
	return &coordinateServiceClient{cc}
}

func (c *coordinateServiceClient) GetCoordinate(ctx context.Context, in *GetCoordinateRequest, opts ...grpc.CallOption) (*GetCoordinateResponse, error) {
	out := new(GetCoordinateResponse)
	err := c.cc.Invoke(ctx, "/pb.CoordinateService/getCoordinate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinateServiceClient) GetCoordinatesByUser(ctx context.Context, in *GetCoordinatesByUserRequest, opts ...grpc.CallOption) (*GetCoordinatesByUserResponse, error) {
	out := new(GetCoordinatesByUserResponse)
	err := c.cc.Invoke(ctx, "/pb.CoordinateService/getCoordinatesByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinateServiceClient) PostCoordinate(ctx context.Context, in *PostCoordinateRequest, opts ...grpc.CallOption) (*PostCoordinateResponse, error) {
	out := new(PostCoordinateResponse)
	err := c.cc.Invoke(ctx, "/pb.CoordinateService/postCoordinate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinateServiceServer is the server API for CoordinateService service.
type CoordinateServiceServer interface {
	GetCoordinate(context.Context, *GetCoordinateRequest) (*GetCoordinateResponse, error)
	GetCoordinatesByUser(context.Context, *GetCoordinatesByUserRequest) (*GetCoordinatesByUserResponse, error)
	PostCoordinate(context.Context, *PostCoordinateRequest) (*PostCoordinateResponse, error)
}

func RegisterCoordinateServiceServer(s *grpc.Server, srv CoordinateServiceServer) {
	s.RegisterService(&_CoordinateService_serviceDesc, srv)
}

func _CoordinateService_GetCoordinate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoordinateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinateServiceServer).GetCoordinate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CoordinateService/GetCoordinate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinateServiceServer).GetCoordinate(ctx, req.(*GetCoordinateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoordinateService_GetCoordinatesByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoordinatesByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinateServiceServer).GetCoordinatesByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CoordinateService/GetCoordinatesByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinateServiceServer).GetCoordinatesByUser(ctx, req.(*GetCoordinatesByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoordinateService_PostCoordinate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCoordinateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinateServiceServer).PostCoordinate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CoordinateService/PostCoordinate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinateServiceServer).PostCoordinate(ctx, req.(*PostCoordinateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoordinateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CoordinateService",
	HandlerType: (*CoordinateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCoordinate",
			Handler:    _CoordinateService_GetCoordinate_Handler,
		},
		{
			MethodName: "getCoordinatesByUser",
			Handler:    _CoordinateService_GetCoordinatesByUser_Handler,
		},
		{
			MethodName: "postCoordinate",
			Handler:    _CoordinateService_PostCoordinate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coordinate.proto",
}

func init() { proto.RegisterFile("coordinate.proto", fileDescriptor_coordinate_9946b7a78021c270) }

var fileDescriptor_coordinate_9946b7a78021c270 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0x1b, 0x5a, 0x61, 0x8a, 0x41, 0x07, 0x5b, 0xd6, 0x28, 0x18, 0xf6, 0x20, 0x3d, 0x05,
	0xa9, 0xf8, 0x03, 0x0a, 0x4a, 0x2e, 0x22, 0x11, 0x4f, 0x9e, 0x4c, 0x33, 0x48, 0x0e, 0x76, 0x63,
	0x76, 0x2b, 0x78, 0xf6, 0xc7, 0xc5, 0x34, 0xe9, 0xee, 0xa6, 0xdb, 0xe3, 0xce, 0xbc, 0x79, 0xf3,
	0xde, 0x9b, 0x85, 0xe3, 0xa5, 0x94, 0x4d, 0x59, 0xad, 0xde, 0x35, 0xa5, 0x75, 0x23, 0xb5, 0x44,
	0x56, 0x17, 0xe2, 0x09, 0xe0, 0x7e, 0x5b, 0xc7, 0x08, 0x58, 0x55, 0xf2, 0x20, 0x09, 0xe6, 0xa3,
	0x9c, 0x55, 0x25, 0xce, 0x60, 0xbc, 0x56, 0xd4, 0x64, 0x25, 0x67, 0x6d, 0xad, 0x7b, 0x21, 0x87,
	0xc3, 0x4a, 0xd3, 0x67, 0x56, 0x2a, 0x1e, 0x26, 0xe1, 0x7c, 0x94, 0xf7, 0x4f, 0x71, 0x05, 0xa7,
	0x1f, 0xa4, 0x0d, 0x65, 0x4e, 0x5f, 0x6b, 0x52, 0x7a, 0xc8, 0x2c, 0x1e, 0x61, 0x3a, 0xc0, 0xa9,
	0x5a, 0xae, 0x14, 0x61, 0x0a, 0x60, 0x84, 0xb6, 0x03, 0x93, 0x45, 0x94, 0xd6, 0x45, 0x6a, 0x61,
	0x2d, 0x84, 0xb8, 0x85, 0x73, 0x87, 0x48, 0xdd, 0xfd, 0xbc, 0x2a, 0x6a, 0xfa, 0xbd, 0xc6, 0x41,
	0x60, 0x3b, 0x10, 0xcf, 0x70, 0xe1, 0x1f, 0xeb, 0x64, 0x5c, 0xc3, 0xc4, 0x2c, 0x51, 0x3c, 0x48,
	0x42, 0x8f, 0x0e, 0x1b, 0x22, 0x32, 0x98, 0xd6, 0x52, 0x79, 0xac, 0xef, 0x91, 0x60, 0x87, 0xc8,
	0xdc, 0x10, 0x39, 0xcc, 0x86, 0x54, 0x1b, 0x59, 0x8b, 0x5f, 0x06, 0x27, 0xa6, 0xfc, 0x42, 0xcd,
	0x77, 0xb5, 0x24, 0x7c, 0x80, 0x23, 0xc7, 0x0c, 0xf2, 0x7f, 0xa1, 0xbe, 0x3b, 0xc4, 0x67, 0x9e,
	0xce, 0x86, 0x5b, 0x1c, 0xe0, 0xdb, 0xe0, 0x78, 0x5d, 0x28, 0x78, 0xb9, 0x33, 0xe4, 0xa6, 0x1c,
	0x27, 0xfb, 0x01, 0x5b, 0xf2, 0x0c, 0x22, 0xd7, 0x14, 0xb6, 0x5a, 0xbc, 0x99, 0xc5, 0xb1, 0xaf,
	0xd5, 0x53, 0x15, 0xe3, 0xf6, 0xff, 0xde, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x3b, 0xfd,
	0x01, 0xd3, 0x02, 0x00, 0x00,
}
