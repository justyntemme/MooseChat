package tor

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

//StartTorService starts tor service and returns url
func StartTorService() string {
	fmt.Printf("Starting Tor Service .....\n")
	cmd := exec.Command("tor", "-f", "/home/user/go/src/github.com/justyntemme/MooseChat/torrc")
	err := cmd.Start()
	out, _ := exec.Command("sleep", "1").Output()
	fmt.Println(out)
	if err != nil {
		log.Fatal(err)
	}

	return getOnionURL()
}

func starttor() {

}

func getOnionURL() string {
	cmd := exec.Command("cat", ".moosechat/hostname")
	stdout, errout := cmd.StdoutPipe()
	if errout != nil {
		log.Fatal(errout)
	}
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(stdout)
	newStr := buf.String()
	return newStr
}
