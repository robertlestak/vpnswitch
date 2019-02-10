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

func run() {
	for {
		c := make(chan string, 1)
		go checkIP(c)
		select {
		case ip := <-c:
			fmt.Println("IP:", ip)
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout Getting IP")
		}
		time.Sleep(time.Second * 10)
	}
}

func main() {
	var e error
	switch os.Args[1] {
	case "start":
		e = vpnswitch.Start()
	case "switch":
		e = vpnswitch.Switch()
	case "stop":
		e = vpnswitch.Stop()
	case "run":
		e = vpnswitch.Start()
		run()
	}
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
}
