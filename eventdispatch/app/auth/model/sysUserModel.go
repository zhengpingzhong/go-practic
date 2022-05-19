package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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
	const query = `select user_id, username, password, salt, email, mobile, expire_date, status, create_user_id, create_time, staff_id from sys_user where username = ?`
	err := m.conn.QueryRow(&resp, query, userName)
	if err != nil {
		logx.Errorf("SysUserModel.FindByUserName error, userName=%v, err=%s", userName, err.Error())
		if err == sqlx.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &resp, nil
}

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn),
	}
}
