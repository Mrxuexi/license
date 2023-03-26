// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	licenseAdminRoleFieldNames          = builder.RawFieldNames(&LicenseAdminRole{})
	licenseAdminRoleRows                = strings.Join(licenseAdminRoleFieldNames, ",")
	licenseAdminRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(licenseAdminRoleFieldNames, "`auto_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	licenseAdminRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(licenseAdminRoleFieldNames, "`auto_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLicenseAdminRoleAutoIdPrefix = "cache:licenseAdminRole:autoId:"
)

type (
	licenseAdminRoleModel interface {
		Insert(ctx context.Context, data *LicenseAdminRole) (sql.Result, error)
		FindOne(ctx context.Context, autoId int64) (*LicenseAdminRole, error)
		Update(ctx context.Context, data *LicenseAdminRole) error
		Delete(ctx context.Context, autoId int64) error
	}

	defaultLicenseAdminRoleModel struct {
		sqlc.CachedConn
		table string
	}

	LicenseAdminRole struct {
		AutoId         int64        `db:"auto_id"`          // 用户角色ID
		LicenseAdminId int64        `db:"license_admin_id"` // 软件服务提供商ID
		RoleId         int64        `db:"role_id"`          // 角色ID
		IsDeleted      int64        `db:"is_deleted"`       // 软删除
		CreatedAt      time.Time    `db:"created_at"`       // 创建时间
		UpdatedAt      time.Time    `db:"updated_at"`       // 更新时间
		DeletedAt      sql.NullTime `db:"deleted_at"`       // 删除时间
	}
)

func newLicenseAdminRoleModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultLicenseAdminRoleModel {
	return &defaultLicenseAdminRoleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`license_admin_role`",
	}
}

func (m *defaultLicenseAdminRoleModel) Delete(ctx context.Context, autoId int64) error {
	licenseAdminRoleAutoIdKey := fmt.Sprintf("%s%v", cacheLicenseAdminRoleAutoIdPrefix, autoId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `auto_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, autoId)
	}, licenseAdminRoleAutoIdKey)
	return err
}

func (m *defaultLicenseAdminRoleModel) FindOne(ctx context.Context, autoId int64) (*LicenseAdminRole, error) {
	licenseAdminRoleAutoIdKey := fmt.Sprintf("%s%v", cacheLicenseAdminRoleAutoIdPrefix, autoId)
	var resp LicenseAdminRole
	err := m.QueryRowCtx(ctx, &resp, licenseAdminRoleAutoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `auto_id` = ? limit 1", licenseAdminRoleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, autoId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLicenseAdminRoleModel) Insert(ctx context.Context, data *LicenseAdminRole) (sql.Result, error) {
	licenseAdminRoleAutoIdKey := fmt.Sprintf("%s%v", cacheLicenseAdminRoleAutoIdPrefix, data.AutoId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, licenseAdminRoleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.LicenseAdminId, data.RoleId, data.IsDeleted, data.DeletedAt)
	}, licenseAdminRoleAutoIdKey)
	return ret, err
}

func (m *defaultLicenseAdminRoleModel) Update(ctx context.Context, data *LicenseAdminRole) error {
	licenseAdminRoleAutoIdKey := fmt.Sprintf("%s%v", cacheLicenseAdminRoleAutoIdPrefix, data.AutoId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `auto_id` = ?", m.table, licenseAdminRoleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.LicenseAdminId, data.RoleId, data.IsDeleted, data.DeletedAt, data.AutoId)
	}, licenseAdminRoleAutoIdKey)
	return err
}

func (m *defaultLicenseAdminRoleModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLicenseAdminRoleAutoIdPrefix, primary)
}

func (m *defaultLicenseAdminRoleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `auto_id` = ? limit 1", licenseAdminRoleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLicenseAdminRoleModel) tableName() string {
	return m.table
}
