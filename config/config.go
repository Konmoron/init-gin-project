package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/iancoleman/strcase"
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

	// LimitRedisHost
	LimitRedisHost string `yaml:"limit_redis_host" mapstructure:"limit_redis_host"`
}

var (
	cfgFile = "config/config.yaml"
	// Cfg : global config
	Cfg = &config{}
)

// read config.yaml
func (c *config) readFromFile() {
	fmt.Println("Read config from ", cfgFile)
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

// read from OS.ENV
// References:
// 	1. https://www.jianshu.com/p/40cb706853cd
// 	2. https://stackoverflow.com/questions/6395076/using-reflect-how-do-you-set-the-value-of-a-struct-field
func (c *config) readFromOSEnv() {
	fmt.Println("read config from OS.ENV")
	// TypeOf 的参数必须为 值类型，不能是指针，否则
	// 在执行 NumField() 的时候，会报错：
	//	panic: reflect: NumField of non-struct type *main.config
	fields := reflect.TypeOf(*c)

	// ValueOf 的参数必须为 指针类型 或者 Interface，
	// 不能为 值类型，否则，在执行 Elem() 会报错：
	//	panic: reflect: call of reflect.Value.Elem on struct Value
	values := reflect.ValueOf(c)
	// fmt.Println("value Kind after reflect.ValueOf(f) is ", values.Kind())

	if values.Kind() != reflect.Ptr || !values.Elem().CanSet() {
		fmt.Println("输入的对象是不可修改的")
		return
	} else {
		values = values.Elem()
		// fmt.Println("value Kind after values.Elem() is ", values.Kind())
	}

	for i := 0; i < values.NumField(); i++ {
		// get field name
		fieldName := fields.Field(i).Name

		// get Env name in os
		osEnvName := strcase.ToScreamingSnake(fieldName)
		// get env
		osEnvValue, exist := os.LookupEnv(osEnvName)

		if exist {
			// FieldByIndex() 的参数是一个数组
			v := values.FieldByIndex([]int{i})

			if !v.IsValid() {
				fmt.Println("没有找到", fieldName, "的字段")
				continue
			}

			// 获取 value 的类型
			fieldType := v.Type().Kind()

			switch fieldType {
			case reflect.String:
				v.SetString(osEnvValue)
				fmt.Println("type of", fieldName, "is", fieldType, ", env name in os is", osEnvName, ", env value in os is", osEnvValue)
			case reflect.Int:
				if i, err := strconv.ParseInt(osEnvValue, 10, 64); err != nil {
					v.SetInt(i)
					fmt.Println("type of", fieldName, "is", fieldType, ", env name in os is", osEnvName, ", env value in os is", osEnvValue)
				} else {
					fmt.Println("strconv.ParseInt(osEnvValue, 10, 64) ERR:", err)
				}
			default:
				fmt.Println("not match")
			}
		}
	}
}

// InitConf : init config
// first read config from config.yaml,
// then read config from env in os
func InitConf() {
	Cfg.readFromFile()
	Cfg.readFromOSEnv()
}
