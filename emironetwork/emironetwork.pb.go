// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        (unknown)
// source: emironetwork.proto

package emironetwork

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query      string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Count      int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	All        bool   `protobuf:"varint,3,opt,name=all,proto3" json:"all,omitempty"`
	RemoteHost string `protobuf:"bytes,4,opt,name=remoteHost,proto3" json:"remoteHost,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emironetwork_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_emironetwork_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_emironetwork_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Query) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Query) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

func (x *Query) GetRemoteHost() string {
	if x != nil {
		return x.RemoteHost
	}
	return ""
}

type QueryFull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *QueryFull) Reset() {
	*x = QueryFull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emironetwork_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFull) ProtoMessage() {}

func (x *QueryFull) ProtoReflect() protoreflect.Message {
	mi := &file_emironetwork_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFull.ProtoReflect.Descriptor instead.
func (*QueryFull) Descriptor() ([]byte, []int) {
	return file_emironetwork_proto_rawDescGZIP(), []int{1}
}

func (x *QueryFull) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description  string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Command      string            `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	Language     string            `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	Path         string            `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Script       bool              `protobuf:"varint,6,opt,name=script,proto3" json:"script,omitempty"`
	Interactive  bool              `protobuf:"varint,7,opt,name=interactive,proto3" json:"interactive,omitempty"`
	Os           []string          `protobuf:"bytes,8,rep,name=os,proto3" json:"os,omitempty"`
	Params       map[string]string `protobuf:"bytes,9,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RemoteOutput string            `protobuf:"bytes,10,opt,name=remoteOutput,proto3" json:"remoteOutput,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emironetwork_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_emironetwork_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_emironetwork_proto_rawDescGZIP(), []int{2}
}

func (x *Answer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Answer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Answer) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Answer) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Answer) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Answer) GetScript() bool {
	if x != nil {
		return x.Script
	}
	return false
}

func (x *Answer) GetInteractive() bool {
	if x != nil {
		return x.Interactive
	}
	return false
}

func (x *Answer) GetOs() []string {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *Answer) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Answer) GetRemoteOutput() string {
	if x != nil {
		return x.RemoteOutput
	}
	return ""
}

type ResponseQueryAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryAnswers []*Answer `protobuf:"bytes,1,rep,name=queryAnswers,proto3" json:"queryAnswers,omitempty"`
}

func (x *ResponseQueryAnswer) Reset() {
	*x = ResponseQueryAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emironetwork_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseQueryAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseQueryAnswer) ProtoMessage() {}

