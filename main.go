package main

import (
	"fmt"
	"os"
	"time"

	"github.com/robertlestak/vpnswitch/vpnswitch"
)

func checkIP(c chan string) {
	fmt.Println("Getting IP...")
	ip, err := vpnswitch.CheckIP()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	c <- ip
}

func checkIPWithTimeout() string {
	c := make(chan string, 1)
	go checkIP(c)
	select {
	case ip := <-c:
		return "IP: " + ip
	case <-time.After(time.Second * 5):
		return "Timeout Getting IP"
	}
}

func run(rc chan error) {
	i := 0
	for {
		nip := checkIPWithTimeout()
		fmt.Println(nip)
		if i >= 3 && origIP == nip {
			i = 0
			fmt.Println("Reconnecting...")
			e := vpnswitch.Switch()
			if e != nil {
				rc <- e
			}
		}
		time.Sleep(time.Second * 30)
		i++
	}
}

var origIP string

func main() {
	var e error
	if len(os.Args) <= 1 {
		os.Args = append(os.Args, "run")
	}
	switch os.Args[1] {
	case "start":
		e = vpnswitch.Start()
	case "switch":
		e = vpnswitch.Switch()
	case "stop":
		e = vpnswitch.Stop()
	case "run":
		origIP = checkIPWithTimeout()
		e = vpnswitch.Start()
		rc := make(chan error, 1)
		go run(rc)
		fmt.Println(<-rc)
	}
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
}
