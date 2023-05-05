// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: services.proto

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
	PortfolioService_GetUser_FullMethodName        = "/pb.PortfolioService/GetUser"
	PortfolioService_UpdateUser_FullMethodName     = "/pb.PortfolioService/UpdateUser"
	PortfolioService_DeleteUser_FullMethodName     = "/pb.PortfolioService/DeleteUser"
	PortfolioService_CreateUser_FullMethodName     = "/pb.PortfolioService/CreateUser"
	PortfolioService_LoginUser_FullMethodName      = "/pb.PortfolioService/LoginUser"
	PortfolioService_LogoutUser_FullMethodName     = "/pb.PortfolioService/LogoutUser"
	PortfolioService_RefreshToken_FullMethodName   = "/pb.PortfolioService/RefreshToken"
	PortfolioService_ResetPassword_FullMethodName  = "/pb.PortfolioService/ResetPassword"
	PortfolioService_ChangePassword_FullMethodName = "/pb.PortfolioService/ChangePassword"
	PortfolioService_CreatePost_FullMethodName     = "/pb.PortfolioService/CreatePost"
	PortfolioService_GetPost_FullMethodName        = "/pb.PortfolioService/GetPost"
	PortfolioService_ListPosts_FullMethodName      = "/pb.PortfolioService/ListPosts"
	PortfolioService_CreateRecipe_FullMethodName   = "/pb.PortfolioService/CreateRecipe"
	PortfolioService_GetRecipe_FullMethodName      = "/pb.PortfolioService/GetRecipe"
	PortfolioService_UpdateRecipe_FullMethodName   = "/pb.PortfolioService/UpdateRecipe"
	PortfolioService_DeleteRecipe_FullMethodName   = "/pb.PortfolioService/DeleteRecipe"
	PortfolioService_ListRecipes_FullMethodName    = "/pb.PortfolioService/ListRecipes"
)

// PortfolioServiceClient is the client API for PortfolioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortfolioServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	// user create account
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	LogoutUser(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
	ListPosts(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error)
	CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error)
	GetRecipe(ctx context.Context, in *GetRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error)
	UpdateRecipe(ctx context.Context, in *UpdateRecipeRequest, opts ...grpc.CallOption) (*UpdateRecipeResponse, error)
	DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error)
	ListRecipes(ctx context.Context, in *ListRecipesRequest, opts ...grpc.CallOption) (*ListRecipesResponse, error)
}

type portfolioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortfolioServiceClient(cc grpc.ClientConnInterface) PortfolioServiceClient {
	return &portfolioServiceClient{cc}
}

