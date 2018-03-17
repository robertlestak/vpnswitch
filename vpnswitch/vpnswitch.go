package vpnswitch

import (
	"bytes"
	"errors"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

func GetDataPath() string {
	if os.Getenv("VPN_DATA_DIR") != "" {
		return os.Getenv("VPN_DATA_DIR")
	}
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	p := path.Join(cwd, "data", "openvpn")
	return p
}

func Connect(l string) error {
	var e error
	fmt.Printf("Connecting to location %s\n", strings.Replace(l, ".ovpn", "", -1))
	params := []string{"openvpn",
		"--config",
		path.Join(GetDataPath(), l),
		"--ca",
		path.Join(GetDataPath(), "ca.rsa.2048.crt"),
		"--crl-verify",
		path.Join(GetDataPath(), "crl.rsa.2048.pem"),
		"--auth-user-pass",
		path.Join(GetDataPath(), "auth.txt"),
	}
	cmd := exec.Command("sudo", params...)
	var output bytes.Buffer
	var outputErr bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &outputErr
	err := cmd.Start()
	if strings.Contains(output.String(), "AUTH_FAILED") {
		e = errors.New("AUTH_FAILED")
	}
	if outputErr.String() != "" {
		e = errors.New(outputErr.String())
	}
	if err != nil {
		e = err
	}
	return e
}

func getLocation() (string, error) {
	var l string
	var e error
	ls, err := ioutil.ReadDir(GetDataPath())
	if err != nil {
		return l, err
	}
	var fs []string
	for _, f := range ls {
		if !strings.Contains(f.Name(), ".ovpn") {
			continue
		}
		fs = append(fs, f.Name())
	}
	rand.Seed(time.Now().Unix())
	l = fs[rand.Intn(len(fs))]
	return l, e
}

func Stop() error {
	var e error
	params := []string{"openvpn"}
	cmd := exec.Command("pkill", params...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return e
}

func Start() error {
	var e error
	l, err := getLocation()
	if err != nil {
		return err
	}
	cerr := Connect(l)
	if cerr != nil {
		return cerr
	}
	return e
}

func Switch() error {
	var e error
	err := Stop()
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 5)
	serr := Start()
	if serr != nil {
		return serr
	}
	return e
}
