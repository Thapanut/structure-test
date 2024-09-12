package checkPermissionService

import (
	"context"

	"github.com/Thapanut/struct-test/constants/e"
	"github.com/Thapanut/struct-test/errs"
	"github.com/Thapanut/struct-test/repositories/accessControl"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type AccessControlService interface {
	CheckPermission(role []string, url string, method string, ctx context.Context) (bool, error)
}

type accessControlService struct {
	accessControlRepo accessControl.AccessControlRepository
}

func NewAccessControlService(accessControlRepo accessControl.AccessControlRepository) AccessControlService {
	return &accessControlService{
		accessControlRepo: accessControlRepo,
	}
}

func (s *accessControlService) CheckPermission(role []string, url string, method string, ctx context.Context) (bool, error) {
	ValidRole, err := s.accessControlRepo.CheckPermission(role, url, method, ctx)
	if err != nil {
		log.Error(err)
		if err == gorm.ErrRecordNotFound {
			return false, errs.AppError{
				Code:    e.DATA_NOT_FOUND,
				BuCode:  "404",
				Message: err.Error(),
			}
		}
		return false, errs.AppError{
			Code:    e.INTERNAL_SERVER_ERROR,
			BuCode:  "500",
			Message: err.Error(),
		}
	}
	return ValidRole, nil
}
