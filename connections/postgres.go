package connections

import (
	"fmt"

	"github.com/Thapanut/struct-test/configs"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func ConnectPG(pgConfig *configs.PgDatabaseConfig) (*gorm.DB, error) {
	host := pgConfig.Host
	user := pgConfig.User
	pwd := pgConfig.Password
	db := pgConfig.DB
	port := pgConfig.Port
	sslMode := pgConfig.SSLMode
	schema := pgConfig.Schema
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
		host, user, pwd, db, port, sslMode, schema,
	)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
	})
	if err != nil {
		log.Errorf("Error Connect PG database: %s\n", err.Error())
		return nil, err
	}
	log.Infof("PG DSN: %v\n", dsn)

	// Setup Tracing for DB
	if err := dbConn.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}

	return dbConn, nil
}
