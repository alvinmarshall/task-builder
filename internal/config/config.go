package config

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	Application struct {
		Name string `default:"my-application"`
	}
	Server struct {
		Port string `env:"PORT"`
	}
	DataSource struct {
		Use      string `default:"postgres"`
		Postgres struct {
			Enabled  bool   `default:"true"`
			Username string `env:"DB_USER"`
			Password string `env:"DB_PASSWORD"`
			Database string `env:"DB_NAME"`
			Port     string `env:"DB_PORT"`
			Host     string `env:"DB_HOST"`
			Dialect  string `default:"postgres"`
		}
	}
	Logger struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"dev"`
		LogLevel    string `default:"debug"`
		FileName    string `default:"app.log"`
	}
	Jwt struct {
		Secret  string `default:"secrete" env:"JWT_SECRET"`
		Expires int64  `env:"JWT_EXPIRES"`
		Issuer  string `default:"my-application" env:"APPLICATION_NAME"`
	}
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := configor.Load(config, "./config.yml")
	if err != nil {
		return nil, err
	}
	return config, nil
}