func (c *portfolioServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, PortfolioService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, PortfolioService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, PortfolioService_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, PortfolioService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, PortfolioService_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) LogoutUser(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, PortfolioService_LogoutUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, PortfolioService_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, PortfolioService_ResetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, PortfolioService_ChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, PortfolioService_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	out := new(GetPostResponse)
	err := c.cc.Invoke(ctx, PortfolioService_GetPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) ListPosts(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error) {
	out := new(ListPostsResponse)
	err := c.cc.Invoke(ctx, PortfolioService_ListPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error) {
	out := new(CreateRecipeResponse)
	err := c.cc.Invoke(ctx, PortfolioService_CreateRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) GetRecipe(ctx context.Context, in *GetRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error) {
	out := new(GetRecipeResponse)
	err := c.cc.Invoke(ctx, PortfolioService_GetRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) UpdateRecipe(ctx context.Context, in *UpdateRecipeRequest, opts ...grpc.CallOption) (*UpdateRecipeResponse, error) {
	out := new(UpdateRecipeResponse)
	err := c.cc.Invoke(ctx, PortfolioService_UpdateRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error) {
	out := new(DeleteRecipeResponse)
	err := c.cc.Invoke(ctx, PortfolioService_DeleteRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) ListRecipes(ctx context.Context, in *ListRecipesRequest, opts ...grpc.CallOption) (*ListRecipesResponse, error) {
	out := new(ListRecipesResponse)
	err := c.cc.Invoke(ctx, PortfolioService_ListRecipes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortfolioServiceServer is the server API for PortfolioService service.
// All implementations must embed UnimplementedPortfolioServiceServer
// for forward compatibility
type PortfolioServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	// user create account
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	LogoutUser(context.Context, *LogoutRequest) (*LogoutResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error)
	ListPosts(context.Context, *ListPostsRequest) (*ListPostsResponse, error)
	CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error)
	GetRecipe(context.Context, *GetRecipeRequest) (*GetRecipeResponse, error)
	UpdateRecipe(context.Context, *UpdateRecipeRequest) (*UpdateRecipeResponse, error)
	DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error)
	ListRecipes(context.Context, *ListRecipesRequest) (*ListRecipesResponse, error)
	mustEmbedUnimplementedPortfolioServiceServer()
}

// UnimplementedPortfolioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortfolioServiceServer struct {
}

func (UnimplementedPortfolioServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedPortfolioServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedPortfolioServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedPortfolioServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedPortfolioServiceServer) LoginUser(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedPortfolioServiceServer) LogoutUser(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutUser not implemented")
}
func (UnimplementedPortfolioServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedPortfolioServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedPortfolioServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedPortfolioServiceServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPortfolioServiceServer) GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPortfolioServiceServer) ListPosts(context.Context, *ListPostsRequest) (*ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (UnimplementedPortfolioServiceServer) CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecipe not implemented")
}
func (UnimplementedPortfolioServiceServer) GetRecipe(context.Context, *GetRecipeRequest) (*GetRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}
func (UnimplementedPortfolioServiceServer) UpdateRecipe(context.Context, *UpdateRecipeRequest) (*UpdateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipe not implemented")
}
func (UnimplementedPortfolioServiceServer) DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipe not implemented")
}
func (UnimplementedPortfolioServiceServer) ListRecipes(context.Context, *ListRecipesRequest) (*ListRecipesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecipes not implemented")
}
func (UnimplementedPortfolioServiceServer) mustEmbedUnimplementedPortfolioServiceServer() {}

// UnsafePortfolioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortfolioServiceServer will
// result in compilation errors.
type UnsafePortfolioServiceServer interface {
	mustEmbedUnimplementedPortfolioServiceServer()
}

func RegisterPortfolioServiceServer(s grpc.ServiceRegistrar, srv PortfolioServiceServer) {
	s.RegisterService(&PortfolioService_ServiceDesc, srv)
}

func _PortfolioService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_LogoutUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).LogoutUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_LogoutUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).LogoutUser(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_ListPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).ListPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_ListPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).ListPosts(ctx, req.(*ListPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_CreateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).CreateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_CreateRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).CreateRecipe(ctx, req.(*CreateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_GetRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).GetRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_GetRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).GetRecipe(ctx, req.(*GetRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_UpdateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).UpdateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_UpdateRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).UpdateRecipe(ctx, req.(*UpdateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_DeleteRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).DeleteRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_DeleteRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).DeleteRecipe(ctx, req.(*DeleteRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_ListRecipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecipesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).ListRecipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_ListRecipes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).ListRecipes(ctx, req.(*ListRecipesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortfolioService_ServiceDesc is the grpc.ServiceDesc for PortfolioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortfolioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PortfolioService",
	HandlerType: (*PortfolioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _PortfolioService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _PortfolioService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _PortfolioService_DeleteUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _PortfolioService_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _PortfolioService_LoginUser_Handler,
		},
		{
			MethodName: "LogoutUser",
			Handler:    _PortfolioService_LogoutUser_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _PortfolioService_RefreshToken_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _PortfolioService_ResetPassword_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _PortfolioService_ChangePassword_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _PortfolioService_CreatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PortfolioService_GetPost_Handler,
		},
		{
			MethodName: "ListPosts",
			Handler:    _PortfolioService_ListPosts_Handler,
		},
		{
			MethodName: "CreateRecipe",
			Handler:    _PortfolioService_CreateRecipe_Handler,
		},
		{
			MethodName: "GetRecipe",
			Handler:    _PortfolioService_GetRecipe_Handler,
		},
		{
			MethodName: "UpdateRecipe",
			Handler:    _PortfolioService_UpdateRecipe_Handler,
		},
		{
			MethodName: "DeleteRecipe",
			Handler:    _PortfolioService_DeleteRecipe_Handler,
		},
		{
			MethodName: "ListRecipes",
			Handler:    _PortfolioService_ListRecipes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