func (x *ResponseQueryAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_emironetwork_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseQueryAnswer.ProtoReflect.Descriptor instead.
func (*ResponseQueryAnswer) Descriptor() ([]byte, []int) {
	return file_emironetwork_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseQueryAnswer) GetQueryAnswers() []*Answer {
	if x != nil {
		return x.QueryAnswers
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emironetwork_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_emironetwork_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_emironetwork_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

var File_emironetwork_proto protoreflect.FileDescriptor

var file_emironetwork_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x6d, 0x69, 0x72, 0x6f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6d, 0x69, 0x72, 0x6f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x22, 0x65, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x09, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x46, 0x75, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xeb, 0x02, 0x0a, 0x06, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12,
	0x38, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x65, 0x6d, 0x69, 0x72, 0x6f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x1a, 0x39, 0x0a,
	0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12,
	0x38, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6d, 0x69, 0x72, 0x6f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x0c, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x32,
	0xfe, 0x01, 0x0a, 0x05, 0x45, 0x6d, 0x69, 0x72, 0x6f, 0x12, 0x45, 0x0a, 0x09, 0x73, 0x65, 0x6e,
	0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x13, 0x2e, 0x65, 0x6d, 0x69, 0x72, 0x6f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x21, 0x2e, 0x65, 0x6d,
	0x69, 0x72, 0x6f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x13, 0x2e, 0x65,
	0x6d, 0x69, 0x72, 0x6f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x1a, 0x14, 0x2e, 0x65, 0x6d, 0x69, 0x72, 0x6f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x45, 0x78, 0x65, 0x63, 0x12, 0x13, 0x2e, 0x65, 0x6d, 0x69, 0x72, 0x6f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x14, 0x2e, 0x65, 0x6d, 0x69,
	0x72, 0x6f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x12, 0x17, 0x2e,
	0x65, 0x6d, 0x69, 0x72, 0x6f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x46, 0x75, 0x6c, 0x6c, 0x1a, 0x16, 0x2e, 0x65, 0x6d, 0x69, 0x72, 0x6f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_emironetwork_proto_rawDescOnce sync.Once
	file_emironetwork_proto_rawDescData = file_emironetwork_proto_rawDesc
)

func file_emironetwork_proto_rawDescGZIP() []byte {
	file_emironetwork_proto_rawDescOnce.Do(func() {
		file_emironetwork_proto_rawDescData = protoimpl.X.CompressGZIP(file_emironetwork_proto_rawDescData)
	})
	return file_emironetwork_proto_rawDescData
}

var file_emironetwork_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_emironetwork_proto_goTypes = []interface{}{
	(*Query)(nil),               // 0: emironetwork.Query
	(*QueryFull)(nil),           // 1: emironetwork.QueryFull
	(*Answer)(nil),              // 2: emironetwork.Answer
	(*ResponseQueryAnswer)(nil), // 3: emironetwork.ResponseQueryAnswer
	(*Response)(nil),            // 4: emironetwork.Response
	nil,                         // 5: emironetwork.Answer.ParamsEntry
}
var file_emironetwork_proto_depIdxs = []int32{
	5, // 0: emironetwork.Answer.params:type_name -> emironetwork.Answer.ParamsEntry
	2, // 1: emironetwork.ResponseQueryAnswer.queryAnswers:type_name -> emironetwork.Answer
	0, // 2: emironetwork.Emiro.sendQuery:input_type -> emironetwork.Query
	0, // 3: emironetwork.Emiro.sendShow:input_type -> emironetwork.Query
	0, // 4: emironetwork.Emiro.sendExec:input_type -> emironetwork.Query
	1, // 5: emironetwork.Emiro.sendNew:input_type -> emironetwork.QueryFull
	3, // 6: emironetwork.Emiro.sendQuery:output_type -> emironetwork.ResponseQueryAnswer
	2, // 7: emironetwork.Emiro.sendShow:output_type -> emironetwork.Answer
	2, // 8: emironetwork.Emiro.sendExec:output_type -> emironetwork.Answer
	4, // 9: emironetwork.Emiro.sendNew:output_type -> emironetwork.Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_emironetwork_proto_init() }
func file_emironetwork_proto_init() {
	if File_emironetwork_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_emironetwork_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_emironetwork_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFull); i {
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
		file_emironetwork_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_emironetwork_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseQueryAnswer); i {
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
		file_emironetwork_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_emironetwork_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_emironetwork_proto_goTypes,
		DependencyIndexes: file_emironetwork_proto_depIdxs,
		MessageInfos:      file_emironetwork_proto_msgTypes,
	}.Build()
	File_emironetwork_proto = out.File
	file_emironetwork_proto_rawDesc = nil
	file_emironetwork_proto_goTypes = nil
	file_emironetwork_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmiroClient is the client API for Emiro service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmiroClient interface {
	SendQuery(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ResponseQueryAnswer, error)
	SendShow(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Answer, error)
	SendExec(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Answer, error)
	SendNew(ctx context.Context, in *QueryFull, opts ...grpc.CallOption) (*Response, error)
}

type emiroClient struct {
	cc grpc.ClientConnInterface
}

func NewEmiroClient(cc grpc.ClientConnInterface) EmiroClient {
	return &emiroClient{cc}
}

func (c *emiroClient) SendQuery(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ResponseQueryAnswer, error) {
	out := new(ResponseQueryAnswer)
	err := c.cc.Invoke(ctx, "/emironetwork.Emiro/sendQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emiroClient) SendShow(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/emironetwork.Emiro/sendShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emiroClient) SendExec(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/emironetwork.Emiro/sendExec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emiroClient) SendNew(ctx context.Context, in *QueryFull, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/emironetwork.Emiro/sendNew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmiroServer is the server API for Emiro service.
type EmiroServer interface {
	SendQuery(context.Context, *Query) (*ResponseQueryAnswer, error)
	SendShow(context.Context, *Query) (*Answer, error)
	SendExec(context.Context, *Query) (*Answer, error)
	SendNew(context.Context, *QueryFull) (*Response, error)
}

// UnimplementedEmiroServer can be embedded to have forward compatible implementations.
type UnimplementedEmiroServer struct {
}

func (*UnimplementedEmiroServer) SendQuery(context.Context, *Query) (*ResponseQueryAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendQuery not implemented")
}
func (*UnimplementedEmiroServer) SendShow(context.Context, *Query) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendShow not implemented")
}
func (*UnimplementedEmiroServer) SendExec(context.Context, *Query) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendExec not implemented")
}
func (*UnimplementedEmiroServer) SendNew(context.Context, *QueryFull) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNew not implemented")
}

func RegisterEmiroServer(s *grpc.Server, srv EmiroServer) {
	s.RegisterService(&_Emiro_serviceDesc, srv)
}

func _Emiro_SendQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmiroServer).SendQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emironetwork.Emiro/SendQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmiroServer).SendQuery(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emiro_SendShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmiroServer).SendShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emironetwork.Emiro/SendShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmiroServer).SendShow(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emiro_SendExec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmiroServer).SendExec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emironetwork.Emiro/SendExec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmiroServer).SendExec(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emiro_SendNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFull)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmiroServer).SendNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emironetwork.Emiro/SendNew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmiroServer).SendNew(ctx, req.(*QueryFull))
	}
	return interceptor(ctx, in, info, handler)
}

var _Emiro_serviceDesc = grpc.ServiceDesc{
	ServiceName: "emironetwork.Emiro",
	HandlerType: (*EmiroServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendQuery",
			Handler:    _Emiro_SendQuery_Handler,
		},
		{
			MethodName: "sendShow",
			Handler:    _Emiro_SendShow_Handler,
		},
		{
			MethodName: "sendExec",
			Handler:    _Emiro_SendExec_Handler,
		},
		{
			MethodName: "sendNew",
			Handler:    _Emiro_SendNew_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emironetwork.proto",
}
