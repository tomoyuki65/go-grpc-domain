// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: chat/chat.proto

package chat

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TextInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TextInput) Reset() {
	*x = TextInput{}
	mi := &file_chat_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TextInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextInput) ProtoMessage() {}

func (x *TextInput) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextInput.ProtoReflect.Descriptor instead.
func (*TextInput) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *TextInput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TextOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TextOutput) Reset() {
	*x = TextOutput{}
	mi := &file_chat_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TextOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextOutput) ProtoMessage() {}

func (x *TextOutput) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextOutput.ProtoReflect.Descriptor instead.
func (*TextOutput) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *TextOutput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_chat_chat_proto protoreflect.FileDescriptor

const file_chat_chat_proto_rawDesc = "" +
	"\n" +
	"\x0fchat/chat.proto\x12\x04chat\x1a\x17validate/validate.proto\"(\n" +
	"\tTextInput\x12\x1b\n" +
	"\x04text\x18\x01 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x04text\" \n" +
	"\n" +
	"TextOutput\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text2G\n" +
	"\vChatService\x128\n" +
	"\rBidirectional\x12\x0f.chat.TextInput\x1a\x10.chat.TextOutput\"\x00(\x010\x01B\tZ\apb/chatb\x06proto3"

var (
	file_chat_chat_proto_rawDescOnce sync.Once
	file_chat_chat_proto_rawDescData []byte
)

func file_chat_chat_proto_rawDescGZIP() []byte {
	file_chat_chat_proto_rawDescOnce.Do(func() {
		file_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_chat_chat_proto_rawDesc), len(file_chat_chat_proto_rawDesc)))
	})
	return file_chat_chat_proto_rawDescData
}

var file_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chat_chat_proto_goTypes = []any{
	(*TextInput)(nil),  // 0: chat.TextInput
	(*TextOutput)(nil), // 1: chat.TextOutput
}
var file_chat_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.Bidirectional:input_type -> chat.TextInput
	1, // 1: chat.ChatService.Bidirectional:output_type -> chat.TextOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_chat_proto_init() }
func file_chat_chat_proto_init() {
	if File_chat_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_chat_chat_proto_rawDesc), len(file_chat_chat_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_chat_proto_goTypes,
		DependencyIndexes: file_chat_chat_proto_depIdxs,
		MessageInfos:      file_chat_chat_proto_msgTypes,
	}.Build()
	File_chat_chat_proto = out.File
	file_chat_chat_proto_goTypes = nil
	file_chat_chat_proto_depIdxs = nil
}
