package common

import (
	"flag"
)

func Banner() {
	banner := `
   __    ___  ___   ___    __    _  _ 
  /__\  / __)/ __) / __)  /__\  ( \( )
 /(__)\ \__ \\__ \( (__  /(__)\  )  ( 
(__)(__)(___/(___/ \___)(__)(__)(_)\_)
                     asscan version: ` + version + `
`
	print(banner)
}

func Flag(Info *HostInfo) {
	Banner()
	flag.StringVar(&Info.Host, "h", "", "IP address of the host you want to scan,for example: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12")
	flag.StringVar(&Ports, "p", DefaultPorts, "Select a port,for example: 22 | 1-65535 | 22,80,3306")
	flag.StringVar(&HostFile, "hf", "", "host file, -hf ip.txt")
	flag.Parse()
}
