package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LicenseUserModel = (*customLicenseUserModel)(nil)

type (
	// LicenseUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLicenseUserModel.
	LicenseUserModel interface {
		licenseUserModel
	}

	customLicenseUserModel struct {
		*defaultLicenseUserModel
	}
)

// NewLicenseUserModel returns a model for the database table.
func NewLicenseUserModel(conn sqlx.SqlConn, c cache.CacheConf) LicenseUserModel {
	return &customLicenseUserModel{
		defaultLicenseUserModel: newLicenseUserModel(conn, c),
	}
}
