package vpnswitch

import (
	"bytes"
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

func Connect(l string) {
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
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if strings.Contains(output.String(), "AUTH_FAILED") {
		fmt.Println("Authorization failed")
		os.Exit(1)
	}
	fmt.Println("Successfully connected")
}

func getLocation() (string, error) {
	var l string
	var e error
	ls, err := ioutil.ReadDir(GetDataPath())
	if err != nil {
		fmt.Println(err)
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

func Stop() {
	params := []string{"killall", "-9", "openvpn"}
	cmd := exec.Command("sudo", params...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection closed")
}

func Start() {
	l, err := getLocation()
	if err != nil {
		fmt.Println(err)
		return
	}
	Connect(l)
}

func Switch() {
	Stop()
	time.Sleep(time.Second * 5)
	Start()
}
