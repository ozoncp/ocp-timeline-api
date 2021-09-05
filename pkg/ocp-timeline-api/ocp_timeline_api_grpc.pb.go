// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_timeline_api

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

// OcpTimelineApiClient is the client API for OcpTimelineApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpTimelineApiClient interface {
	//CreateTimelineV1 adds a Timeline
	CreateTimelineV1(ctx context.Context, in *CreateTimelineV1Request, opts ...grpc.CallOption) (*CreateTimelineV1Response, error)
	GetTimelineV1(ctx context.Context, in *GetTimelineV1Request, opts ...grpc.CallOption) (*GetTimelineV1Response, error)
	ListTimelineV1(ctx context.Context, in *ListTimelineV1Request, opts ...grpc.CallOption) (*ListTimelineV1Response, error)
	RemoveTimelineV1(ctx context.Context, in *RemoveTimelineV1Request, opts ...grpc.CallOption) (*RemoveTimelineV1Response, error)
	UpdateTimelineV1(ctx context.Context, in *UpdateTimelineV1Request, opts ...grpc.CallOption) (*UpdateTimelineV1Response, error)
	MultiCreateTimelinesV1(ctx context.Context, in *MultiCreateTimelinesV1Request, opts ...grpc.CallOption) (*MultiCreateTimelinesV1Response, error)
}

type ocpTimelineApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpTimelineApiClient(cc grpc.ClientConnInterface) OcpTimelineApiClient {
	return &ocpTimelineApiClient{cc}
}

