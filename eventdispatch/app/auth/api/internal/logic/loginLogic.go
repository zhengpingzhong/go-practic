package logic

import (
	"context"
	"encoding/base64"
	"eventdispatch/common/tool"
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
	var userName = req.Username
	decodeString, err := base64.StdEncoding.DecodeString(userName)
	if err != nil {
		return nil, err
	}
	decrypt, err := tool.RsaDecrypt(decodeString)
	userName = string(decrypt)

	var password = req.Password
	passwordDecodeString, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return nil, err
	}
	passwordDecrypt, err := tool.RsaDecrypt(passwordDecodeString)
	password = string(passwordDecrypt)
	//TODO Sha256Hash(password,salt) 处理密码后比对
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
