// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"github.com/asuratu/user-rpc/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateUserRequest  = user.CreateUserRequest
	CreateUserResponse = user.CreateUserResponse
	DefaultResponse    = user.DefaultResponse
	DeleteUserRequest  = user.DeleteUserRequest
	GetUserRequest     = user.GetUserRequest
	GetUserResponse    = user.GetUserResponse
	Request            = user.Request
	Response           = user.Response
	UpdateUserRequest  = user.UpdateUserRequest

	User interface {
		GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
		CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
		UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
		DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUser) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}
