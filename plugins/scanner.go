package plugins

import (
	"fmt"
	"sync"

	"github.com/Fxe-h/asscan/common"
	"github.com/Fxe-h/asscan/lib"
)

func Scan(info common.HostInfo) {
	fmt.Println("start infoscan")
	Hosts, err := common.ParseIP(info.Host, common.HostFile, common.NoHosts)
	if err != nil {
		fmt.Println("len(hosts)==0", err)
		return
	}
	lib.Inithttp()
	var ch = make(chan struct{}, common.Threads)
	var wg = sync.WaitGroup{}
	if len(Hosts) > 0 || len(common.HostPort) > 0 {
		AlivePorts = PortScan(Hosts, common.Ports, common.Timeout)
	}
}
