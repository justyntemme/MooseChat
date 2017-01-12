package main

import (
	"fmt"

	"github.com/burntsushi/toml"
	"github.com/justyntemme/MooseChat/src/tor"
)

func main() {
	var conf Config

	onionURL := tor.StartTorService()
	fmt.Println(onionURL)
}

func readConfig() Config{

}