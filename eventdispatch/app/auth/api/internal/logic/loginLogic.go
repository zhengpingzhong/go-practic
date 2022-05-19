package logic

import (
	"context"
	"encoding/base64"
	"github.com/jinzhu/copier"
	"strings"

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
	var userName string = req.Username
	var password string = req.Password
	decodeString, err := base64.StdEncoding.DecodeString(userName)
	if err != nil {
		return nil, err
	}
	userName = string(decodeString)
	sysuser, err := l.svcCtx.UserModel.FindByUserName(l.ctx, userName)
	if !strings.EqualFold(password, sysuser.Password.String) {
		return &types.LoginRes{
			Code: "200",
			Msg:  "密码错误！",
		}, nil
	}
	if err != nil {
		return nil, err
	}
	var resp types.LoginRes
	_ = copier.Copy(&resp, sysuser)
	return &resp, nil
}
