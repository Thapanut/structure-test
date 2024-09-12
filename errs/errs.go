package errs

import (
	"fmt"

	"github.com/Thapanut/struct-test/models"

	"github.com/Thapanut/struct-test/constants/e"
	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code    int
	BuCode  string
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

// for service layer
func ReturnAppError(c *fiber.Ctx, err error) error {
	resp := models.ResponseStandard{}
	appErr, ok := err.(AppError)
	if ok {
		resp = models.ResponseStandard{
			Code:    appErr.BuCode,
			Status:  e.GetMsg(e.ERROR),
			Message: fmt.Sprintf("%s: %s", e.GetMsg(appErr.Code), appErr.Message),
		}
		return c.Status(appErr.Code).JSON(resp)
	}
	resp = models.ResponseStandard{
		Code:    "500",
		Status:  e.GetMsg(e.ERROR),
		Message: fmt.Sprintf("%s: %s", e.GetMsg(e.INTERNAL_SERVER_ERROR), err.Error()),
	}
	return c.Status(500).JSON(resp)
}

// for controller layer
func ReturnError(c *fiber.Ctx, code int, buCode string, msg string) error {
	resp := models.ResponseStandard{
		Code:    buCode,
		Status:  e.GetMsg(e.ERROR),
		Message: fmt.Sprintf("%s: %s", e.GetMsg(code), msg),
	}
	return c.Status(code).JSON(resp)
}

func ReturnErrorWithData(c *fiber.Ctx, code int, buCode string, msg string, data interface{}) error {
	resp := models.ResponseStandard{
		Code:    buCode,
		Status:  e.GetMsg(e.ERROR),
		Message: fmt.Sprintf("%s: %s", e.GetMsg(code), msg),
		Data:    data,
	}
	return c.Status(code).JSON(resp)
}
