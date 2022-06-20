package main

import (
	"errors"
	"fmt"
	"github.com/FuradWho/Cherry/internal/apiserver"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func initConfig(cfgFile string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {

		/*
			viper.SetDefault("ContentDir", "content")
			viper.SetDefault("LayoutDir", "layouts")
			viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
		*/

		return errors.New("no config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()           // 读取匹配的环境变量
	viper.SetEnvPrefix("GOSERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	return nil

}

func main() {
	initConfig("D:\\projects\\Cherry\\configs\\goserver.yaml")
	apiserver.NewGoServer()
}
