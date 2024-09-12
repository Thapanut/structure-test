package controllers

import (
	"github.com/Thapanut/struct-test/configs"
	"github.com/Thapanut/struct-test/connections"
	"github.com/Thapanut/struct-test/routes/middleware"

	repositories "github.com/Thapanut/struct-test/repositories/searchList"
	SearchSv "github.com/Thapanut/struct-test/services/SearchService"
	"github.com/gofiber/fiber/v2/log"
)

type API struct {
	middleware   middleware.AuthMiddleware
	SearchSvConf SearchSv.IServiceSearch
}
type Server struct {
	api    API
	Config configs.AppConfig
}

func NewServer() *Server {
	conf := *configs.Config
	commonPgConfig := configs.GetCommonPgDatabaseConfig()
	commonSchema, err := connections.ConnectPG(commonPgConfig)
	if err != nil {
		log.Errorf("Failed to connect  common to PG. %v\n", err)
		panic(err)
	}
	searchListRepo := repositories.NewSearchListDaoServices(commonSchema)
	searchListSer := SearchSv.NewSearchService(searchListRepo)
	return &Server{
		api: API{
			SearchSvConf: searchListSer,
		},
		Config: conf,
	}
}
