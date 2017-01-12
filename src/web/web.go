package web

import "golang.org/x/net/proxy"

func CreateDialer() Dialer {
dialer, err :=proxy.SOCKS5("tcp", "localhost", nil, proxy.Direct)
if err != nil{
  log.Fatal(err)
return dialer
}

func SendMSG(string s,Dialer dialer, string url) {
  conn, err :=dialer.Dial("tcp", url)
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()
}
