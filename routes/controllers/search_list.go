package controllers

import (
	"fmt"

	"github.com/Thapanut/struct-test/constants/e"
	"github.com/Thapanut/struct-test/errs"
	searchList "github.com/Thapanut/struct-test/services/SearchService"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (server *Server) SearchListHandle(c *fiber.Ctx) error {

	req := searchList.SearchListReq{}
	response := searchList.SearchSanctionListResponseStandard{}
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Error parse request body: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		log.Error(err.Error())
		return errs.ReturnError(c, e.BAD_REQUEST, fmt.Sprint(e.BAD_REQUEST), err.Error())
	}
	resp, err := server.api.SearchSvConf.SearchList(req)
	if err != nil {
		return errs.ReturnAppError(c, err)
	}
	if resp != nil {
		return errs.ReturnError(c, e.SUCCESS, fmt.Sprint(e.DATA_NOT_FOUND), "data not found")
	}
	response.Code = fmt.Sprint(e.SUCCESS)
	response.Message = "search list success"
	response.Status = e.GetMsg(e.SUCCESS)
	response.Data = resp
	return c.Status(e.SUCCESS).JSON(response)
}
