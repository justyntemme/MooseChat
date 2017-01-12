package crypto

import "os/exec"

func decrypt(string s) string{
	cmd := exec.Command("gpg", "-d", s)
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
