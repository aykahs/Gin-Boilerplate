package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql  Mysql
	Server Server
	Jwt    Jwt
}

var EnvConfig *Config

func LoadConfig(env string) *Config {
	println(env)
	path, err := os.Getwd() // get curent path
	if err != nil {
		panic(err)
	}

	viper.SetConfigName(env + ".config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/configs") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
	// switch gin.Mode() {
	// case gin.DebugMode:
	// 	config = config
	// case gin.ReleaseMode:
	// 	config = config
	// default:
	// 	fmt.Printf(fmt.Sprintf("Unknown gin mode %s", gin.Mode()))
	// }
	return config
}

func Init(env string) {
	EnvConfig = LoadConfig(env)
}
