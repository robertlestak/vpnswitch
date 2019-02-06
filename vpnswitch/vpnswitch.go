package vpnswitch

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func dataPath() string {
	if os.Getenv("VPN_DATA_DIR") != "" {
		return os.Getenv("VPN_DATA_DIR")
	}
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	p := path.Join(cwd, "data", "openvpn")
	return p
}

// Connect connects to VPN server with configuration file l
func Connect(l string) error {
	var e error
	fmt.Printf("Connecting to location %s\n", strings.Replace(l, ".ovpn", "", -1))
	params := []string{"openvpn",
		"--config",
		path.Join(dataPath(), l),
		"--ca",
		path.Join(dataPath(), "ca.rsa.2048.crt"),
		"--crl-verify",
		path.Join(dataPath(), "crl.rsa.2048.pem"),
		"--auth-user-pass",
		path.Join(dataPath(), "auth.txt"),
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
	ls, err := ioutil.ReadDir(dataPath())
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

// Stop stops all openvpn processes
func Stop() error {
	var e error
	params := []string{"pkill", "openvpn"}
	cmd := exec.Command("sudo", params...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return e
}

// CreateAuth creates openvpn authorization file
func CreateAuth() error {
	var e error
	d := []byte(os.Getenv("OPENVPN_USERNAME") + "\n" + os.Getenv("OPENVPN_PASSWORD"))
	p := path.Join(dataPath(), "auth.txt")
	err := ioutil.WriteFile(p, d, 0644)
	if err != nil {
		return err
	}
	return e
}

// Start starts openvpn connection
func Start() error {
	var e error
	if aerr := CreateAuth(); aerr != nil {
		return aerr
	}
	l, err := getLocation()
	if err != nil {
		return err
	}
	if l == "" {
		return errors.New("Configuration file not found")
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
