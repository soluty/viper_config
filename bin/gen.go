// generated by viper-config, do not change.

package config

import (
	"flag"
	"fmt"

	"time"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var C = Config{}
var watchFunc func()

func Init(watch func(), filePath ...string) {
	watchFunc = watch
	setFlags()
	setEnv()
	readConfigFile(filePath)
	setDefaults()
	err := viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}
}

func setFlags() {

	pflag.Int("a", 2, "一个整数a")
	pflag.Duration("nest.b", 0 * time.Nanosecond, "")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

func setEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("app")

	_ = viper.BindEnv("a")
	_ = viper.BindEnv("nest.b")
}

func setDefaults() {

	viper.SetDefault("a", 2)
	viper.SetDefault("nest.b", 0 * time.Nanosecond)
}

func readConfigFile(path []string) {
	viper.SetConfigName("config")
	if len(path) == 0 {
		viper.AddConfigPath(".")
	}else {
		for _, value := range path {
			viper.AddConfigPath(value)
		}
	}
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	if watchFunc != nil {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			err := viper.Unmarshal(&C)
			if err != nil {
				fmt.Println(err)
				return
			}
			watchFunc()
		})
	}
}

