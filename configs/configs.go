package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// TODO: fetch configs and return structs of instances.

type Config struct {
	Postgres
	Redis
}

type Postgres struct {
	Host     string
	Port     string
	Username string
	DBname   string
	Password string
	SSLMode  string
}

type Redis struct {
	Addr     string
	Password string
}

func Get() (Config, error) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	var cfg Config

	if err := newConfig(); err != nil {
		return cfg, fmt.Errorf("error initializing configs: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		return cfg, fmt.Errorf("error loading env variables: %s", err)
	}

	cfg = Config{
		Postgres{
			Host:     viper.GetString("psql.host"),
			Port:     viper.GetString("psql.port"),
			Username: viper.GetString("psql.username"),
			DBname:   viper.GetString("psql.dbname"),
			Password: os.Getenv("PSQL_PASSWORD"),
			SSLMode:  viper.GetString("psql.sslmode"),
		},
		Redis{
			Addr:     viper.GetString("redis.addr"),
			Password: "",
		},
	}
	return cfg, nil
}

func newConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