func (c *ocpTimelineApiClient) CreateTimelineV1(ctx context.Context, in *CreateTimelineV1Request, opts ...grpc.CallOption) (*CreateTimelineV1Response, error) {
	out := new(CreateTimelineV1Response)
	err := c.cc.Invoke(ctx, "/ocp.timeline.api.OcpTimelineApi/CreateTimelineV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpTimelineApiClient) GetTimelineV1(ctx context.Context, in *GetTimelineV1Request, opts ...grpc.CallOption) (*GetTimelineV1Response, error) {
	out := new(GetTimelineV1Response)
	err := c.cc.Invoke(ctx, "/ocp.timeline.api.OcpTimelineApi/GetTimelineV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpTimelineApiClient) ListTimelineV1(ctx context.Context, in *ListTimelineV1Request, opts ...grpc.CallOption) (*ListTimelineV1Response, error) {
	out := new(ListTimelineV1Response)
	err := c.cc.Invoke(ctx, "/ocp.timeline.api.OcpTimelineApi/ListTimelineV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpTimelineApiClient) RemoveTimelineV1(ctx context.Context, in *RemoveTimelineV1Request, opts ...grpc.CallOption) (*RemoveTimelineV1Response, error) {
	out := new(RemoveTimelineV1Response)
	err := c.cc.Invoke(ctx, "/ocp.timeline.api.OcpTimelineApi/RemoveTimelineV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpTimelineApiClient) UpdateTimelineV1(ctx context.Context, in *UpdateTimelineV1Request, opts ...grpc.CallOption) (*UpdateTimelineV1Response, error) {
	out := new(UpdateTimelineV1Response)
	err := c.cc.Invoke(ctx, "/ocp.timeline.api.OcpTimelineApi/UpdateTimelineV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpTimelineApiClient) MultiCreateTimelinesV1(ctx context.Context, in *MultiCreateTimelinesV1Request, opts ...grpc.CallOption) (*MultiCreateTimelinesV1Response, error) {
	out := new(MultiCreateTimelinesV1Response)
	err := c.cc.Invoke(ctx, "/ocp.timeline.api.OcpTimelineApi/MultiCreateTimelinesV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpTimelineApiServer is the server API for OcpTimelineApi service.
// All implementations must embed UnimplementedOcpTimelineApiServer
// for forward compatibility
type OcpTimelineApiServer interface {
	//CreateTimelineV1 adds a Timeline
	CreateTimelineV1(context.Context, *CreateTimelineV1Request) (*CreateTimelineV1Response, error)
	GetTimelineV1(context.Context, *GetTimelineV1Request) (*GetTimelineV1Response, error)
	ListTimelineV1(context.Context, *ListTimelineV1Request) (*ListTimelineV1Response, error)
	RemoveTimelineV1(context.Context, *RemoveTimelineV1Request) (*RemoveTimelineV1Response, error)
	UpdateTimelineV1(context.Context, *UpdateTimelineV1Request) (*UpdateTimelineV1Response, error)
	MultiCreateTimelinesV1(context.Context, *MultiCreateTimelinesV1Request) (*MultiCreateTimelinesV1Response, error)
	mustEmbedUnimplementedOcpTimelineApiServer()
}

// UnimplementedOcpTimelineApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpTimelineApiServer struct {
}

func (UnimplementedOcpTimelineApiServer) CreateTimelineV1(context.Context, *CreateTimelineV1Request) (*CreateTimelineV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTimelineV1 not implemented")
}
func (UnimplementedOcpTimelineApiServer) GetTimelineV1(context.Context, *GetTimelineV1Request) (*GetTimelineV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimelineV1 not implemented")
}
func (UnimplementedOcpTimelineApiServer) ListTimelineV1(context.Context, *ListTimelineV1Request) (*ListTimelineV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTimelineV1 not implemented")
}
func (UnimplementedOcpTimelineApiServer) RemoveTimelineV1(context.Context, *RemoveTimelineV1Request) (*RemoveTimelineV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTimelineV1 not implemented")
}
func (UnimplementedOcpTimelineApiServer) UpdateTimelineV1(context.Context, *UpdateTimelineV1Request) (*UpdateTimelineV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTimelineV1 not implemented")
}
func (UnimplementedOcpTimelineApiServer) MultiCreateTimelinesV1(context.Context, *MultiCreateTimelinesV1Request) (*MultiCreateTimelinesV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateTimelinesV1 not implemented")
}
func (UnimplementedOcpTimelineApiServer) mustEmbedUnimplementedOcpTimelineApiServer() {}

// UnsafeOcpTimelineApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpTimelineApiServer will
// result in compilation errors.
type UnsafeOcpTimelineApiServer interface {
	mustEmbedUnimplementedOcpTimelineApiServer()
}

func RegisterOcpTimelineApiServer(s grpc.ServiceRegistrar, srv OcpTimelineApiServer) {
	s.RegisterService(&OcpTimelineApi_ServiceDesc, srv)
}

func _OcpTimelineApi_CreateTimelineV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTimelineV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpTimelineApiServer).CreateTimelineV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.timeline.api.OcpTimelineApi/CreateTimelineV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpTimelineApiServer).CreateTimelineV1(ctx, req.(*CreateTimelineV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpTimelineApi_GetTimelineV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimelineV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpTimelineApiServer).GetTimelineV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.timeline.api.OcpTimelineApi/GetTimelineV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpTimelineApiServer).GetTimelineV1(ctx, req.(*GetTimelineV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpTimelineApi_ListTimelineV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTimelineV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpTimelineApiServer).ListTimelineV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.timeline.api.OcpTimelineApi/ListTimelineV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpTimelineApiServer).ListTimelineV1(ctx, req.(*ListTimelineV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpTimelineApi_RemoveTimelineV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTimelineV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpTimelineApiServer).RemoveTimelineV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.timeline.api.OcpTimelineApi/RemoveTimelineV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpTimelineApiServer).RemoveTimelineV1(ctx, req.(*RemoveTimelineV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpTimelineApi_UpdateTimelineV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTimelineV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpTimelineApiServer).UpdateTimelineV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.timeline.api.OcpTimelineApi/UpdateTimelineV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpTimelineApiServer).UpdateTimelineV1(ctx, req.(*UpdateTimelineV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpTimelineApi_MultiCreateTimelinesV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateTimelinesV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpTimelineApiServer).MultiCreateTimelinesV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.timeline.api.OcpTimelineApi/MultiCreateTimelinesV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpTimelineApiServer).MultiCreateTimelinesV1(ctx, req.(*MultiCreateTimelinesV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpTimelineApi_ServiceDesc is the grpc.ServiceDesc for OcpTimelineApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpTimelineApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.timeline.api.OcpTimelineApi",
	HandlerType: (*OcpTimelineApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTimelineV1",
			Handler:    _OcpTimelineApi_CreateTimelineV1_Handler,
		},
		{
			MethodName: "GetTimelineV1",
			Handler:    _OcpTimelineApi_GetTimelineV1_Handler,
		},
		{
			MethodName: "ListTimelineV1",
			Handler:    _OcpTimelineApi_ListTimelineV1_Handler,
		},
		{
			MethodName: "RemoveTimelineV1",
			Handler:    _OcpTimelineApi_RemoveTimelineV1_Handler,
		},
		{
			MethodName: "UpdateTimelineV1",
			Handler:    _OcpTimelineApi_UpdateTimelineV1_Handler,
		},
		{
			MethodName: "MultiCreateTimelinesV1",
			Handler:    _OcpTimelineApi_MultiCreateTimelinesV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-timeline-api/ocp_timeline_api.proto",
}
