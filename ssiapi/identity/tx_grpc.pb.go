// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: identity/identity/tx.proto

package identity

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
	Msg_UpdateParams_FullMethodName  = "/identity.identity.Msg/UpdateParams"
	Msg_CreateUser_FullMethodName    = "/identity.identity.Msg/CreateUser"
	Msg_UpdateUser_FullMethodName    = "/identity.identity.Msg/UpdateUser"
	Msg_DeleteUser_FullMethodName    = "/identity.identity.Msg/DeleteUser"
	Msg_CreateAddress_FullMethodName = "/identity.identity.Msg/CreateAddress"
	Msg_UpdateAddress_FullMethodName = "/identity.identity.Msg/UpdateAddress"
	Msg_DeleteAddress_FullMethodName = "/identity.identity.Msg/DeleteAddress"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateUser(ctx context.Context, in *MsgCreateUser, opts ...grpc.CallOption) (*MsgCreateUserResponse, error)
	UpdateUser(ctx context.Context, in *MsgUpdateUser, opts ...grpc.CallOption) (*MsgUpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *MsgDeleteUser, opts ...grpc.CallOption) (*MsgDeleteUserResponse, error)
	CreateAddress(ctx context.Context, in *MsgCreateAddress, opts ...grpc.CallOption) (*MsgCreateAddressResponse, error)
	UpdateAddress(ctx context.Context, in *MsgUpdateAddress, opts ...grpc.CallOption) (*MsgUpdateAddressResponse, error)
	DeleteAddress(ctx context.Context, in *MsgDeleteAddress, opts ...grpc.CallOption) (*MsgDeleteAddressResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateUser(ctx context.Context, in *MsgCreateUser, opts ...grpc.CallOption) (*MsgCreateUserResponse, error) {
	out := new(MsgCreateUserResponse)
	err := c.cc.Invoke(ctx, Msg_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateUser(ctx context.Context, in *MsgUpdateUser, opts ...grpc.CallOption) (*MsgUpdateUserResponse, error) {
	out := new(MsgUpdateUserResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteUser(ctx context.Context, in *MsgDeleteUser, opts ...grpc.CallOption) (*MsgDeleteUserResponse, error) {
	out := new(MsgDeleteUserResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateAddress(ctx context.Context, in *MsgCreateAddress, opts ...grpc.CallOption) (*MsgCreateAddressResponse, error) {
	out := new(MsgCreateAddressResponse)
	err := c.cc.Invoke(ctx, Msg_CreateAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateAddress(ctx context.Context, in *MsgUpdateAddress, opts ...grpc.CallOption) (*MsgUpdateAddressResponse, error) {
	out := new(MsgUpdateAddressResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteAddress(ctx context.Context, in *MsgDeleteAddress, opts ...grpc.CallOption) (*MsgDeleteAddressResponse, error) {
	out := new(MsgDeleteAddressResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateUser(context.Context, *MsgCreateUser) (*MsgCreateUserResponse, error)
	UpdateUser(context.Context, *MsgUpdateUser) (*MsgUpdateUserResponse, error)
	DeleteUser(context.Context, *MsgDeleteUser) (*MsgDeleteUserResponse, error)
	CreateAddress(context.Context, *MsgCreateAddress) (*MsgCreateAddressResponse, error)
	UpdateAddress(context.Context, *MsgUpdateAddress) (*MsgUpdateAddressResponse, error)
	DeleteAddress(context.Context, *MsgDeleteAddress) (*MsgDeleteAddressResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateUser(context.Context, *MsgCreateUser) (*MsgCreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMsgServer) UpdateUser(context.Context, *MsgUpdateUser) (*MsgUpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedMsgServer) DeleteUser(context.Context, *MsgDeleteUser) (*MsgDeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedMsgServer) CreateAddress(context.Context, *MsgCreateAddress) (*MsgCreateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
}
func (UnimplementedMsgServer) UpdateAddress(context.Context, *MsgUpdateAddress) (*MsgUpdateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}
func (UnimplementedMsgServer) DeleteAddress(context.Context, *MsgDeleteAddress) (*MsgDeleteAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddress not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateUser(ctx, req.(*MsgCreateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateUser(ctx, req.(*MsgUpdateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteUser(ctx, req.(*MsgDeleteUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAddress(ctx, req.(*MsgCreateAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateAddress(ctx, req.(*MsgUpdateAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteAddress(ctx, req.(*MsgDeleteAddress))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identity.identity.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Msg_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Msg_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Msg_DeleteUser_Handler,
		},
		{
			MethodName: "CreateAddress",
			Handler:    _Msg_CreateAddress_Handler,
		},
		{
			MethodName: "UpdateAddress",
			Handler:    _Msg_UpdateAddress_Handler,
		},
		{
			MethodName: "DeleteAddress",
			Handler:    _Msg_DeleteAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/identity/tx.proto",
}