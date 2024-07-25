package main

import (
	"fmt"
	"time"

	"github.com/Fxe-h/asscan/common"
	"github.com/Fxe-h/asscan/plugins"
)

func main() {
	start := time.Now()
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	plugins.Scan(Info)
	fmt.Printf("[*] 扫描结束,耗时: %s\n", time.Since(start))
}
