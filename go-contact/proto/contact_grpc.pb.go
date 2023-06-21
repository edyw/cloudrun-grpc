// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: contact.proto

package contactpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Contact_GetContact_FullMethodName = "/contact.Contact/GetContact"
)

// ContactClient is the client API for Contact service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactClient interface {
	GetContact(ctx context.Context, in *ContactRequest, opts ...grpc.CallOption) (*ContactReply, error)
}

type contactClient struct {
	cc grpc.ClientConnInterface
}

func NewContactClient(cc grpc.ClientConnInterface) ContactClient {
	return &contactClient{cc}
}

func (c *contactClient) GetContact(ctx context.Context, in *ContactRequest, opts ...grpc.CallOption) (*ContactReply, error) {
	out := new(ContactReply)
	err := c.cc.Invoke(ctx, Contact_GetContact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServer is the server API for Contact service.
// All implementations must embed UnimplementedContactServer
// for forward compatibility
type ContactServer interface {
	GetContact(context.Context, *ContactRequest) (*ContactReply, error)
	mustEmbedUnimplementedContactServer()
}

// UnimplementedContactServer must be embedded to have forward compatible implementations.
type UnimplementedContactServer struct {
}

func (UnimplementedContactServer) GetContact(context.Context, *ContactRequest) (*ContactReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContact not implemented")
}
func (UnimplementedContactServer) mustEmbedUnimplementedContactServer() {}

// UnsafeContactServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactServer will
// result in compilation errors.
type UnsafeContactServer interface {
	mustEmbedUnimplementedContactServer()
}

func RegisterContactServer(s grpc.ServiceRegistrar, srv ContactServer) {
	s.RegisterService(&Contact_ServiceDesc, srv)
}

func _Contact_GetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServer).GetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contact_GetContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServer).GetContact(ctx, req.(*ContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Contact_ServiceDesc is the grpc.ServiceDesc for Contact service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Contact_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contact.Contact",
	HandlerType: (*ContactServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContact",
			Handler:    _Contact_GetContact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact.proto",
}
