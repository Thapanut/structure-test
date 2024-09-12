package configs

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type App struct {
	Port     int    `mapstructure:"port"`
	TimeZone string `mapstructure:"time_zone"`
}

type DB struct {
	PGSchema SchemaConfig `mapstructure:"pg_schema"`
}

type SchemaConfig struct {
	User                   string `mapstructure:"user"`
	Password               string `mapstructure:"password"`
	SchemaName             string `mapstructure:"schema_name"`
	Database               string `mapstructure:"database"`
	Port                   int    `mapstructure:"port"`
	Host                   string `mapstructure:"host"`
	SSLMode                string `mapstructure:"ssl_mode"`
	MaxIdleConns           int    `mapstructure:"max_idle_connection"`
	MaxOpenConns           int    `mapstructure:"max_open_connection"`
	MaxLifeTimeConnsMinute int    `mapstructure:"max_life_time_connection_minute"`
}

type AppConfig struct {
	App            App    `mapstructure:"app"`
	DB             DB     `mapstructure:"db"`
	JaegerEndpoint string `mapstructure:"jaeger_endpoint"`
	ServiceName    string `mapstructure:"service_name"`
	FileSize       int    `mapstructure:"file_size"`
}

type PgDatabaseConfig struct {
	User     string
	Password string
	DB       string
	Port     int
	Host     string
	SSLMode  string
	Schema   string
}

var Config = &AppConfig{}

func Setup() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.yaml': %v", err)
	}

	// Unmarshalling the config file into the AppConfig struct
	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	log.Info(Config.FileSize)
}

func GetCommonPgDatabaseConfig() *PgDatabaseConfig {
	return &PgDatabaseConfig{
		User:     Config.DB.PGSchema.User,
		Password: Config.DB.PGSchema.Password,
		DB:       Config.DB.PGSchema.Database,
		Port:     Config.DB.PGSchema.Port,
		Host:     Config.DB.PGSchema.Host,
		SSLMode:  Config.DB.PGSchema.SSLMode,
		Schema:   Config.DB.PGSchema.SchemaName,
	}
}
