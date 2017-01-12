package main

import (
	"fmt"

	"github.com/justyntemme/MooseChat/src/tor"
)

func main() {
	onionURL := tor.StartTorService()
	fmt.Println(onionURL)
}
