// Code generated by protoc-gen-go. DO NOT EDIT.
// source: emironetwork.proto

package emironetwork

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

type Query struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	All                  bool     `protobuf:"varint,3,opt,name=all,proto3" json:"all,omitempty"`
	RemoteHost           string   `protobuf:"bytes,4,opt,name=remoteHost,proto3" json:"remoteHost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd0f112475f4ca2e, []int{0}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *Query) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Query) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func (m *Query) GetRemoteHost() string {
	if m != nil {
		return m.RemoteHost
	}
	return ""
}

type QueryFull struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryFull) Reset()         { *m = QueryFull{} }
func (m *QueryFull) String() string { return proto.CompactTextString(m) }
func (*QueryFull) ProtoMessage()    {}
func (*QueryFull) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd0f112475f4ca2e, []int{1}
}

func (m *QueryFull) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryFull.Unmarshal(m, b)
}
func (m *QueryFull) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryFull.Marshal(b, m, deterministic)
}
func (m *QueryFull) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFull.Merge(m, src)
}
func (m *QueryFull) XXX_Size() int {
	return xxx_messageInfo_QueryFull.Size(m)
}
func (m *QueryFull) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFull.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFull proto.InternalMessageInfo

func (m *QueryFull) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Answer struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Command              string            `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	Language             string            `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	Path                 string            `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Script               bool              `protobuf:"varint,6,opt,name=script,proto3" json:"script,omitempty"`
	Interactive          bool              `protobuf:"varint,7,opt,name=interactive,proto3" json:"interactive,omitempty"`
	Os                   []string          `protobuf:"bytes,8,rep,name=os,proto3" json:"os,omitempty"`
	Params               map[string]string `protobuf:"bytes,9,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RemoteOutput         string            `protobuf:"bytes,10,opt,name=remoteOutput,proto3" json:"remoteOutput,omitempty"`
	Id                   string            `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd0f112475f4ca2e, []int{2}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Answer) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Answer) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Answer) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Answer) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Answer) GetScript() bool {
	if m != nil {
		return m.Script
	}
	return false
}

func (m *Answer) GetInteractive() bool {
	if m != nil {
		return m.Interactive
	}
	return false
}

func (m *Answer) GetOs() []string {
	if m != nil {
		return m.Os
	}
	return nil
}

func (m *Answer) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Answer) GetRemoteOutput() string {
	if m != nil {
		return m.RemoteOutput
	}
	return ""
}

