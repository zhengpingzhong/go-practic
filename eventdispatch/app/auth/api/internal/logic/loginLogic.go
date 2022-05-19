package logic

import (
	"context"
	"eventdispatch/app/auth/rpc/user"
	"github.com/jinzhu/copier"

	"eventdispatch/app/auth/api/internal/svc"
	"eventdispatch/app/auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginRes, error) {
	var userName string = ""
	var password string = ""
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Username: userName,
		Password: password,
		Client:   req.Client,
	})
	if err != nil {
		return nil, err
	}
	var resp types.LoginRes
	_ = copier.Copy(&resp, loginResp)
	return &resp, nil
}
