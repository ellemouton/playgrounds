// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: service.proto

package pb_new

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type ListItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	New bool `protobuf:"varint,1,opt,name=new,proto3" json:"new,omitempty"`
}

func (x *ListItemsRequest) Reset() {
	*x = ListItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsRequest) ProtoMessage() {}

func (x *ListItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsRequest.ProtoReflect.Descriptor instead.
func (*ListItemsRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListItemsRequest) GetNew() bool {
	if x != nil {
		return x.New
	}
	return false
}

type ListItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Messages:
	//	*ListItemsResponse_OldMsg
	//	*ListItemsResponse_NewMsg
	Messages isListItemsResponse_Messages `protobuf_oneof:"messages"`
}

func (x *ListItemsResponse) Reset() {
	*x = ListItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsResponse) ProtoMessage() {}

func (x *ListItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsResponse.ProtoReflect.Descriptor instead.
func (*ListItemsResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (m *ListItemsResponse) GetMessages() isListItemsResponse_Messages {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (x *ListItemsResponse) GetOldMsg() *OldMessage {
	if x, ok := x.GetMessages().(*ListItemsResponse_OldMsg); ok {
		return x.OldMsg
	}
	return nil
}

func (x *ListItemsResponse) GetNewMsg() *NewMessage {
	if x, ok := x.GetMessages().(*ListItemsResponse_NewMsg); ok {
		return x.NewMsg
	}
	return nil
}

type isListItemsResponse_Messages interface {
	isListItemsResponse_Messages()
}

type ListItemsResponse_OldMsg struct {
	OldMsg *OldMessage `protobuf:"bytes,1,opt,name=old_msg,json=oldMsg,proto3,oneof"`
}

type ListItemsResponse_NewMsg struct {
	NewMsg *NewMessage `protobuf:"bytes,2,opt,name=new_msg,json=newMsg,proto3,oneof"`
}

func (*ListItemsResponse_OldMsg) isListItemsResponse_Messages() {}

func (*ListItemsResponse_NewMsg) isListItemsResponse_Messages() {}

type OldMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *OldMessage) Reset() {
	*x = OldMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OldMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OldMessage) ProtoMessage() {}

func (x *OldMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OldMessage.ProtoReflect.Descriptor instead.
func (*OldMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *OldMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type NewMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg   string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *NewMessage) Reset() {
	*x = NewMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMessage) ProtoMessage() {}

func (x *NewMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMessage.ProtoReflect.Descriptor instead.
func (*NewMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *NewMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *NewMessage) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x62, 0x5f, 0x6f, 0x6c, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x24, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x6e, 0x65, 0x77, 0x22, 0x7d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x6f,
	0x6c, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x62, 0x5f, 0x6f, 0x6c, 0x64, 0x2e, 0x4f, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x06, 0x6f, 0x6c, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x07, 0x6e, 0x65,
	0x77, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62,
	0x5f, 0x6f, 0x6c, 0x64, 0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x4d, 0x73, 0x67, 0x42, 0x0a, 0x0a, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x1e, 0x0a, 0x0a, 0x4f, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x34, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x4c, 0x0a, 0x08, 0x43,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x40, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x5f, 0x6f, 0x6c, 0x64, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x62, 0x5f, 0x6f, 0x6c, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x6c, 0x65, 0x6d, 0x6f, 0x75, 0x74,
	0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2f, 0x70, 0x62, 0x5f, 0x6e, 0x65, 0x77, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_proto_goTypes = []interface{}{
	(*Empty)(nil),             // 0: pb_old.Empty
	(*ListItemsRequest)(nil),  // 1: pb_old.ListItemsRequest
	(*ListItemsResponse)(nil), // 2: pb_old.ListItemsResponse
	(*OldMessage)(nil),        // 3: pb_old.OldMessage
	(*NewMessage)(nil),        // 4: pb_old.NewMessage
}
var file_service_proto_depIdxs = []int32{
	3, // 0: pb_old.ListItemsResponse.old_msg:type_name -> pb_old.OldMessage
	4, // 1: pb_old.ListItemsResponse.new_msg:type_name -> pb_old.NewMessage
	1, // 2: pb_old.Calendar.ListItems:input_type -> pb_old.ListItemsRequest
	2, // 3: pb_old.Calendar.ListItems:output_type -> pb_old.ListItemsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OldMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ListItemsResponse_OldMsg)(nil),
		(*ListItemsResponse_NewMsg)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsResponse, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsResponse, error) {
	out := new(ListItemsResponse)
	err := c.cc.Invoke(ctx, "/pb_old.Calendar/ListItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	ListItems(context.Context, *ListItemsRequest) (*ListItemsResponse, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) ListItems(context.Context, *ListItemsRequest) (*ListItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_old.Calendar/ListItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListItems(ctx, req.(*ListItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb_old.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListItems",
			Handler:    _Calendar_ListItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
