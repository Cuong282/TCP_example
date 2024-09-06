package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("../../.././config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	fmt.Println("viper:", viper)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failded to read configuration %v \n", err))

	}

	fmt.Println("server port::", viper.GetInt("server.port"))
	fmt.Println("server port::", viper.GetString("security.jwt.key"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}

	fmt.Println("config port:", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Println("database User: %s, password: %s,host: %s\n", db.User, db.Password, db.Host)
	}
}
