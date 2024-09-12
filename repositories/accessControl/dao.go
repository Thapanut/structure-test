package accessControl

import (
	"context"

	"gorm.io/gorm"
)

type AccessControlRepository interface {
	CheckPermission(role []string, url string, method string, ctx context.Context) (bool, error)
}

type accessControlRepository struct {
	db *gorm.DB
}

func NewAccessControlRepository(db *gorm.DB) accessControlRepository {
	return accessControlRepository{db: db}
}

func (r *accessControlRepository) CheckPermission(roles []string, url string, method string, ctx context.Context) (bool, error) {
	var accessControl AccessControl
	for _, role := range roles {
		err := r.db.Where("role_code = ? and url  = ? and  http_method = ? ", role, url, method).Find(&accessControl).Error
		if err != nil {
			return false, err
		}
		if role == accessControl.RoleCode && url == accessControl.URL && method == accessControl.HTTPMethod {
			return true, nil
		}
	}
	return false, nil
}