func (m *Answer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ResponseQueryAnswer struct {
	QueryAnswers         []*Answer `protobuf:"bytes,1,rep,name=queryAnswers,proto3" json:"queryAnswers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResponseQueryAnswer) Reset()         { *m = ResponseQueryAnswer{} }
func (m *ResponseQueryAnswer) String() string { return proto.CompactTextString(m) }
func (*ResponseQueryAnswer) ProtoMessage()    {}
func (*ResponseQueryAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd0f112475f4ca2e, []int{3}
}

func (m *ResponseQueryAnswer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseQueryAnswer.Unmarshal(m, b)
}
func (m *ResponseQueryAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseQueryAnswer.Marshal(b, m, deterministic)
}
func (m *ResponseQueryAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseQueryAnswer.Merge(m, src)
}
func (m *ResponseQueryAnswer) XXX_Size() int {
	return xxx_messageInfo_ResponseQueryAnswer.Size(m)
}
func (m *ResponseQueryAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseQueryAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseQueryAnswer proto.InternalMessageInfo

func (m *ResponseQueryAnswer) GetQueryAnswers() []*Answer {
	if m != nil {
		return m.QueryAnswers
	}
	return nil
}

type Response struct {
	Succeed              bool     `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd0f112475f4ca2e, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

func init() {
	proto.RegisterType((*Query)(nil), "emironetwork.Query")
	proto.RegisterType((*QueryFull)(nil), "emironetwork.QueryFull")
	proto.RegisterType((*Answer)(nil), "emironetwork.Answer")
	proto.RegisterMapType((map[string]string)(nil), "emironetwork.Answer.ParamsEntry")
	proto.RegisterType((*ResponseQueryAnswer)(nil), "emironetwork.ResponseQueryAnswer")
	proto.RegisterType((*Response)(nil), "emironetwork.Response")
}

func init() {
	proto.RegisterFile("emironetwork.proto", fileDescriptor_bd0f112475f4ca2e)
}

var fileDescriptor_bd0f112475f4ca2e = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x6d, 0xb2, 0x4d, 0x36, 0x99, 0x5d, 0x21, 0xe4, 0x56, 0xc5, 0xda, 0x03, 0x84, 0x88, 0x43,
	0x4e, 0x7b, 0x28, 0x07, 0x96, 0x8f, 0x0b, 0x12, 0x8b, 0x38, 0x51, 0x30, 0xbf, 0xc0, 0x24, 0xa3,
	0x36, 0x6a, 0x62, 0xa7, 0xb6, 0xd3, 0xa5, 0xbf, 0x8d, 0x9f, 0xc6, 0x05, 0xd9, 0x49, 0xd8, 0x44,
	0x0a, 0x87, 0xde, 0xe6, 0x3d, 0x7b, 0xe6, 0xbd, 0xc9, 0x73, 0x80, 0x60, 0x5d, 0x2a, 0x29, 0xd0,
	0x1c, 0xa4, 0xba, 0xdd, 0x36, 0x4a, 0x1a, 0x49, 0xd6, 0x63, 0x2e, 0x45, 0x08, 0xbe, 0xb7, 0xa8,
	0x1e, 0xc8, 0x39, 0x04, 0x77, 0xb6, 0xa0, 0x5e, 0xe2, 0x65, 0x31, 0xeb, 0x80, 0x65, 0x73, 0xd9,
	0x0a, 0x43, 0xfd, 0xc4, 0xcb, 0x02, 0xd6, 0x01, 0xf2, 0x14, 0x16, 0xbc, 0xaa, 0xe8, 0x22, 0xf1,
	0xb2, 0x88, 0xd9, 0x92, 0x3c, 0x07, 0x50, 0x58, 0x4b, 0x83, 0x5f, 0xa4, 0x36, 0xf4, 0xd4, 0x8d,
	0x18, 0x31, 0xe9, 0x0b, 0x88, 0x9d, 0xcc, 0xe7, 0xb6, 0xaa, 0x08, 0x81, 0xd3, 0x82, 0x1b, 0xee,
	0x94, 0xd6, 0xcc, 0xd5, 0xe9, 0x1f, 0x1f, 0xc2, 0x8f, 0x42, 0x1f, 0x50, 0xd9, 0x63, 0xc1, 0x6b,
	0xec, 0x8d, 0xb8, 0x9a, 0x24, 0xb0, 0x2a, 0x50, 0xe7, 0xaa, 0x6c, 0x4c, 0x29, 0x85, 0x73, 0x13,
	0xb3, 0x31, 0x45, 0x28, 0x2c, 0x73, 0x59, 0xd7, 0x5c, 0x14, 0xce, 0x57, 0xcc, 0x06, 0x48, 0x36,
	0x10, 0x55, 0x5c, 0x5c, 0xb7, 0xfc, 0x1a, 0x7b, 0x67, 0xff, 0xb0, 0xd5, 0x6a, 0xb8, 0xb9, 0xa1,
	0x41, 0xa7, 0x65, 0x6b, 0x72, 0x01, 0x61, 0x37, 0x96, 0x86, 0x6e, 0xc1, 0x1e, 0x59, 0x0f, 0xa5,
	0x30, 0xa8, 0x78, 0x6e, 0xca, 0x7b, 0xa4, 0x4b, 0x77, 0x38, 0xa6, 0xc8, 0x13, 0xf0, 0xa5, 0xa6,
	0x51, 0xb2, 0xc8, 0x62, 0xe6, 0x4b, 0x4d, 0x76, 0x10, 0x36, 0x5c, 0xf1, 0x5a, 0xd3, 0x38, 0x59,
	0x64, 0xab, 0xcb, 0x64, 0x3b, 0xc9, 0xa3, 0xdb, 0x77, 0xfb, 0xcd, 0x5d, 0xd9, 0x0b, 0xa3, 0x1e,
	0x58, 0x7f, 0x9f, 0xa4, 0xb0, 0xee, 0xbe, 0xde, 0x55, 0x6b, 0x9a, 0xd6, 0x50, 0x70, 0xfe, 0x26,
	0x9c, 0x55, 0x2b, 0x0b, 0xba, 0x72, 0x27, 0x7e, 0x59, 0x6c, 0xde, 0xc2, 0x6a, 0x34, 0xca, 0x86,
	0x74, 0x8b, 0x43, 0x9c, 0xb6, 0xb4, 0x61, 0xde, 0xf3, 0xaa, 0xc5, 0xfe, 0xf3, 0x75, 0xe0, 0x9d,
	0xbf, 0xf3, 0xd2, 0x2b, 0x38, 0x63, 0xa8, 0x1b, 0x29, 0x34, 0xba, 0x98, 0xfa, 0x24, 0x76, 0xb0,
	0xbe, 0x3b, 0x42, 0x4d, 0x3d, 0xb7, 0xc5, 0xf9, 0xdc, 0x16, 0x6c, 0x72, 0x33, 0x7d, 0x05, 0xd1,
	0x30, 0xd0, 0x26, 0xa3, 0xdb, 0x3c, 0x47, 0x2c, 0x9c, 0x99, 0x88, 0x0d, 0xf0, 0xf2, 0xb7, 0x0f,
	0xc1, 0xde, 0xce, 0x22, 0x7b, 0x88, 0x35, 0x8a, 0xa2, 0x7b, 0x8a, 0x67, 0x53, 0x01, 0x47, 0x6e,
	0x5e, 0x4e, 0xc9, 0x19, 0xbb, 0xe9, 0x09, 0x79, 0x03, 0x91, 0x1d, 0xf3, 0xe3, 0x46, 0x1e, 0xe6,
	0xa7, 0xcc, 0x7a, 0x3f, 0x36, 0xee, 0x7f, 0x61, 0xfe, 0xb8, 0xc6, 0x0f, 0xb0, 0xb4, 0x8d, 0x5f,
	0xf1, 0x40, 0x9e, 0xcd, 0xf4, 0xd9, 0xf7, 0xbe, 0xb9, 0x98, 0xb7, 0x9e, 0x9e, 0x90, 0xf7, 0x00,
	0xb6, 0xfb, 0x13, 0x56, 0x68, 0x70, 0x5e, 0xf8, 0xbf, 0xcd, 0x3f, 0x43, 0xf7, 0x3f, 0xbf, 0xfe,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x7c, 0x1a, 0xc4, 0xe5, 0x03, 0x00, 0x00,
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
	SendDelete(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
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

func (c *emiroClient) SendDelete(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/emironetwork.Emiro/sendDelete", in, out, opts...)
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
	SendDelete(context.Context, *Query) (*Response, error)
}

// UnimplementedEmiroServer can be embedded to have forward compatible implementations.
type UnimplementedEmiroServer struct {
}

func (*UnimplementedEmiroServer) SendQuery(ctx context.Context, req *Query) (*ResponseQueryAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendQuery not implemented")
}
func (*UnimplementedEmiroServer) SendShow(ctx context.Context, req *Query) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendShow not implemented")
}
func (*UnimplementedEmiroServer) SendExec(ctx context.Context, req *Query) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendExec not implemented")
}
func (*UnimplementedEmiroServer) SendNew(ctx context.Context, req *QueryFull) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNew not implemented")
}
func (*UnimplementedEmiroServer) SendDelete(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDelete not implemented")
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

func _Emiro_SendDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmiroServer).SendDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emironetwork.Emiro/SendDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmiroServer).SendDelete(ctx, req.(*Query))
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
		{
			MethodName: "sendDelete",
			Handler:    _Emiro_SendDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emironetwork.proto",
}
