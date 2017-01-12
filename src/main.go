package main

import (
	"fmt"
	"log"

	"github.com/burntsushi/toml"
	"github.com/justyntemme/MooseChat/src/tor"
)

type tomlConfig struct {
	target     string
	encryptKey string
	decryptKey string
}

func main() {
	var conf tomlConfig

	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}
	onionURL := tor.StartTorService()
	fmt.Println(onionURL)
}
