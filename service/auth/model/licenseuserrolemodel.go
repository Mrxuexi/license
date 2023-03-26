package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LicenseUserRoleModel = (*customLicenseUserRoleModel)(nil)

type (
	// LicenseUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLicenseUserRoleModel.
	LicenseUserRoleModel interface {
		licenseUserRoleModel
	}

	customLicenseUserRoleModel struct {
		*defaultLicenseUserRoleModel
	}
)

// NewLicenseUserRoleModel returns a model for the database table.
func NewLicenseUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf) LicenseUserRoleModel {
	return &customLicenseUserRoleModel{
		defaultLicenseUserRoleModel: newLicenseUserRoleModel(conn, c),
	}
}
