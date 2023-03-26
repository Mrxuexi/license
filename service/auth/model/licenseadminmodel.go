package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LicenseAdminModel = (*customLicenseAdminModel)(nil)

type (
	// LicenseAdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLicenseAdminModel.
	LicenseAdminModel interface {
		licenseAdminModel
	}

	customLicenseAdminModel struct {
		*defaultLicenseAdminModel
	}
)

// NewLicenseAdminModel returns a model for the database table.
func NewLicenseAdminModel(conn sqlx.SqlConn, c cache.CacheConf) LicenseAdminModel {
	return &customLicenseAdminModel{
		defaultLicenseAdminModel: newLicenseAdminModel(conn, c),
	}
}
