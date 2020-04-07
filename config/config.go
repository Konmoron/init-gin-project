package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// config
type config struct {
	// app name
	AppName string `yaml:"app_name" mapstructure:"app_name"`

	// http addr
	Addr string `yaml:"addr" mapstructure:"addr"`

	// gin model
	// "debug" "release"
	GinModel string `yaml:"gin_model" mapstructure:"gin_model"`

	// log level
	LogLevel string `yaml:"log_level" mapstructure:"log_level"`
}

var (
	cfgFile = "config/config.yaml"
	// Cfg : global config
	Cfg = &config{}
)

// InitConf : init config
func InitConf() {
	fmt.Println("Read CfgFile: ", cfgFile)
	viper.SetConfigFile(cfgFile)

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Read Config ERROR", err)
		return
	}

	// Unmarshal to rc
	err = viper.Unmarshal(Cfg)
	if err != nil {
		fmt.Println("Unmarshal Config ERROR", err)
		return
	}
}
