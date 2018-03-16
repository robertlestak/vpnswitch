package main

import (
	"github.com/adal-io/vpnswitch/vpnswitch"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		vpnswitch.Start()
		return
	}
	switch os.Args[1] {
	case "start":
		vpnswitch.Start()
	case "switch":
		vpnswitch.Switch()
	case "stop":
		vpnswitch.Stop()
	}
}
