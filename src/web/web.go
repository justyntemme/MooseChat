package web

import (
	"log"

	"golang.org/x/net/proxy"
)

//CreateDialer creates a proxy to the local tor connection returning the dialer to send msg
func CreateDialer() proxy.Dialer {
	dialer, err := proxy.SOCKS5("tcp", "localhost:9050", nil, proxy.Direct)
	if err != nil {
		log.Fatal(err)
	}
	return dialer
}

//SendMsg sends encrypted texts to hidden onion address
func SendMsg(dialer proxy.Dialer, target string, message string) {
	conn, err := dialer.Dial("tcp", target)
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte(message))
	defer conn.Close()
}
