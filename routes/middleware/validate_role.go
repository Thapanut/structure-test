package middleware

import (
	"github.com/Thapanut/struct-test/services/checkPermissionService"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	accessControlService checkPermissionService.AccessControlService
	// db                   *gorm.DB
}

func NewAuthMiddleware(accessControlService checkPermissionService.AccessControlService) AuthMiddleware {
	return AuthMiddleware{accessControlService: accessControlService}
}

func (m *AuthMiddleware) ValidatePermission(c *fiber.Ctx) error {

	return c.Next()
}
