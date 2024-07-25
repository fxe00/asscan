package plugins

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/Fxe-h/asscan/common"
	"github.com/cheggaaa/pb"
)

type Addr struct {
	ip   string
	port int
}

func PortScan(hostslist []string, ports string, timeout int64) []string {
	var AliveAddress []string
	probePorts := common.ParsePort(ports)
	if len(probePorts) == 0 {
		fmt.Printf("[-] parse port %s error, please check your port format\n", ports)
		return AliveAddress
	}

	workers := common.Threads
	totalTasks := len(hostslist) * len(probePorts)
	bar := pb.StartNew(totalTasks).Prefix("Scanning")

	Addrs := make(chan Addr, totalTasks)
	results := make(chan string, totalTasks)
	var wg sync.WaitGroup

	// 创建文件以存储结果
	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Failed to create result file:", err)
		return AliveAddress
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// 接收结果
	go func() {
		for found := range results {
			AliveAddress = append(AliveAddress, found)
			_, _ = writer.WriteString(found + "\n") // 写入文件
			_ = writer.Flush()                       // 刷新缓冲区
			wg.Done()
		}
	}()

	// 多线程扫描
	for i := 0; i < workers; i++ {
		go func() {
			for addr := range Addrs {
				PortConnect(addr, results, timeout, &wg)
				wg.Done()
				bar.Increment() // 更新进度条
			}
		}()
	}

	// 添加扫描目标
	for _, port := range probePorts {
		for _, host := range hostslist {
			wg.Add(1)
			Addrs <- Addr{host, port}
		}
	}
	wg.Wait()
	close(Addrs)
	close(results)
	bar.FinishPrint("Done!") // 完成进度条

	return AliveAddress
}

func PortConnect(addr Addr, respondingHosts chan<- string, adjustedTimeout int64, wg *sync.WaitGroup) {
	host, port := addr.ip, addr.port
	conn, err := common.WrapperTcpWithTimeout("tcp4", fmt.Sprintf("%s:%v", host, port), time.Duration(adjustedTimeout)*time.Second)
	if err == nil {
		defer conn.Close()
		address := host + ":" + strconv.Itoa(port)
		fmt.Printf("[*] %s open\n", address)
		wg.Add(1)
		respondingHosts <- address
	}
}