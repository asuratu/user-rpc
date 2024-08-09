package userlogic

import (
	"context"

	"github.com/asuratu/user-rpc/internal/svc"
	"github.com/asuratu/user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.DeleteUserRequest) (*user.DefaultResponse, error) {
	// todo: add your logic here and delete this line

	return &user.DefaultResponse{}, nil
}
