// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/client/client.proto

package client

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

type ClientBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   int64                        `protobuf:"varint,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Messages []*ClientBatch_SingleMessage `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"` // a batch of client requests
	Id       string                       `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`             // unique identifier for a batch
}

func (x *ClientBatch) Reset() {
	*x = ClientBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_client_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientBatch) ProtoMessage() {}

func (x *ClientBatch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_client_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientBatch.ProtoReflect.Descriptor instead.
func (*ClientBatch) Descriptor() ([]byte, []int) {
	return file_proto_client_client_proto_rawDescGZIP(), []int{0}
}

func (x *ClientBatch) GetSender() int64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *ClientBatch) GetMessages() []*ClientBatch_SingleMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ClientBatch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ClientStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    int64  `protobuf:"varint,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Operation int64  `protobuf:"varint,2,opt,name=operation,proto3" json:"operation,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"` // optional
}

func (x *ClientStatus) Reset() {
	*x = ClientStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_client_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStatus) ProtoMessage() {}

func (x *ClientStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_client_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStatus.ProtoReflect.Descriptor instead.
func (*ClientStatus) Descriptor() ([]byte, []int) {
	return file_proto_client_client_proto_rawDescGZIP(), []int{1}
}

func (x *ClientStatus) GetSender() int64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *ClientStatus) GetOperation() int64 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *ClientStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ClientBatch_SingleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // a single client request
}

func (x *ClientBatch_SingleMessage) Reset() {
	*x = ClientBatch_SingleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_client_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientBatch_SingleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientBatch_SingleMessage) ProtoMessage() {}

func (x *ClientBatch_SingleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_client_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientBatch_SingleMessage.ProtoReflect.Descriptor instead.
func (*ClientBatch_SingleMessage) Descriptor() ([]byte, []int) {
	return file_proto_client_client_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ClientBatch_SingleMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_client_client_proto protoreflect.FileDescriptor

var file_proto_client_client_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0b,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x29, 0x0a, 0x0d, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5e, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_client_client_proto_rawDescOnce sync.Once
	file_proto_client_client_proto_rawDescData = file_proto_client_client_proto_rawDesc
)

func file_proto_client_client_proto_rawDescGZIP() []byte {
	file_proto_client_client_proto_rawDescOnce.Do(func() {
		file_proto_client_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_client_client_proto_rawDescData)
	})
	return file_proto_client_client_proto_rawDescData
}

var file_proto_client_client_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_client_client_proto_goTypes = []interface{}{
	(*ClientBatch)(nil),               // 0: ClientBatch
	(*ClientStatus)(nil),              // 1: ClientStatus
	(*ClientBatch_SingleMessage)(nil), // 2: ClientBatch.SingleMessage
}
var file_proto_client_client_proto_depIdxs = []int32{
	2, // 0: ClientBatch.messages:type_name -> ClientBatch.SingleMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_client_client_proto_init() }
func file_proto_client_client_proto_init() {
	if File_proto_client_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_client_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientBatch); i {
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
		file_proto_client_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStatus); i {
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
		file_proto_client_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientBatch_SingleMessage); i {
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
			RawDescriptor: file_proto_client_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_client_client_proto_goTypes,
		DependencyIndexes: file_proto_client_client_proto_depIdxs,
		MessageInfos:      file_proto_client_client_proto_msgTypes,
	}.Build()
	File_proto_client_client_proto = out.File
	file_proto_client_client_proto_rawDesc = nil
	file_proto_client_client_proto_goTypes = nil
	file_proto_client_client_proto_depIdxs = nil
}
