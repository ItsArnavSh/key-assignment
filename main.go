package main

import (
	"stack/src/core/codegen"
	"stack/src/server"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	if viper.GetBool("generate") {
		codegen.GenerateLargeCodebase()
	}
	server := server.NewServer()
	server.ScanCodebase(viper.GetString("target_loc"), viper.GetString("target_url"))
}
