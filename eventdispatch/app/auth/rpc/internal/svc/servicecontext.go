package svc

import (
	"eventdispatch/app/auth/model"
	"eventdispatch/app/auth/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
