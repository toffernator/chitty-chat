// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: ClientToServer.proto

package protofiles

import (
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

type Status_StatusCode int32

const (
	Status_OK                  Status_StatusCode = 0
	Status_INVALIDMSG          Status_StatusCode = 1
	Status_INCOMPLETEBROADCAST Status_StatusCode = 2
)

// Enum value maps for Status_StatusCode.
var (
	Status_StatusCode_name = map[int32]string{
		0: "OK",
		1: "INVALIDMSG",
		2: "INCOMPLETEBROADCAST",
	}
	Status_StatusCode_value = map[string]int32{
		"OK":                  0,
		"INVALIDMSG":          1,
		"INCOMPLETEBROADCAST": 2,
	}
)

func (x Status_StatusCode) Enum() *Status_StatusCode {
	p := new(Status_StatusCode)
	*p = x
	return p
}

func (x Status_StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_ClientToServer_proto_enumTypes[0].Descriptor()
}

func (Status_StatusCode) Type() protoreflect.EnumType {
	return &file_ClientToServer_proto_enumTypes[0]
}

func (x Status_StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_StatusCode.Descriptor instead.
func (Status_StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_ClientToServer_proto_rawDescGZIP(), []int{2, 0}
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTs int64  `protobuf:"varint,1,opt,name=lamportTs,proto3" json:"lamportTs,omitempty"`
	Address   string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientToServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_ClientToServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_ClientToServer_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetLamportTs() int64 {
	if x != nil {
		return x.LamportTs
	}
	return 0
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type StatusOk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTs int64 `protobuf:"varint,1,opt,name=lamportTs,proto3" json:"lamportTs,omitempty"`
}

func (x *StatusOk) Reset() {
	*x = StatusOk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientToServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusOk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusOk) ProtoMessage() {}

func (x *StatusOk) ProtoReflect() protoreflect.Message {
	mi := &file_ClientToServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusOk.ProtoReflect.Descriptor instead.
func (*StatusOk) Descriptor() ([]byte, []int) {
	return file_ClientToServer_proto_rawDescGZIP(), []int{1}
}

func (x *StatusOk) GetLamportTs() int64 {
	if x != nil {
		return x.LamportTs
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTs int64 `protobuf:"varint,1,opt,name=lamportTs,proto3" json:"lamportTs,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientToServer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_ClientToServer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_ClientToServer_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetLamportTs() int64 {
	if x != nil {
		return x.LamportTs
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTs int64  `protobuf:"varint,1,opt,name=lamportTs,proto3" json:"lamportTs,omitempty"`
	Contents  string `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientToServer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_ClientToServer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_ClientToServer_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetLamportTs() int64 {
	if x != nil {
		return x.LamportTs
	}
	return 0
}

func (x *Message) GetContents() string {
	if x != nil {
		return x.Contents
	}
	return ""
}

var File_ClientToServer_proto protoreflect.FileDescriptor

var file_ClientToServer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x28, 0x0a, 0x08, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4f, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x73, 0x22, 0x65, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x73, 0x22, 0x3d, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x4d, 0x53, 0x47, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x42,
	0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x10, 0x02, 0x22, 0x43, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x32,
	0x70, 0x0a, 0x15, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e,
	0x12, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x09, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4f, 0x6b, 0x12, 0x1c, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x08,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x09, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4f, 0x6b, 0x12, 0x1c, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x08,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x68, 0x69, 0x74,
	0x74, 0x79, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClientToServer_proto_rawDescOnce sync.Once
	file_ClientToServer_proto_rawDescData = file_ClientToServer_proto_rawDesc
)

func file_ClientToServer_proto_rawDescGZIP() []byte {
	file_ClientToServer_proto_rawDescOnce.Do(func() {
		file_ClientToServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientToServer_proto_rawDescData)
	})
	return file_ClientToServer_proto_rawDescData
}

var file_ClientToServer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ClientToServer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ClientToServer_proto_goTypes = []interface{}{
	(Status_StatusCode)(0), // 0: Status.StatusCode
	(*Address)(nil),        // 1: Address
	(*StatusOk)(nil),       // 2: StatusOk
	(*Status)(nil),         // 3: Status
	(*Message)(nil),        // 4: Message
}
var file_ClientToServer_proto_depIdxs = []int32{
	1, // 0: ClientToServerService.Join:input_type -> Address
	1, // 1: ClientToServerService.Leave:input_type -> Address
	4, // 2: ClientToServerService.Publish:input_type -> Message
	2, // 3: ClientToServerService.Join:output_type -> StatusOk
	2, // 4: ClientToServerService.Leave:output_type -> StatusOk
	3, // 5: ClientToServerService.Publish:output_type -> Status
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ClientToServer_proto_init() }
func file_ClientToServer_proto_init() {
	if File_ClientToServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ClientToServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_ClientToServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusOk); i {
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
		file_ClientToServer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_ClientToServer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ClientToServer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ClientToServer_proto_goTypes,
		DependencyIndexes: file_ClientToServer_proto_depIdxs,
		EnumInfos:         file_ClientToServer_proto_enumTypes,
		MessageInfos:      file_ClientToServer_proto_msgTypes,
	}.Build()
	File_ClientToServer_proto = out.File
	file_ClientToServer_proto_rawDesc = nil
	file_ClientToServer_proto_goTypes = nil
	file_ClientToServer_proto_depIdxs = nil
}