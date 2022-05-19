package svc

import (
	"eventdispatch/app/auth/api/internal/config"
	"eventdispatch/app/auth/model"
	"eventdispatch/app/auth/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserRpc   user.User
	UserModel model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewSysUserModel(sqlx.NewMysql(c.DB.DataSource), nil),
	}
}
