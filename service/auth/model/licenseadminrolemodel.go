package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LicenseAdminRoleModel = (*customLicenseAdminRoleModel)(nil)

type (
	// LicenseAdminRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLicenseAdminRoleModel.
	LicenseAdminRoleModel interface {
		licenseAdminRoleModel
	}

	customLicenseAdminRoleModel struct {
		*defaultLicenseAdminRoleModel
	}
)

// NewLicenseAdminRoleModel returns a model for the database table.
func NewLicenseAdminRoleModel(conn sqlx.SqlConn, c cache.CacheConf) LicenseAdminRoleModel {
	return &customLicenseAdminRoleModel{
		defaultLicenseAdminRoleModel: newLicenseAdminRoleModel(conn, c),
	}
}
