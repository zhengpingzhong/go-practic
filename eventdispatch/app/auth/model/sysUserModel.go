package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		FindByUserName(ctx context.Context, userName string) (*SysUser, error)
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
)

func (m *defaultSysUserModel) FindByUserName(ctx context.Context, userName string) (*SysUser, error) {
	var resp SysUser
	var query string = "select username,password from sys_user where username = ?"
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userName)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn, c),
	}
}
