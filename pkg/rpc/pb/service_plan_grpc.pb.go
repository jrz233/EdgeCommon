// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_plan.proto

package pb

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
	PlanService_CreatePlan_FullMethodName                 = "/pb.PlanService/createPlan"
	PlanService_UpdatePlan_FullMethodName                 = "/pb.PlanService/updatePlan"
	PlanService_DeletePlan_FullMethodName                 = "/pb.PlanService/deletePlan"
	PlanService_FindEnabledPlan_FullMethodName            = "/pb.PlanService/findEnabledPlan"
	PlanService_FindBasicPlan_FullMethodName              = "/pb.PlanService/findBasicPlan"
	PlanService_CountAllEnabledPlans_FullMethodName       = "/pb.PlanService/countAllEnabledPlans"
	PlanService_ListEnabledPlans_FullMethodName           = "/pb.PlanService/listEnabledPlans"
	PlanService_FindAllAvailablePlans_FullMethodName      = "/pb.PlanService/findAllAvailablePlans"
	PlanService_FindAllAvailableBasicPlans_FullMethodName = "/pb.PlanService/findAllAvailableBasicPlans"
	PlanService_SortPlans_FullMethodName                  = "/pb.PlanService/sortPlans"
)

// PlanServiceClient is the client API for PlanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlanServiceClient interface {
	// 创建套餐
	CreatePlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*CreatePlanResponse, error)
	// 修改套餐
	UpdatePlan(ctx context.Context, in *UpdatePlanRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除套餐
	DeletePlan(ctx context.Context, in *DeletePlanRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个套餐
	FindEnabledPlan(ctx context.Context, in *FindEnabledPlanRequest, opts ...grpc.CallOption) (*FindEnabledPlanResponse, error)
	// 查找套餐基本信息
	FindBasicPlan(ctx context.Context, in *FindBasicPlanRequest, opts ...grpc.CallOption) (*FindBasicPlanResponse, error)
	// 计算套餐数量
	CountAllEnabledPlans(ctx context.Context, in *CountAllEnabledPlansRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页套餐
	ListEnabledPlans(ctx context.Context, in *ListEnabledPlansRequest, opts ...grpc.CallOption) (*ListEnabledPlansResponse, error)
	// 列出所有可用的套餐
	FindAllAvailablePlans(ctx context.Context, in *FindAllAvailablePlansRequest, opts ...grpc.CallOption) (*FindAllAvailablePlansResponse, error)
	// 列出所有可用的套餐的基本信息
	FindAllAvailableBasicPlans(ctx context.Context, in *FindAllAvailableBasicPlansRequest, opts ...grpc.CallOption) (*FindAllAvailableBasicPlansResponse, error)
	// 对套餐进行排序
	SortPlans(ctx context.Context, in *SortPlansRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type planServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlanServiceClient(cc grpc.ClientConnInterface) PlanServiceClient {
	return &planServiceClient{cc}
}

func (c *planServiceClient) CreatePlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*CreatePlanResponse, error) {
	out := new(CreatePlanResponse)
	err := c.cc.Invoke(ctx, PlanService_CreatePlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) UpdatePlan(ctx context.Context, in *UpdatePlanRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, PlanService_UpdatePlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) DeletePlan(ctx context.Context, in *DeletePlanRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, PlanService_DeletePlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) FindEnabledPlan(ctx context.Context, in *FindEnabledPlanRequest, opts ...grpc.CallOption) (*FindEnabledPlanResponse, error) {
	out := new(FindEnabledPlanResponse)
	err := c.cc.Invoke(ctx, PlanService_FindEnabledPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) FindBasicPlan(ctx context.Context, in *FindBasicPlanRequest, opts ...grpc.CallOption) (*FindBasicPlanResponse, error) {
	out := new(FindBasicPlanResponse)
	err := c.cc.Invoke(ctx, PlanService_FindBasicPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) CountAllEnabledPlans(ctx context.Context, in *CountAllEnabledPlansRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, PlanService_CountAllEnabledPlans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) ListEnabledPlans(ctx context.Context, in *ListEnabledPlansRequest, opts ...grpc.CallOption) (*ListEnabledPlansResponse, error) {
	out := new(ListEnabledPlansResponse)
	err := c.cc.Invoke(ctx, PlanService_ListEnabledPlans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) FindAllAvailablePlans(ctx context.Context, in *FindAllAvailablePlansRequest, opts ...grpc.CallOption) (*FindAllAvailablePlansResponse, error) {
	out := new(FindAllAvailablePlansResponse)
	err := c.cc.Invoke(ctx, PlanService_FindAllAvailablePlans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) FindAllAvailableBasicPlans(ctx context.Context, in *FindAllAvailableBasicPlansRequest, opts ...grpc.CallOption) (*FindAllAvailableBasicPlansResponse, error) {
	out := new(FindAllAvailableBasicPlansResponse)
	err := c.cc.Invoke(ctx, PlanService_FindAllAvailableBasicPlans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) SortPlans(ctx context.Context, in *SortPlansRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, PlanService_SortPlans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlanServiceServer is the server API for PlanService service.
// All implementations should embed UnimplementedPlanServiceServer
// for forward compatibility
type PlanServiceServer interface {
	// 创建套餐
	CreatePlan(context.Context, *CreatePlanRequest) (*CreatePlanResponse, error)
	// 修改套餐
	UpdatePlan(context.Context, *UpdatePlanRequest) (*RPCSuccess, error)
	// 删除套餐
	DeletePlan(context.Context, *DeletePlanRequest) (*RPCSuccess, error)
	// 查找单个套餐
	FindEnabledPlan(context.Context, *FindEnabledPlanRequest) (*FindEnabledPlanResponse, error)
	// 查找套餐基本信息
	FindBasicPlan(context.Context, *FindBasicPlanRequest) (*FindBasicPlanResponse, error)
	// 计算套餐数量
	CountAllEnabledPlans(context.Context, *CountAllEnabledPlansRequest) (*RPCCountResponse, error)
	// 列出单页套餐
	ListEnabledPlans(context.Context, *ListEnabledPlansRequest) (*ListEnabledPlansResponse, error)
	// 列出所有可用的套餐
	FindAllAvailablePlans(context.Context, *FindAllAvailablePlansRequest) (*FindAllAvailablePlansResponse, error)
	// 列出所有可用的套餐的基本信息
	FindAllAvailableBasicPlans(context.Context, *FindAllAvailableBasicPlansRequest) (*FindAllAvailableBasicPlansResponse, error)
	// 对套餐进行排序
	SortPlans(context.Context, *SortPlansRequest) (*RPCSuccess, error)
}

// UnimplementedPlanServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPlanServiceServer struct {
}

func (UnimplementedPlanServiceServer) CreatePlan(context.Context, *CreatePlanRequest) (*CreatePlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlan not implemented")
}
func (UnimplementedPlanServiceServer) UpdatePlan(context.Context, *UpdatePlanRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlan not implemented")
}
func (UnimplementedPlanServiceServer) DeletePlan(context.Context, *DeletePlanRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlan not implemented")
}
func (UnimplementedPlanServiceServer) FindEnabledPlan(context.Context, *FindEnabledPlanRequest) (*FindEnabledPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledPlan not implemented")
}
func (UnimplementedPlanServiceServer) FindBasicPlan(context.Context, *FindBasicPlanRequest) (*FindBasicPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBasicPlan not implemented")
}
func (UnimplementedPlanServiceServer) CountAllEnabledPlans(context.Context, *CountAllEnabledPlansRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledPlans not implemented")
}
func (UnimplementedPlanServiceServer) ListEnabledPlans(context.Context, *ListEnabledPlansRequest) (*ListEnabledPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledPlans not implemented")
}
func (UnimplementedPlanServiceServer) FindAllAvailablePlans(context.Context, *FindAllAvailablePlansRequest) (*FindAllAvailablePlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllAvailablePlans not implemented")
}
func (UnimplementedPlanServiceServer) FindAllAvailableBasicPlans(context.Context, *FindAllAvailableBasicPlansRequest) (*FindAllAvailableBasicPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllAvailableBasicPlans not implemented")
}
func (UnimplementedPlanServiceServer) SortPlans(context.Context, *SortPlansRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortPlans not implemented")
}

// UnsafePlanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlanServiceServer will
// result in compilation errors.
type UnsafePlanServiceServer interface {
	mustEmbedUnimplementedPlanServiceServer()
}

func RegisterPlanServiceServer(s grpc.ServiceRegistrar, srv PlanServiceServer) {
	s.RegisterService(&PlanService_ServiceDesc, srv)
}

func _PlanService_CreatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).CreatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_CreatePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).CreatePlan(ctx, req.(*CreatePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_UpdatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).UpdatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_UpdatePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).UpdatePlan(ctx, req.(*UpdatePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_DeletePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).DeletePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_DeletePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).DeletePlan(ctx, req.(*DeletePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_FindEnabledPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).FindEnabledPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_FindEnabledPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).FindEnabledPlan(ctx, req.(*FindEnabledPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_FindBasicPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBasicPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).FindBasicPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_FindBasicPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).FindBasicPlan(ctx, req.(*FindBasicPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_CountAllEnabledPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).CountAllEnabledPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_CountAllEnabledPlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).CountAllEnabledPlans(ctx, req.(*CountAllEnabledPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_ListEnabledPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).ListEnabledPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_ListEnabledPlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).ListEnabledPlans(ctx, req.(*ListEnabledPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_FindAllAvailablePlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllAvailablePlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).FindAllAvailablePlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_FindAllAvailablePlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).FindAllAvailablePlans(ctx, req.(*FindAllAvailablePlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_FindAllAvailableBasicPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllAvailableBasicPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).FindAllAvailableBasicPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_FindAllAvailableBasicPlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).FindAllAvailableBasicPlans(ctx, req.(*FindAllAvailableBasicPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_SortPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).SortPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_SortPlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).SortPlans(ctx, req.(*SortPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlanService_ServiceDesc is the grpc.ServiceDesc for PlanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PlanService",
	HandlerType: (*PlanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createPlan",
			Handler:    _PlanService_CreatePlan_Handler,
		},
		{
			MethodName: "updatePlan",
			Handler:    _PlanService_UpdatePlan_Handler,
		},
		{
			MethodName: "deletePlan",
			Handler:    _PlanService_DeletePlan_Handler,
		},
		{
			MethodName: "findEnabledPlan",
			Handler:    _PlanService_FindEnabledPlan_Handler,
		},
		{
			MethodName: "findBasicPlan",
			Handler:    _PlanService_FindBasicPlan_Handler,
		},
		{
			MethodName: "countAllEnabledPlans",
			Handler:    _PlanService_CountAllEnabledPlans_Handler,
		},
		{
			MethodName: "listEnabledPlans",
			Handler:    _PlanService_ListEnabledPlans_Handler,
		},
		{
			MethodName: "findAllAvailablePlans",
			Handler:    _PlanService_FindAllAvailablePlans_Handler,
		},
		{
			MethodName: "findAllAvailableBasicPlans",
			Handler:    _PlanService_FindAllAvailableBasicPlans_Handler,
		},
		{
			MethodName: "sortPlans",
			Handler:    _PlanService_SortPlans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_plan.proto",
}
