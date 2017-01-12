package main

import (
	"fmt"

	"github.com/burntsushi/toml"
	"github.com/justyntemme/MooseChat/src/tor"
)

func main() {
	onionURL := tor.StartTorService()
	fmt.Println(onionURL)
}

func readConfig() {

}