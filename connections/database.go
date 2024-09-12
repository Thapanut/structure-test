package connections

import (
	"fmt"
	"time"

	"github.com/Thapanut/struct-test/configs"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type SqlLogger struct {
	logger.Interface
}

func SetupDB(schemaConfig configs.SchemaConfig) (*gorm.DB, error) {
	var err error

	host := schemaConfig.Host
	user := schemaConfig.User
	pass := schemaConfig.Password
	db := schemaConfig.Database
	port := fmt.Sprint(schemaConfig.Port)
	sslMode := schemaConfig.SSLMode

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, pass, db, port, sslMode,
	)
	log.Info(dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
	})
	if err != nil {
		log.Error("models.Setup err: ", err)
		return nil, err
	}
	sqlDB, err := DB.DB()
	if err != nil {
		log.Error("models.Setup err: ", err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(schemaConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(schemaConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(schemaConfig.MaxLifeTimeConnsMinute))
	return DB, nil
}
