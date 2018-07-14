// Code generated by protoc-gen-go. DO NOT EDIT.
// source: destination.proto

/*
Package destination is a generated protocol buffer package.

It is generated from these files:
	destination.proto

It has these top-level messages:
	GetDestination
	Update
	AddrSet
	WeightedAddrSet
	WeightedAddr
	TlsIdentity
	NoEndpoints
*/
package destination

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import io_linkerd_proxy_net "github.com/linkerd/linkerd2-proxy-api/go/net"

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

type GetDestination struct {
	Scheme string `protobuf:"bytes,1,opt,name=scheme" json:"scheme,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *GetDestination) Reset()                    { *m = GetDestination{} }
func (m *GetDestination) String() string            { return proto.CompactTextString(m) }
func (*GetDestination) ProtoMessage()               {}
func (*GetDestination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetDestination) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *GetDestination) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Update struct {
	// Types that are valid to be assigned to Update:
	//	*Update_Add
	//	*Update_Remove
	//	*Update_NoEndpoints
	Update isUpdate_Update `protobuf_oneof:"update"`
}

func (m *Update) Reset()                    { *m = Update{} }
func (m *Update) String() string            { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()               {}
func (*Update) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isUpdate_Update interface{ isUpdate_Update() }

type Update_Add struct {
	Add *WeightedAddrSet `protobuf:"bytes,1,opt,name=add,oneof"`
}
type Update_Remove struct {
	Remove *AddrSet `protobuf:"bytes,2,opt,name=remove,oneof"`
}
type Update_NoEndpoints struct {
	NoEndpoints *NoEndpoints `protobuf:"bytes,3,opt,name=no_endpoints,json=noEndpoints,oneof"`
}

func (*Update_Add) isUpdate_Update()         {}
func (*Update_Remove) isUpdate_Update()      {}
func (*Update_NoEndpoints) isUpdate_Update() {}

func (m *Update) GetUpdate() isUpdate_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *Update) GetAdd() *WeightedAddrSet {
	if x, ok := m.GetUpdate().(*Update_Add); ok {
		return x.Add
	}
	return nil
}

func (m *Update) GetRemove() *AddrSet {
	if x, ok := m.GetUpdate().(*Update_Remove); ok {
		return x.Remove
	}
	return nil
}

func (m *Update) GetNoEndpoints() *NoEndpoints {
	if x, ok := m.GetUpdate().(*Update_NoEndpoints); ok {
		return x.NoEndpoints
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Update) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Update_OneofMarshaler, _Update_OneofUnmarshaler, _Update_OneofSizer, []interface{}{
		(*Update_Add)(nil),
		(*Update_Remove)(nil),
		(*Update_NoEndpoints)(nil),
	}
}

func _Update_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Update)
	// update
	switch x := m.Update.(type) {
	case *Update_Add:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Add); err != nil {
			return err
		}
	case *Update_Remove:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Remove); err != nil {
			return err
		}
	case *Update_NoEndpoints:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NoEndpoints); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Update.Update has unexpected type %T", x)
	}
	return nil
}

func _Update_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Update)
	switch tag {
	case 1: // update.add
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WeightedAddrSet)
		err := b.DecodeMessage(msg)
		m.Update = &Update_Add{msg}
		return true, err
	case 2: // update.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AddrSet)
		err := b.DecodeMessage(msg)
		m.Update = &Update_Remove{msg}
		return true, err
	case 3: // update.no_endpoints
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NoEndpoints)
		err := b.DecodeMessage(msg)
		m.Update = &Update_NoEndpoints{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Update_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Update)
	// update
	switch x := m.Update.(type) {
	case *Update_Add:
		s := proto.Size(x.Add)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_Remove:
		s := proto.Size(x.Remove)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_NoEndpoints:
		s := proto.Size(x.NoEndpoints)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AddrSet struct {
	Addrs []*io_linkerd_proxy_net.TcpAddress `protobuf:"bytes,1,rep,name=addrs" json:"addrs,omitempty"`
}

func (m *AddrSet) Reset()                    { *m = AddrSet{} }
func (m *AddrSet) String() string            { return proto.CompactTextString(m) }
func (*AddrSet) ProtoMessage()               {}
func (*AddrSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddrSet) GetAddrs() []*io_linkerd_proxy_net.TcpAddress {
	if m != nil {
		return m.Addrs
	}
	return nil
}

type WeightedAddrSet struct {
	Addrs        []*WeightedAddr   `protobuf:"bytes,1,rep,name=addrs" json:"addrs,omitempty"`
	MetricLabels map[string]string `protobuf:"bytes,2,rep,name=metric_labels,json=metricLabels" json:"metric_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *WeightedAddrSet) Reset()                    { *m = WeightedAddrSet{} }
func (m *WeightedAddrSet) String() string            { return proto.CompactTextString(m) }
func (*WeightedAddrSet) ProtoMessage()               {}
func (*WeightedAddrSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WeightedAddrSet) GetAddrs() []*WeightedAddr {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func (m *WeightedAddrSet) GetMetricLabels() map[string]string {
	if m != nil {
		return m.MetricLabels
	}
	return nil
}

type WeightedAddr struct {
	Addr         *io_linkerd_proxy_net.TcpAddress `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Weight       uint32                           `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	MetricLabels map[string]string                `protobuf:"bytes,4,rep,name=metric_labels,json=metricLabels" json:"metric_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TlsIdentity  *TlsIdentity                     `protobuf:"bytes,5,opt,name=tls_identity,json=tlsIdentity" json:"tls_identity,omitempty"`
}

func (m *WeightedAddr) Reset()                    { *m = WeightedAddr{} }
func (m *WeightedAddr) String() string            { return proto.CompactTextString(m) }
func (*WeightedAddr) ProtoMessage()               {}
func (*WeightedAddr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WeightedAddr) GetAddr() *io_linkerd_proxy_net.TcpAddress {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *WeightedAddr) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *WeightedAddr) GetMetricLabels() map[string]string {
	if m != nil {
		return m.MetricLabels
	}
	return nil
}

func (m *WeightedAddr) GetTlsIdentity() *TlsIdentity {
	if m != nil {
		return m.TlsIdentity
	}
	return nil
}

// Which strategy should be used for verifying TLS.
type TlsIdentity struct {
	// Types that are valid to be assigned to Strategy:
	//	*TlsIdentity_K8SPodIdentity_
	Strategy isTlsIdentity_Strategy `protobuf_oneof:"strategy"`
}

func (m *TlsIdentity) Reset()                    { *m = TlsIdentity{} }
func (m *TlsIdentity) String() string            { return proto.CompactTextString(m) }
func (*TlsIdentity) ProtoMessage()               {}
func (*TlsIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isTlsIdentity_Strategy interface{ isTlsIdentity_Strategy() }

type TlsIdentity_K8SPodIdentity_ struct {
	K8SPodIdentity *TlsIdentity_K8SPodIdentity `protobuf:"bytes,2,opt,name=k8s_pod_identity,json=k8sPodIdentity,oneof"`
}

func (*TlsIdentity_K8SPodIdentity_) isTlsIdentity_Strategy() {}

func (m *TlsIdentity) GetStrategy() isTlsIdentity_Strategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (m *TlsIdentity) GetK8SPodIdentity() *TlsIdentity_K8SPodIdentity {
	if x, ok := m.GetStrategy().(*TlsIdentity_K8SPodIdentity_); ok {
		return x.K8SPodIdentity
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TlsIdentity) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TlsIdentity_OneofMarshaler, _TlsIdentity_OneofUnmarshaler, _TlsIdentity_OneofSizer, []interface{}{
		(*TlsIdentity_K8SPodIdentity_)(nil),
	}
}

func _TlsIdentity_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TlsIdentity)
	// strategy
	switch x := m.Strategy.(type) {
	case *TlsIdentity_K8SPodIdentity_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.K8SPodIdentity); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TlsIdentity.Strategy has unexpected type %T", x)
	}
	return nil
}

func _TlsIdentity_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TlsIdentity)
	switch tag {
	case 2: // strategy.k8s_pod_identity
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TlsIdentity_K8SPodIdentity)
		err := b.DecodeMessage(msg)
		m.Strategy = &TlsIdentity_K8SPodIdentity_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TlsIdentity_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TlsIdentity)
	// strategy
	switch x := m.Strategy.(type) {
	case *TlsIdentity_K8SPodIdentity_:
		s := proto.Size(x.K8SPodIdentity)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Verify the certificate based on the Kubernetes pod identity.
type TlsIdentity_K8SPodIdentity struct {
	// The pod_identity string is of the format:
	// {owner_name}.{owner_type}.{pod_ns}.linkerd-managed.{controller_ns}.svc.cluster.local
	PodIdentity string `protobuf:"bytes,1,opt,name=pod_identity,json=podIdentity" json:"pod_identity,omitempty"`
	// The Kubernetes namespace of the pod's Linkerd2 control plane.
	ControllerNs string `protobuf:"bytes,2,opt,name=controller_ns,json=controllerNs" json:"controller_ns,omitempty"`
}

func (m *TlsIdentity_K8SPodIdentity) Reset()                    { *m = TlsIdentity_K8SPodIdentity{} }
func (m *TlsIdentity_K8SPodIdentity) String() string            { return proto.CompactTextString(m) }
func (*TlsIdentity_K8SPodIdentity) ProtoMessage()               {}
func (*TlsIdentity_K8SPodIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *TlsIdentity_K8SPodIdentity) GetPodIdentity() string {
	if m != nil {
		return m.PodIdentity
	}
	return ""
}

func (m *TlsIdentity_K8SPodIdentity) GetControllerNs() string {
	if m != nil {
		return m.ControllerNs
	}
	return ""
}

type NoEndpoints struct {
	Exists bool `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
}

func (m *NoEndpoints) Reset()                    { *m = NoEndpoints{} }
func (m *NoEndpoints) String() string            { return proto.CompactTextString(m) }
func (*NoEndpoints) ProtoMessage()               {}
func (*NoEndpoints) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *NoEndpoints) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func init() {
	proto.RegisterType((*GetDestination)(nil), "io.linkerd.proxy.destination.GetDestination")
	proto.RegisterType((*Update)(nil), "io.linkerd.proxy.destination.Update")
	proto.RegisterType((*AddrSet)(nil), "io.linkerd.proxy.destination.AddrSet")
	proto.RegisterType((*WeightedAddrSet)(nil), "io.linkerd.proxy.destination.WeightedAddrSet")
	proto.RegisterType((*WeightedAddr)(nil), "io.linkerd.proxy.destination.WeightedAddr")
	proto.RegisterType((*TlsIdentity)(nil), "io.linkerd.proxy.destination.TlsIdentity")
	proto.RegisterType((*TlsIdentity_K8SPodIdentity)(nil), "io.linkerd.proxy.destination.TlsIdentity.K8sPodIdentity")
	proto.RegisterType((*NoEndpoints)(nil), "io.linkerd.proxy.destination.NoEndpoints")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Destination service

type DestinationClient interface {
	// Given a destination, return all addresses in that destination as a long-
	// running stream of updates.
	Get(ctx context.Context, in *GetDestination, opts ...grpc.CallOption) (Destination_GetClient, error)
}

type destinationClient struct {
	cc *grpc.ClientConn
}

func NewDestinationClient(cc *grpc.ClientConn) DestinationClient {
	return &destinationClient{cc}
}

func (c *destinationClient) Get(ctx context.Context, in *GetDestination, opts ...grpc.CallOption) (Destination_GetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Destination_serviceDesc.Streams[0], c.cc, "/io.linkerd.proxy.destination.Destination/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &destinationGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Destination_GetClient interface {
	Recv() (*Update, error)
	grpc.ClientStream
}

type destinationGetClient struct {
	grpc.ClientStream
}

func (x *destinationGetClient) Recv() (*Update, error) {
	m := new(Update)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Destination service

type DestinationServer interface {
	// Given a destination, return all addresses in that destination as a long-
	// running stream of updates.
	Get(*GetDestination, Destination_GetServer) error
}

func RegisterDestinationServer(s *grpc.Server, srv DestinationServer) {
	s.RegisterService(&_Destination_serviceDesc, srv)
}

func _Destination_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDestination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DestinationServer).Get(m, &destinationGetServer{stream})
}

type Destination_GetServer interface {
	Send(*Update) error
	grpc.ServerStream
}

type destinationGetServer struct {
	grpc.ServerStream
}

func (x *destinationGetServer) Send(m *Update) error {
	return x.ServerStream.SendMsg(m)
}

var _Destination_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.linkerd.proxy.destination.Destination",
	HandlerType: (*DestinationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _Destination_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "destination.proto",
}

func init() { proto.RegisterFile("destination.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0xcd, 0xe6, 0xcb, 0xf4, 0xee, 0x26, 0xb6, 0x83, 0xc8, 0x12, 0x7c, 0x88, 0xab, 0x85, 0x2a,
	0x66, 0x23, 0xb1, 0x94, 0x20, 0x85, 0x98, 0x60, 0x69, 0xc4, 0x1a, 0x64, 0xad, 0x28, 0x82, 0x84,
	0x4d, 0x66, 0x48, 0x86, 0x6c, 0x66, 0x96, 0x9d, 0x49, 0x6d, 0xfe, 0xa4, 0xef, 0xfe, 0x01, 0x7f,
	0x84, 0x4f, 0xb2, 0xb3, 0x13, 0x76, 0xd3, 0x42, 0x1a, 0xf1, 0x29, 0x7b, 0x6f, 0xce, 0x39, 0x77,
	0xee, 0x99, 0x3b, 0x17, 0x0e, 0x30, 0x11, 0x92, 0x32, 0x5f, 0x52, 0xce, 0xdc, 0x30, 0xe2, 0x92,
	0xa3, 0x47, 0x94, 0xbb, 0x01, 0x65, 0x73, 0x12, 0xe1, 0x38, 0x73, 0xbd, 0x72, 0x33, 0x98, 0xfa,
	0x1e, 0x23, 0x32, 0x01, 0x3a, 0xa7, 0x50, 0x3b, 0x27, 0xf2, 0x6d, 0xfa, 0x27, 0x7a, 0x08, 0x65,
	0x31, 0x99, 0x91, 0x05, 0xb1, 0x8d, 0x86, 0x71, 0xb4, 0xe7, 0xe9, 0x08, 0x21, 0x28, 0x86, 0xbe,
	0x9c, 0xd9, 0x79, 0x95, 0x55, 0xdf, 0xce, 0x6f, 0x03, 0xca, 0x9f, 0x43, 0xec, 0x4b, 0x82, 0x7a,
	0x50, 0xf0, 0x31, 0x56, 0x1c, 0xb3, 0xdd, 0x74, 0xb7, 0xd5, 0x77, 0xbf, 0x10, 0x3a, 0x9d, 0x49,
	0x82, 0x7b, 0x18, 0x47, 0x9f, 0x88, 0x1c, 0xe4, 0xbc, 0x98, 0x8b, 0xba, 0x50, 0x8e, 0xc8, 0x82,
	0x5f, 0x11, 0x55, 0xc3, 0x6c, 0x1f, 0x6e, 0x57, 0x49, 0xd9, 0x9a, 0x86, 0x86, 0x60, 0x31, 0x3e,
	0x22, 0x0c, 0x87, 0x9c, 0x32, 0x29, 0xec, 0x82, 0x92, 0x79, 0xb6, 0x5d, 0x66, 0xc8, 0xcf, 0xd6,
	0x84, 0x41, 0xce, 0x33, 0x59, 0x1a, 0xf6, 0x2b, 0x50, 0x5e, 0xaa, 0xee, 0x9c, 0x1e, 0xdc, 0xd3,
	0xe5, 0xd0, 0x09, 0x94, 0x7c, 0x8c, 0x23, 0x61, 0x1b, 0x8d, 0xc2, 0x91, 0xd9, 0x6e, 0xdc, 0x56,
	0x8f, 0xdd, 0xbd, 0x9c, 0x84, 0x31, 0x81, 0x08, 0xe1, 0x25, 0x70, 0xe7, 0x8f, 0x01, 0xf7, 0x6f,
	0x34, 0x8e, 0xde, 0x6c, 0x6a, 0x3d, 0xdf, 0xdd, 0x36, 0xad, 0x8a, 0x30, 0x54, 0x17, 0x44, 0x46,
	0x74, 0x32, 0x0a, 0xfc, 0x31, 0x09, 0x84, 0x9d, 0x57, 0x4a, 0xdd, 0x7f, 0xba, 0x00, 0xf7, 0x83,
	0x92, 0xb8, 0x50, 0x0a, 0x67, 0x4c, 0x46, 0x2b, 0xcf, 0x5a, 0x64, 0x52, 0xf5, 0x2e, 0x1c, 0xdc,
	0x82, 0xa0, 0x7d, 0x28, 0xcc, 0xc9, 0x4a, 0x4f, 0x49, 0xfc, 0x89, 0x1e, 0x40, 0xe9, 0xca, 0x0f,
	0x96, 0x44, 0xcf, 0x48, 0x12, 0xbc, 0xce, 0x77, 0x0c, 0xe7, 0x67, 0x1e, 0xac, 0x6c, 0x51, 0x74,
	0x0c, 0xc5, 0xb8, 0x01, 0x3d, 0x2f, 0x77, 0x9b, 0xa8, 0xd0, 0xf1, 0x6c, 0xfe, 0x50, 0x2a, 0xea,
	0x6a, 0xab, 0x9e, 0x8e, 0x90, 0x7f, 0xd3, 0x85, 0xa2, 0x72, 0xe1, 0x74, 0x77, 0x17, 0xee, 0xb2,
	0x00, 0x5d, 0x80, 0x25, 0x03, 0x31, 0xa2, 0x98, 0x30, 0x49, 0xe5, 0xca, 0x2e, 0xed, 0x32, 0x5b,
	0x97, 0x81, 0x78, 0xa7, 0x09, 0x9e, 0x29, 0xd3, 0xe0, 0xff, 0x0d, 0xfd, 0x65, 0x80, 0x99, 0x51,
	0x47, 0x18, 0xf6, 0xe7, 0x1d, 0x31, 0x0a, 0x39, 0x4e, 0x8f, 0x98, 0xbc, 0xa2, 0xce, 0xce, 0x47,
	0x74, 0xdf, 0x77, 0xc4, 0x47, 0x8e, 0xd7, 0xe1, 0x20, 0xe7, 0xd5, 0xe6, 0x1b, 0x99, 0xfa, 0x57,
	0xa8, 0x6d, 0x62, 0xd0, 0x63, 0xb0, 0x36, 0x6a, 0x26, 0x87, 0x37, 0xc3, 0x0c, 0xe4, 0x09, 0x54,
	0x27, 0x9c, 0xc9, 0x88, 0x07, 0x01, 0x89, 0x46, 0x4c, 0xe8, 0x66, 0xac, 0x34, 0x39, 0x14, 0x7d,
	0x80, 0x8a, 0x90, 0x91, 0x2f, 0xc9, 0x74, 0xe5, 0x1c, 0x82, 0x99, 0x79, 0x94, 0xf1, 0xa5, 0x93,
	0x6b, 0x2a, 0xa4, 0x50, 0xe2, 0x15, 0x4f, 0x47, 0xed, 0x00, 0xcc, 0xec, 0xde, 0xfa, 0x0e, 0x85,
	0x73, 0x22, 0xd1, 0x8b, 0xed, 0xed, 0x6e, 0x2e, 0xbb, 0xfa, 0xd3, 0xed, 0xe8, 0x64, 0xb7, 0x39,
	0xb9, 0x97, 0x46, 0xff, 0xe4, 0xdb, 0xf1, 0x94, 0xca, 0xd9, 0x72, 0xec, 0x4e, 0xf8, 0xa2, 0xa5,
	0x29, 0xeb, 0xdf, 0x76, 0x53, 0x71, 0x9b, 0x7e, 0x48, 0x5b, 0x53, 0xde, 0xca, 0x48, 0x8c, 0xcb,
	0x6a, 0xcf, 0xbe, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xf6, 0x78, 0xfb, 0xa5, 0x05, 0x00,
	0x00,
}
