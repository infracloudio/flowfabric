// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flowfabric.proto

package proto

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

// CaptureRequest containing network capture request
type CaptureRequest struct {
	Pod                  string   `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CaptureRequest) Reset()         { *m = CaptureRequest{} }
func (m *CaptureRequest) String() string { return proto.CompactTextString(m) }
func (*CaptureRequest) ProtoMessage()    {}
func (*CaptureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f228048f95a44b08, []int{0}
}

func (m *CaptureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaptureRequest.Unmarshal(m, b)
}
func (m *CaptureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaptureRequest.Marshal(b, m, deterministic)
}
func (m *CaptureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaptureRequest.Merge(m, src)
}
func (m *CaptureRequest) XXX_Size() int {
	return xxx_messageInfo_CaptureRequest.Size(m)
}
func (m *CaptureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CaptureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CaptureRequest proto.InternalMessageInfo

func (m *CaptureRequest) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

// CaputreResponse containing network details
type CaptureResponse struct {
	Src                  string   `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	Dst                  string   `protobuf:"bytes,2,opt,name=dst,proto3" json:"dst,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CaptureResponse) Reset()         { *m = CaptureResponse{} }
func (m *CaptureResponse) String() string { return proto.CompactTextString(m) }
func (*CaptureResponse) ProtoMessage()    {}
func (*CaptureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f228048f95a44b08, []int{1}
}

func (m *CaptureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaptureResponse.Unmarshal(m, b)
}
func (m *CaptureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaptureResponse.Marshal(b, m, deterministic)
}
func (m *CaptureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaptureResponse.Merge(m, src)
}
func (m *CaptureResponse) XXX_Size() int {
	return xxx_messageInfo_CaptureResponse.Size(m)
}
func (m *CaptureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CaptureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CaptureResponse proto.InternalMessageInfo

func (m *CaptureResponse) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *CaptureResponse) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

func init() {
	proto.RegisterType((*CaptureRequest)(nil), "proto.CaptureRequest")
	proto.RegisterType((*CaptureResponse)(nil), "proto.CaptureResponse")
}

func init() { proto.RegisterFile("flowfabric.proto", fileDescriptor_f228048f95a44b08) }

var fileDescriptor_f228048f95a44b08 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcb, 0xc9, 0x2f,
	0x4f, 0x4b, 0x4c, 0x2a, 0xca, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53,
	0x4a, 0x4a, 0x5c, 0x7c, 0xce, 0x89, 0x05, 0x25, 0xa5, 0x45, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0x05, 0xf9, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x20, 0xa6, 0x92, 0x29, 0x17, 0x3f, 0x5c, 0x4d, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x48, 0x51,
	0x71, 0x51, 0x32, 0x4c, 0x51, 0x71, 0x51, 0x32, 0x48, 0x24, 0xa5, 0xb8, 0x44, 0x82, 0x09, 0x22,
	0x92, 0x52, 0x5c, 0x62, 0xe4, 0xc7, 0xc5, 0xe7, 0x97, 0x5a, 0x52, 0x9e, 0x5f, 0x94, 0x0d, 0xd5,
	0x2d, 0x64, 0xc3, 0xc5, 0x0e, 0x63, 0x8a, 0x42, 0x9c, 0xa1, 0x87, 0x6a, 0xb9, 0x94, 0x18, 0xba,
	0x30, 0xc4, 0x3e, 0x25, 0x06, 0x03, 0xc6, 0x24, 0x36, 0xb0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xe2, 0x08, 0xc1, 0x89, 0xcc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkCaptureClient is the client API for NetworkCapture service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkCaptureClient interface {
	// Capture captures network
	Capture(ctx context.Context, in *CaptureRequest, opts ...grpc.CallOption) (NetworkCapture_CaptureClient, error)
}

type networkCaptureClient struct {
	cc *grpc.ClientConn
}

func NewNetworkCaptureClient(cc *grpc.ClientConn) NetworkCaptureClient {
	return &networkCaptureClient{cc}
}

func (c *networkCaptureClient) Capture(ctx context.Context, in *CaptureRequest, opts ...grpc.CallOption) (NetworkCapture_CaptureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkCapture_serviceDesc.Streams[0], "/proto.NetworkCapture/Capture", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkCaptureCaptureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetworkCapture_CaptureClient interface {
	Recv() (*CaptureResponse, error)
	grpc.ClientStream
}

type networkCaptureCaptureClient struct {
	grpc.ClientStream
}

func (x *networkCaptureCaptureClient) Recv() (*CaptureResponse, error) {
	m := new(CaptureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetworkCaptureServer is the server API for NetworkCapture service.
type NetworkCaptureServer interface {
	// Capture captures network
	Capture(*CaptureRequest, NetworkCapture_CaptureServer) error
}

// UnimplementedNetworkCaptureServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkCaptureServer struct {
}

func (*UnimplementedNetworkCaptureServer) Capture(req *CaptureRequest, srv NetworkCapture_CaptureServer) error {
	return status.Errorf(codes.Unimplemented, "method Capture not implemented")
}

func RegisterNetworkCaptureServer(s *grpc.Server, srv NetworkCaptureServer) {
	s.RegisterService(&_NetworkCapture_serviceDesc, srv)
}

func _NetworkCapture_Capture_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CaptureRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetworkCaptureServer).Capture(m, &networkCaptureCaptureServer{stream})
}

type NetworkCapture_CaptureServer interface {
	Send(*CaptureResponse) error
	grpc.ServerStream
}

type networkCaptureCaptureServer struct {
	grpc.ServerStream
}

func (x *networkCaptureCaptureServer) Send(m *CaptureResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NetworkCapture_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NetworkCapture",
	HandlerType: (*NetworkCaptureServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Capture",
			Handler:       _NetworkCapture_Capture_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "flowfabric.proto",
}