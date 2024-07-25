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
		ports := ""
		for i := 1; i <= 65535; i++ {
			if i > 1 {
				ports += ","
			}
			ports += fmt.Sprintf("%d", i)
		}
		Ports = ports
	}
}
