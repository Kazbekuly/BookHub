package config

import (
	"BookHub/internal/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIp string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Db struct {
		Username string `yaml:"username" env-default:"username"`
		Password string `yaml:"password" env-default:"password"`
		Host     string `yaml:"host" env-default:"host"`
		Port     string `yaml:"port" env-default:"port"`
		Dbname   string `yaml:"dbname" env-default:"dbname"`
		SslMode  string `yaml:"sslmode" env-default:"sslmode"`
	} `yaml:"db"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
