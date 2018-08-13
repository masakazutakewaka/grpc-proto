// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

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

type Item struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_d2cced31af62beaf, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (dst *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(dst, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type GetItemRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemRequest) Reset()         { *m = GetItemRequest{} }
func (m *GetItemRequest) String() string { return proto.CompactTextString(m) }
func (*GetItemRequest) ProtoMessage()    {}
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_d2cced31af62beaf, []int{1}
}
func (m *GetItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemRequest.Unmarshal(m, b)
}
func (m *GetItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemRequest.Marshal(b, m, deterministic)
}
func (dst *GetItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemRequest.Merge(dst, src)
}
func (m *GetItemRequest) XXX_Size() int {
	return xxx_messageInfo_GetItemRequest.Size(m)
}
func (m *GetItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemRequest proto.InternalMessageInfo

func (m *GetItemRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetItemResponse struct {
	Item                 *Item    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemResponse) Reset()         { *m = GetItemResponse{} }
func (m *GetItemResponse) String() string { return proto.CompactTextString(m) }
func (*GetItemResponse) ProtoMessage()    {}
func (*GetItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_d2cced31af62beaf, []int{2}
}
func (m *GetItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemResponse.Unmarshal(m, b)
}
func (m *GetItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemResponse.Marshal(b, m, deterministic)
}
func (dst *GetItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemResponse.Merge(dst, src)
}
func (m *GetItemResponse) XXX_Size() int {
	return xxx_messageInfo_GetItemResponse.Size(m)
}
func (m *GetItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemResponse proto.InternalMessageInfo

func (m *GetItemResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type GetItemsRequest struct {
	Skip                 int32    `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Take                 int32    `protobuf:"varint,2,opt,name=take,proto3" json:"take,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemsRequest) Reset()         { *m = GetItemsRequest{} }
func (m *GetItemsRequest) String() string { return proto.CompactTextString(m) }
func (*GetItemsRequest) ProtoMessage()    {}
func (*GetItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_d2cced31af62beaf, []int{3}
}
func (m *GetItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemsRequest.Unmarshal(m, b)
}
func (m *GetItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemsRequest.Marshal(b, m, deterministic)
}
func (dst *GetItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemsRequest.Merge(dst, src)
}
func (m *GetItemsRequest) XXX_Size() int {
	return xxx_messageInfo_GetItemsRequest.Size(m)
}
func (m *GetItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemsRequest proto.InternalMessageInfo

func (m *GetItemsRequest) GetSkip() int32 {
	if m != nil {
		return m.Skip
	}
	return 0
}

func (m *GetItemsRequest) GetTake() int32 {
	if m != nil {
		return m.Take
	}
	return 0
}

type GetItemsResponse struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemsResponse) Reset()         { *m = GetItemsResponse{} }
func (m *GetItemsResponse) String() string { return proto.CompactTextString(m) }
func (*GetItemsResponse) ProtoMessage()    {}
func (*GetItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_d2cced31af62beaf, []int{4}
}
func (m *GetItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemsResponse.Unmarshal(m, b)
}
func (m *GetItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemsResponse.Marshal(b, m, deterministic)
}
func (dst *GetItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemsResponse.Merge(dst, src)
}
func (m *GetItemsResponse) XXX_Size() int {
	return xxx_messageInfo_GetItemsResponse.Size(m)
}
func (m *GetItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemsResponse proto.InternalMessageInfo

func (m *GetItemsResponse) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type PostItemRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostItemRequest) Reset()         { *m = PostItemRequest{} }
func (m *PostItemRequest) String() string { return proto.CompactTextString(m) }
func (*PostItemRequest) ProtoMessage()    {}
func (*PostItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_d2cced31af62beaf, []int{5}
}
func (m *PostItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostItemRequest.Unmarshal(m, b)
}
func (m *PostItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostItemRequest.Marshal(b, m, deterministic)
}
func (dst *PostItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostItemRequest.Merge(dst, src)
}
func (m *PostItemRequest) XXX_Size() int {
	return xxx_messageInfo_PostItemRequest.Size(m)
}
func (m *PostItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostItemRequest proto.InternalMessageInfo

func (m *PostItemRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PostItemRequest) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type PostItemResponse struct {
	Item                 *Item    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostItemResponse) Reset()         { *m = PostItemResponse{} }
func (m *PostItemResponse) String() string { return proto.CompactTextString(m) }
func (*PostItemResponse) ProtoMessage()    {}
func (*PostItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_d2cced31af62beaf, []int{6}
}
func (m *PostItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostItemResponse.Unmarshal(m, b)
}
func (m *PostItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostItemResponse.Marshal(b, m, deterministic)
}
func (dst *PostItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostItemResponse.Merge(dst, src)
}
func (m *PostItemResponse) XXX_Size() int {
	return xxx_messageInfo_PostItemResponse.Size(m)
}
func (m *PostItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostItemResponse proto.InternalMessageInfo

func (m *PostItemResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "pb.Item")
	proto.RegisterType((*GetItemRequest)(nil), "pb.getItemRequest")
	proto.RegisterType((*GetItemResponse)(nil), "pb.getItemResponse")
	proto.RegisterType((*GetItemsRequest)(nil), "pb.getItemsRequest")
	proto.RegisterType((*GetItemsResponse)(nil), "pb.getItemsResponse")
	proto.RegisterType((*PostItemRequest)(nil), "pb.postItemRequest")
	proto.RegisterType((*PostItemResponse)(nil), "pb.postItemResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemServiceClient interface {
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
	GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*GetItemsResponse, error)
	PostItem(ctx context.Context, in *PostItemRequest, opts ...grpc.CallOption) (*PostItemResponse, error)
}

type itemServiceClient struct {
	cc *grpc.ClientConn
}

func NewItemServiceClient(cc *grpc.ClientConn) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, "/pb.ItemService/getItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*GetItemsResponse, error) {
	out := new(GetItemsResponse)
	err := c.cc.Invoke(ctx, "/pb.ItemService/getItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) PostItem(ctx context.Context, in *PostItemRequest, opts ...grpc.CallOption) (*PostItemResponse, error) {
	out := new(PostItemResponse)
	err := c.cc.Invoke(ctx, "/pb.ItemService/postItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServiceServer is the server API for ItemService service.
type ItemServiceServer interface {
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
	GetItems(context.Context, *GetItemsRequest) (*GetItemsResponse, error)
	PostItem(context.Context, *PostItemRequest) (*PostItemResponse, error)
}

func RegisterItemServiceServer(s *grpc.Server, srv ItemServiceServer) {
	s.RegisterService(&_ItemService_serviceDesc, srv)
}

func _ItemService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ItemService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ItemService/GetItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItems(ctx, req.(*GetItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_PostItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).PostItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ItemService/PostItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).PostItem(ctx, req.(*PostItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getItem",
			Handler:    _ItemService_GetItem_Handler,
		},
		{
			MethodName: "getItems",
			Handler:    _ItemService_GetItems_Handler,
		},
		{
			MethodName: "postItem",
			Handler:    _ItemService_PostItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "item.proto",
}

func init() { proto.RegisterFile("item.proto", fileDescriptor_item_d2cced31af62beaf) }

var fileDescriptor_item_d2cced31af62beaf = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0xac, 0xdd, 0x04, 0xca, 0x56, 0x6a, 0xab, 0xa5, 0x87, 0xa8, 0x42, 0x28, 0xf2, 0xa9, 0xa7,
	0x82, 0x02, 0x12, 0x42, 0x5c, 0xb8, 0x72, 0x0d, 0x2f, 0x68, 0xda, 0x15, 0xb2, 0xaa, 0x34, 0x26,
	0x36, 0xfc, 0x8c, 0xff, 0x21, 0xdb, 0x71, 0x1b, 0x47, 0x1c, 0xb8, 0xad, 0x27, 0x33, 0xb3, 0x33,
	0xab, 0x00, 0x48, 0x43, 0xf5, 0x46, 0xb5, 0x8d, 0x69, 0x90, 0xab, 0x4a, 0xbc, 0x42, 0xf2, 0x66,
	0xa8, 0xc6, 0x19, 0x70, 0xb9, 0xcf, 0x58, 0xce, 0xd6, 0x69, 0xc9, 0xe5, 0x1e, 0x11, 0x92, 0xe3,
	0xb6, 0xa6, 0x8c, 0xe7, 0x6c, 0x7d, 0x55, 0xba, 0x19, 0x97, 0x90, 0xaa, 0x56, 0xee, 0x28, 0x1b,
	0x3b, 0x9a, 0x7f, 0x88, 0x1c, 0x66, 0x1f, 0x64, 0xac, 0x49, 0x49, 0x9f, 0x5f, 0xa4, 0xcd, 0xd0,
	0x4b, 0xdc, 0xc1, 0xfc, 0xc4, 0xd0, 0xaa, 0x39, 0x6a, 0xc2, 0x1b, 0x48, 0x6c, 0x10, 0x47, 0x9a,
	0x16, 0x93, 0x8d, 0xaa, 0x36, 0xee, 0xbb, 0x43, 0xc5, 0xf3, 0x49, 0xa0, 0x83, 0x27, 0x42, 0xa2,
	0x0f, 0x52, 0x75, 0xae, 0x6e, 0xb6, 0x98, 0xd9, 0x1e, 0x7c, 0xc6, 0xb4, 0x74, 0xb3, 0x28, 0x60,
	0x71, 0x96, 0x76, 0xcb, 0x6e, 0x21, 0xb5, 0xb6, 0x3a, 0x63, 0xf9, 0x38, 0xda, 0xe6, 0x61, 0xf1,
	0x02, 0x73, 0xd5, 0xe8, 0xa8, 0x42, 0xa8, 0xcf, 0xfe, 0xaa, 0xcf, 0xfb, 0xf5, 0xef, 0x61, 0x71,
	0x16, 0xff, 0xa7, 0x5d, 0xf1, 0xc3, 0x60, 0x6a, 0x9f, 0xef, 0xd4, 0x7e, 0xcb, 0x1d, 0xe1, 0x23,
	0x5c, 0x76, 0x91, 0x11, 0x2d, 0x35, 0xbe, 0xe6, 0xea, 0x3a, 0xc2, 0xfc, 0x06, 0x31, 0xc2, 0x27,
	0x98, 0x84, 0xa2, 0xd8, 0xa7, 0x84, 0x8b, 0xad, 0x96, 0x31, 0xd8, 0x17, 0x86, 0xc0, 0x5e, 0x38,
	0xe8, 0xee, 0x85, 0xc3, 0x4e, 0x62, 0x54, 0x5d, 0xb8, 0xbf, 0xe6, 0xe1, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x0d, 0x3e, 0xd6, 0x98, 0x43, 0x02, 0x00, 0x00,
}