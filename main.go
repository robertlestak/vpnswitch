package main

import (
	"fmt"
	"os"

	"github.com/robertlestak/vpnswitch/vpnswitch"
)

func main() {
	if len(os.Args) < 2 {
		vpnswitch.Switch()
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
