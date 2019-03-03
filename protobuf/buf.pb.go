package buf

import (
	"fmt"
	"math"

	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ context.Context
var _ grpc.ClientConn

const _ = proto.ProtoPackageIsVersion3
const _ = grpc.SupportPackageIsVersion4

func init() {
	//proto.RegisterFile("buf.proto", fileDescriptor_buf_71e208cbdc16936b)
	proto.RegisterType((*ObjetoEntrada)(nil), "TestProto.ObjetoEntrada")
	proto.RegisterType((*ObjetoSaida)(nil), "TestProto.ObjetoSaida")
}

// ObjetoEntrada : The request message containing the user's name.
type ObjetoEntrada struct {
	Nome                 string   `protobuf:"bytes,1,opt,name=nome,proto3" json:"nome,omitempty"`
	Idade                string   `protobuf:"bytes,2,opt,name=idade,proto3" json:"idade,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type ObjetoSaida struct {
	Nome                 string   `protobuf:"bytes,1,opt,name=nome,proto3" json:"nome,omitempty"`
	Idade                string   `protobuf:"bytes,2,opt,name=idade,proto3" json:"idade,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjetoEntrada) Reset()         { *m = ObjetoEntrada{} }
func (m *ObjetoEntrada) String() string { return proto.CompactTextString(m) }

func (*ObjetoEntrada) ProtoMessage() {}

func (m *ObjetoEntrada) GetNome() string {
	if m != nil {
		return m.Nome
	}
	return ""
}
func (m *ObjetoEntrada) GetIdade() string {
	if m != nil {
		return m.Idade
	}
	return "0"
}

func (m *ObjetoSaida) Reset() { *m = ObjetoSaida{} }

func (m *ObjetoSaida) String() string { return proto.CompactTextString(m) }

func (*ObjetoSaida) ProtoMessage() {}

func (m *ObjetoSaida) GetNome() string {
	if m != nil {
		return m.Nome
	}
	return ""
}

func (m *ObjetoSaida) GetIdade() string {
	if m != nil {
		return m.Idade
	}
	return "0"
}

// SERVER

type MarotoServer interface {
	TesteMaroto(context.Context, *ObjetoEntrada) (*ObjetoSaida, error)
}

func RegisterMarotoServer(s *grpc.Server, srv MarotoServer) {
	s.RegisterService(&_Maroto_serviceDesc, srv)
}

var _Maroto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TestProto.Maroto",
	HandlerType: (*MarotoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TesteMaroto",
			Handler:    _Maroto_TesteMaroto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "TestProto.proto",
}

func _Maroto_TesteMaroto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjetoEntrada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarotoServer).TesteMaroto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestProto.Maroto/TesteMaroto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarotoServer).TesteMaroto(ctx, req.(*ObjetoEntrada))
	}
	return interceptor(ctx, in, info, handler)
}

// cliente

type MarotoClient interface {
	TesteMaroto(ctx context.Context, in *ObjetoEntrada, opts ...grpc.CallOption) (*ObjetoSaida, error)
}

type marotoClient struct {
	cc *grpc.ClientConn
}

func NewMarotoClient(cc *grpc.ClientConn) MarotoClient {
	return &marotoClient{cc}
}

func (c *marotoClient) TesteMaroto(ctx context.Context, in *ObjetoEntrada, opts ...grpc.CallOption) (*ObjetoSaida, error) {
	out := new(ObjetoSaida)
	err := c.cc.Invoke(ctx, "/TestProto.Maroto/TesteMaroto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
