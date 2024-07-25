package plugins

import (
	"fmt"
	"sync"

	"github.com/Fxe-h/asscan/common"
)

func Scan(info common.HostInfo) {
	fmt.Println("start infoscan")
	Hosts, err := common.ParseIP(info.Host, common.HostFile)
	if err != nil {
		fmt.Println("len(hosts)==0", err)
		return
	}
	var wg = sync.WaitGroup{}
	if len(Hosts) > 0 || len(common.HostPort) > 0 {
		var AlivePorts []string = []string{}
		// fmt.Printf("common.Ports: %v\n", common.Ports)
		AlivePorts = PortScan(Hosts, common.Ports, common.Timeout)
		fmt.Println("[*] alive ports len is:", len(AlivePorts))
		wg.Wait()
	}
}
