package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Environment string
	Postgre     PostgreConfiguration
}

type PostgreConfiguration struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

func GetConfig() Configuration {
	conf := Configuration{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return conf
}

func GetURL() string {
	conf := GetConfig()
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		conf.Postgre.User,
		conf.Postgre.Password,
		conf.Postgre.Host,
		conf.Postgre.Port,
		conf.Postgre.DBname,
	)
	return url
}
