package config

import "github.com/jinzhu/configor"

type Config struct {
	Application struct {
		Name string `default:"my-application"`
	}
	Server struct {
		Port string `default:"8080"`
	}
	DataSource struct {
		Use      string `default:"postgres"`
		Postgres struct {
			Enabled  bool   `default:"true"`
			Username string `default:"postgres"`
			Password string `default:"postgres"`
			Database string
			Port     string `default:"5432"`
			Host     string `default:"localhost"`
			Dialect  string `default:"postgres"`
		}
	}
	Logger struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"dev"`
		LogLevel    string `default:"debug"`
		FileName    string `default:"app.log"`
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
