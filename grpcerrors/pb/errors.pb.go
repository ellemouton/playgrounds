// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: errors.proto

package pb

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
		mi := &file_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_errors_proto_msgTypes[0]
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
	return file_errors_proto_rawDescGZIP(), []int{0}
}

type StreamItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item string `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *StreamItem) Reset() {
	*x = StreamItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamItem) ProtoMessage() {}

func (x *StreamItem) ProtoReflect() protoreflect.Message {
	mi := &file_errors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamItem.ProtoReflect.Descriptor instead.
func (*StreamItem) Descriptor() ([]byte, []int) {
	return file_errors_proto_rawDescGZIP(), []int{1}
}

func (x *StreamItem) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Error:
	//	*Error_RuleViolationErr
	//	*Error_InternalErr
	Error isError_Error `protobuf_oneof:"error"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_errors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_errors_proto_rawDescGZIP(), []int{2}
}

func (m *Error) GetError() isError_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (x *Error) GetRuleViolationErr() *ErrRuleViolation {
	if x, ok := x.GetError().(*Error_RuleViolationErr); ok {
		return x.RuleViolationErr
	}
	return nil
}

func (x *Error) GetInternalErr() *ErrInternal {
	if x, ok := x.GetError().(*Error_InternalErr); ok {
		return x.InternalErr
	}
	return nil
}

type isError_Error interface {
	isError_Error()
}

type Error_RuleViolationErr struct {
	RuleViolationErr *ErrRuleViolation `protobuf:"bytes,1,opt,name=rule_violation_err,json=ruleViolationErr,proto3,oneof"`
}

type Error_InternalErr struct {
	InternalErr *ErrInternal `protobuf:"bytes,2,opt,name=internal_err,json=internalErr,proto3,oneof"`
}

func (*Error_RuleViolationErr) isError_Error() {}

func (*Error_InternalErr) isError_Error() {}

type ErrRuleViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleName string `protobuf:"bytes,1,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	Err      string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ErrRuleViolation) Reset() {
	*x = ErrRuleViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrRuleViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrRuleViolation) ProtoMessage() {}

func (x *ErrRuleViolation) ProtoReflect() protoreflect.Message {
	mi := &file_errors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrRuleViolation.ProtoReflect.Descriptor instead.
func (*ErrRuleViolation) Descriptor() ([]byte, []int) {
	return file_errors_proto_rawDescGZIP(), []int{3}
}

func (x *ErrRuleViolation) GetRuleName() string {
	if x != nil {
		return x.RuleName
	}
	return ""
}

func (x *ErrRuleViolation) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type ErrInternal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ErrInternal) Reset() {
	*x = ErrInternal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrInternal) ProtoMessage() {}

func (x *ErrInternal) ProtoReflect() protoreflect.Message {
	mi := &file_errors_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrInternal.ProtoReflect.Descriptor instead.
func (*ErrInternal) Descriptor() ([]byte, []int) {
	return file_errors_proto_rawDescGZIP(), []int{4}
}

func (x *ErrInternal) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_errors_proto protoreflect.FileDescriptor

var file_errors_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x0a, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x8c, 0x01,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x44, 0x0a, 0x12, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x72, 0x72, 0x52, 0x75, 0x6c, 0x65,
	0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x72, 0x75, 0x6c,
	0x65, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x12, 0x34, 0x0a,
	0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x72, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x45, 0x72, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x41, 0x0a, 0x10,
	0x45, 0x72, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22,
	0x1f, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x32, 0xdd, 0x02, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x07, 0x4e,
	0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x08,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x6b, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x0d, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x09, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x14, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x30, 0x01, 0x12, 0x30,
	0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x30, 0x01,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6c, 0x6c, 0x65, 0x6d, 0x6f, 0x75, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_proto_rawDescOnce sync.Once
	file_errors_proto_rawDescData = file_errors_proto_rawDesc
)

func file_errors_proto_rawDescGZIP() []byte {
	file_errors_proto_rawDescOnce.Do(func() {
		file_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_proto_rawDescData)
	})
	return file_errors_proto_rawDescData
}

var file_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_errors_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: pb.Empty
	(*StreamItem)(nil),       // 1: pb.StreamItem
	(*Error)(nil),            // 2: pb.Error
	(*ErrRuleViolation)(nil), // 3: pb.ErrRuleViolation
	(*ErrInternal)(nil),      // 4: pb.ErrInternal
}
var file_errors_proto_depIdxs = []int32{
	3,  // 0: pb.Error.rule_violation_err:type_name -> pb.ErrRuleViolation
	4,  // 1: pb.Error.internal_err:type_name -> pb.ErrInternal
	0,  // 2: pb.Errors.NoError:input_type -> pb.Empty
	0,  // 3: pb.Errors.StatusOk:input_type -> pb.Empty
	0,  // 4: pb.Errors.StatusError:input_type -> pb.Empty
	0,  // 5: pb.Errors.StatusErrorWithDetails:input_type -> pb.Empty
	0,  // 6: pb.Errors.NonStatusError:input_type -> pb.Empty
	0,  // 7: pb.Errors.StreamNoError:input_type -> pb.Empty
	0,  // 8: pb.Errors.StreamNonStatusError:input_type -> pb.Empty
	0,  // 9: pb.Errors.StreamStatusError:input_type -> pb.Empty
	0,  // 10: pb.Errors.NoError:output_type -> pb.Empty
	0,  // 11: pb.Errors.StatusOk:output_type -> pb.Empty
	0,  // 12: pb.Errors.StatusError:output_type -> pb.Empty
	0,  // 13: pb.Errors.StatusErrorWithDetails:output_type -> pb.Empty
	0,  // 14: pb.Errors.NonStatusError:output_type -> pb.Empty
	1,  // 15: pb.Errors.StreamNoError:output_type -> pb.StreamItem
	1,  // 16: pb.Errors.StreamNonStatusError:output_type -> pb.StreamItem
	1,  // 17: pb.Errors.StreamStatusError:output_type -> pb.StreamItem
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_errors_proto_init() }
func file_errors_proto_init() {
	if File_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_errors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamItem); i {
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
		file_errors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_errors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrRuleViolation); i {
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
		file_errors_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrInternal); i {
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
	file_errors_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Error_RuleViolationErr)(nil),
		(*Error_InternalErr)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_errors_proto_goTypes,
		DependencyIndexes: file_errors_proto_depIdxs,
		MessageInfos:      file_errors_proto_msgTypes,
	}.Build()
	File_errors_proto = out.File
	file_errors_proto_rawDesc = nil
	file_errors_proto_goTypes = nil
	file_errors_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ErrorsClient is the client API for Errors service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ErrorsClient interface {
	NoError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	StatusOk(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	StatusError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	StatusErrorWithDetails(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	NonStatusError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	StreamNoError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Errors_StreamNoErrorClient, error)
	StreamNonStatusError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Errors_StreamNonStatusErrorClient, error)
	StreamStatusError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Errors_StreamStatusErrorClient, error)
}

type errorsClient struct {
	cc grpc.ClientConnInterface
}

func NewErrorsClient(cc grpc.ClientConnInterface) ErrorsClient {
	return &errorsClient{cc}
}

func (c *errorsClient) NoError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.Errors/NoError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *errorsClient) StatusOk(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.Errors/StatusOk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *errorsClient) StatusError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.Errors/StatusError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *errorsClient) StatusErrorWithDetails(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.Errors/StatusErrorWithDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *errorsClient) NonStatusError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.Errors/NonStatusError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *errorsClient) StreamNoError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Errors_StreamNoErrorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Errors_serviceDesc.Streams[0], "/pb.Errors/StreamNoError", opts...)
	if err != nil {
		return nil, err
	}
	x := &errorsStreamNoErrorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Errors_StreamNoErrorClient interface {
	Recv() (*StreamItem, error)
	grpc.ClientStream
}

type errorsStreamNoErrorClient struct {
	grpc.ClientStream
}

func (x *errorsStreamNoErrorClient) Recv() (*StreamItem, error) {
	m := new(StreamItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *errorsClient) StreamNonStatusError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Errors_StreamNonStatusErrorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Errors_serviceDesc.Streams[1], "/pb.Errors/StreamNonStatusError", opts...)
	if err != nil {
		return nil, err
	}
	x := &errorsStreamNonStatusErrorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Errors_StreamNonStatusErrorClient interface {
	Recv() (*StreamItem, error)
	grpc.ClientStream
}

type errorsStreamNonStatusErrorClient struct {
	grpc.ClientStream
}

func (x *errorsStreamNonStatusErrorClient) Recv() (*StreamItem, error) {
	m := new(StreamItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *errorsClient) StreamStatusError(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Errors_StreamStatusErrorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Errors_serviceDesc.Streams[2], "/pb.Errors/StreamStatusError", opts...)
	if err != nil {
		return nil, err
	}
	x := &errorsStreamStatusErrorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Errors_StreamStatusErrorClient interface {
	Recv() (*StreamItem, error)
	grpc.ClientStream
}

type errorsStreamStatusErrorClient struct {
	grpc.ClientStream
}

func (x *errorsStreamStatusErrorClient) Recv() (*StreamItem, error) {
	m := new(StreamItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ErrorsServer is the server API for Errors service.
type ErrorsServer interface {
	NoError(context.Context, *Empty) (*Empty, error)
	StatusOk(context.Context, *Empty) (*Empty, error)
	StatusError(context.Context, *Empty) (*Empty, error)
	StatusErrorWithDetails(context.Context, *Empty) (*Empty, error)
	NonStatusError(context.Context, *Empty) (*Empty, error)
	StreamNoError(*Empty, Errors_StreamNoErrorServer) error
	StreamNonStatusError(*Empty, Errors_StreamNonStatusErrorServer) error
	StreamStatusError(*Empty, Errors_StreamStatusErrorServer) error
}

// UnimplementedErrorsServer can be embedded to have forward compatible implementations.
type UnimplementedErrorsServer struct {
}

func (*UnimplementedErrorsServer) NoError(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoError not implemented")
}
func (*UnimplementedErrorsServer) StatusOk(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusOk not implemented")
}
func (*UnimplementedErrorsServer) StatusError(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusError not implemented")
}
func (*UnimplementedErrorsServer) StatusErrorWithDetails(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusErrorWithDetails not implemented")
}
func (*UnimplementedErrorsServer) NonStatusError(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NonStatusError not implemented")
}
func (*UnimplementedErrorsServer) StreamNoError(*Empty, Errors_StreamNoErrorServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamNoError not implemented")
}
func (*UnimplementedErrorsServer) StreamNonStatusError(*Empty, Errors_StreamNonStatusErrorServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamNonStatusError not implemented")
}
func (*UnimplementedErrorsServer) StreamStatusError(*Empty, Errors_StreamStatusErrorServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamStatusError not implemented")
}

func RegisterErrorsServer(s *grpc.Server, srv ErrorsServer) {
	s.RegisterService(&_Errors_serviceDesc, srv)
}

func _Errors_NoError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorsServer).NoError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Errors/NoError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorsServer).NoError(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Errors_StatusOk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorsServer).StatusOk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Errors/StatusOk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorsServer).StatusOk(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Errors_StatusError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorsServer).StatusError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Errors/StatusError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorsServer).StatusError(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Errors_StatusErrorWithDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorsServer).StatusErrorWithDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Errors/StatusErrorWithDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorsServer).StatusErrorWithDetails(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Errors_NonStatusError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorsServer).NonStatusError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Errors/NonStatusError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorsServer).NonStatusError(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Errors_StreamNoError_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ErrorsServer).StreamNoError(m, &errorsStreamNoErrorServer{stream})
}

type Errors_StreamNoErrorServer interface {
	Send(*StreamItem) error
	grpc.ServerStream
}

type errorsStreamNoErrorServer struct {
	grpc.ServerStream
}

func (x *errorsStreamNoErrorServer) Send(m *StreamItem) error {
	return x.ServerStream.SendMsg(m)
}

func _Errors_StreamNonStatusError_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ErrorsServer).StreamNonStatusError(m, &errorsStreamNonStatusErrorServer{stream})
}

type Errors_StreamNonStatusErrorServer interface {
	Send(*StreamItem) error
	grpc.ServerStream
}

type errorsStreamNonStatusErrorServer struct {
	grpc.ServerStream
}

func (x *errorsStreamNonStatusErrorServer) Send(m *StreamItem) error {
	return x.ServerStream.SendMsg(m)
}

func _Errors_StreamStatusError_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ErrorsServer).StreamStatusError(m, &errorsStreamStatusErrorServer{stream})
}

type Errors_StreamStatusErrorServer interface {
	Send(*StreamItem) error
	grpc.ServerStream
}

type errorsStreamStatusErrorServer struct {
	grpc.ServerStream
}

func (x *errorsStreamStatusErrorServer) Send(m *StreamItem) error {
	return x.ServerStream.SendMsg(m)
}

var _Errors_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Errors",
	HandlerType: (*ErrorsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NoError",
			Handler:    _Errors_NoError_Handler,
		},
		{
			MethodName: "StatusOk",
			Handler:    _Errors_StatusOk_Handler,
		},
		{
			MethodName: "StatusError",
			Handler:    _Errors_StatusError_Handler,
		},
		{
			MethodName: "StatusErrorWithDetails",
			Handler:    _Errors_StatusErrorWithDetails_Handler,
		},
		{
			MethodName: "NonStatusError",
			Handler:    _Errors_NonStatusError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNoError",
			Handler:       _Errors_StreamNoError_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamNonStatusError",
			Handler:       _Errors_StreamNonStatusError_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamStatusError",
			Handler:       _Errors_StreamStatusError_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "errors.proto",
}
