// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: crud_service.proto

package crud_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_crud_service_proto protoreflect.FileDescriptor

var file_crud_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xed, 0x01, 0x0a, 0x0b,
	0x43, 0x72, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x63,
	0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x12, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e,
	0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_crud_service_proto_goTypes = []interface{}{
	(*ListReq)(nil),       // 0: crud_service.ListReq
	(*ById)(nil),          // 1: crud_service.ById
	(*Post)(nil),          // 2: crud_service.Post
	(*ListRespPost)(nil),  // 3: crud_service.ListRespPost
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_crud_service_proto_depIdxs = []int32{
	0, // 0: crud_service.CrudService.GetAll:input_type -> crud_service.ListReq
	1, // 1: crud_service.CrudService.Get:input_type -> crud_service.ById
	2, // 2: crud_service.CrudService.Update:input_type -> crud_service.Post
	1, // 3: crud_service.CrudService.Delete:input_type -> crud_service.ById
	3, // 4: crud_service.CrudService.GetAll:output_type -> crud_service.ListRespPost
	2, // 5: crud_service.CrudService.Get:output_type -> crud_service.Post
	4, // 6: crud_service.CrudService.Update:output_type -> google.protobuf.Empty
	4, // 7: crud_service.CrudService.Delete:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crud_service_proto_init() }
func file_crud_service_proto_init() {
	if File_crud_service_proto != nil {
		return
	}
	file_crud_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crud_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crud_service_proto_goTypes,
		DependencyIndexes: file_crud_service_proto_depIdxs,
	}.Build()
	File_crud_service_proto = out.File
	file_crud_service_proto_rawDesc = nil
	file_crud_service_proto_goTypes = nil
	file_crud_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CrudServiceClient is the client API for CrudService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrudServiceClient interface {
	GetAll(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListRespPost, error)
	Get(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Post, error)
	Update(ctx context.Context, in *Post, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type crudServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrudServiceClient(cc grpc.ClientConnInterface) CrudServiceClient {
	return &crudServiceClient{cc}
}

func (c *crudServiceClient) GetAll(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListRespPost, error) {
	out := new(ListRespPost)
	err := c.cc.Invoke(ctx, "/crud_service.CrudService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) Get(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/crud_service.CrudService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) Update(ctx context.Context, in *Post, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/crud_service.CrudService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/crud_service.CrudService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrudServiceServer is the server API for CrudService service.
type CrudServiceServer interface {
	GetAll(context.Context, *ListReq) (*ListRespPost, error)
	Get(context.Context, *ById) (*Post, error)
	Update(context.Context, *Post) (*emptypb.Empty, error)
	Delete(context.Context, *ById) (*emptypb.Empty, error)
}

// UnimplementedCrudServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCrudServiceServer struct {
}

func (*UnimplementedCrudServiceServer) GetAll(context.Context, *ListReq) (*ListRespPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedCrudServiceServer) Get(context.Context, *ById) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCrudServiceServer) Update(context.Context, *Post) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCrudServiceServer) Delete(context.Context, *ById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCrudServiceServer(s *grpc.Server, srv CrudServiceServer) {
	s.RegisterService(&_CrudService_serviceDesc, srv)
}

func _CrudService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_service.CrudService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).GetAll(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_service.CrudService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).Get(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_service.CrudService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).Update(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_service.CrudService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

var _CrudService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crud_service.CrudService",
	HandlerType: (*CrudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _CrudService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CrudService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CrudService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CrudService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud_service.proto",
}
