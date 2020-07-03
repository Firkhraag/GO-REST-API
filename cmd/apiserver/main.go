package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/firkhraag/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		color.Set(color.FgRed)
		log.Println(err.Error())
		color.Unset()
		os.Exit(1)
	}

	if err := apiserver.StartServer(config); err != nil {
		color.Set(color.FgRed)
		log.Println(err.Error())
		color.Unset()
		os.Exit(1)
	}
}
