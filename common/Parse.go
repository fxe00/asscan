package common

import (
	"flag"
	"fmt"
	"os"
)

func Parse(Info *HostInfo) {

}

func ParseInput(Info *HostInfo) {
	if Info.Host == "" && HostFile == "" {
		fmt.Println("Host is none")
		flag.Usage()
		os.Exit(0)
	}

	if Ports == DefaultPorts {
		Ports += "," + Webport
	}
}
