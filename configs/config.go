package configs

import (
	"sync"

	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port     string `yaml:"port"`
	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var testConfig AppConfig
	testConfig.Port = "1324"
	testConfig.Database.Driver = "mysql"
	testConfig.Database.Name = "todo_test_db"
	testConfig.Database.Host = "localhost"
	testConfig.Database.Port = "3306"
	testConfig.Database.Username = "root"
	testConfig.Database.Password = "root"

	viper.SetConfigFile("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("Please make a config file!")
		return &testConfig
	}

	var finalConfig AppConfig

	err := viper.Unmarshal(&finalConfig)

	if err != nil {
		panic("Failed extract external config!")
	}

	return &finalConfig
}
