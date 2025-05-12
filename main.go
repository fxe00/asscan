package main

import (
	"log"
	"runtime"
	"time"

	"github.com/Fxe-h/asscan/common"
	"github.com/Fxe-h/asscan/plugins"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	start := time.Now()
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	plugins.Scan(Info)
	log.Printf("[*] 扫描结束,耗时: %s\n", time.Since(start))
}
