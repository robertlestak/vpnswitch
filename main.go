package main

import (
	"fmt"
	"github.com/adal-io/vpnswitch/vpnswitch"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		vpnswitch.Start()
		return
	}
	var e error
	switch os.Args[1] {
	case "start":
		e = vpnswitch.Start()
	case "switch":
		e = vpnswitch.Switch()
	case "stop":
		e = vpnswitch.Stop()
	}
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
}
