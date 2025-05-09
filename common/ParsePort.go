package common

import (
	"sort"
	"strconv"
	"strings"
)

var (
	nmapTopPortsSet map[int]struct{}
)

// init 函数在包初始化时执行，用于预处理 NmapTop1000 列表
func init() {
	nmapTopPortsSet = parseStringToPortSet(NmapTop1000)
}

// parseStringToPortSet 将端口字符串（如 "80,443,1000-1005"）解析为一个整数集合
func parseStringToPortSet(portsStr string) map[int]struct{} {
	set := make(map[int]struct{})
	if portsStr == "" {
		return set
	}
	slices := strings.Split(portsStr, ",")
	for _, portEntry := range slices {
		portEntry = strings.TrimSpace(portEntry)
		if portEntry == "" {
			continue
		}
		if strings.Contains(portEntry, "-") {
			ranges := strings.Split(portEntry, "-")
			if len(ranges) < 2 {
				continue // 格式错误的范围
			}
			startPortStr := strings.TrimSpace(ranges[0])
			endPortStr := strings.TrimSpace(ranges[1])

			startPort, errStart := strconv.Atoi(startPortStr)
			endPort, errEnd := strconv.Atoi(endPortStr)

			if errStart != nil || errEnd != nil {
				continue // 端口号转换失败
			}
			if startPort < 1 || endPort > 65535 || startPort > endPort {
				continue // 无效的端口范围
			}
			for i := startPort; i <= endPort; i++ {
				set[i] = struct{}{}
			}
		} else {
			port, err := strconv.Atoi(portEntry)
			if err == nil && port >= 1 && port <= 65535 {
				set[port] = struct{}{}
			}
		}
	}
	return set
}

// ParsePort 解析端口字符串，优先排列 Top 1000 端口并排序
func ParsePort(ports string) []int {
	var allFoundPorts []int // 存储所有初步解析到的端口，可能有重复
	// 定义一个内部递归函数来处理端口字符串的解析
	// 它可以访问外部的 PortGroup 和 allFoundPorts
	var parseInternal func(currentPortsStr string)
	parseInternal = func(currentPortsStr string) {
		if currentPortsStr == "" {
			return
		}
		slices := strings.Split(currentPortsStr, ",")
		for _, portStr := range slices {
			portStr = strings.TrimSpace(portStr)
			if portStr == "" {
				continue
			}
			// 检查 portStr 是否为 PortGroup 中的键
			if groupDefinition, ok := PortGroup[portStr]; ok {
				parseInternal(groupDefinition) // 递归解析组定义
				continue
			}
			// 处理单个端口或范围
			var startPort, endPort int
			var err error
			if strings.Contains(portStr, "-") {
				ranges := strings.Split(portStr, "-")
				if len(ranges) < 2 {
					continue // 格式错误的范围
				}
				rawStartPort := strings.TrimSpace(ranges[0])
				rawEndPort := strings.TrimSpace(ranges[1])

				startPort, err = strconv.Atoi(rawStartPort)
				if err != nil {
					continue // 转换失败
				}
				endPort, err = strconv.Atoi(rawEndPort)
				if err != nil {
					continue // 转换失败
				}
				if startPort > endPort { // 确保 startPort <= endPort
					startPort, endPort = endPort, startPort
				}
			} else {
				startPort, err = strconv.Atoi(portStr)
				if err != nil {
					continue // 转换失败
				}
				endPort = startPort // 单个端口
			}
			// 添加有效端口到 allFoundPorts
			for i := startPort; i <= endPort; i++ {
				if i >= 1 && i <= 65535 {
					allFoundPorts = append(allFoundPorts, i)
				}
			}
		}
	}
	parseInternal(ports) // 开始解析输入的端口字符串
	// 1. 去重
	uniquePorts := removeDuplicate(allFoundPorts)
	// 2. 分类
	var topPriorityPorts []int
	var otherPorts []int
	for _, p := range uniquePorts {
		if _, isTopPort := nmapTopPortsSet[p]; isTopPort {
			topPriorityPorts = append(topPriorityPorts, p)
		} else {
			otherPorts = append(otherPorts, p)
		}
	}
	// 3. 排序
	sort.Ints(topPriorityPorts)
	sort.Ints(otherPorts)
	// 4. 合并
	finalResultPorts := append(topPriorityPorts, otherPorts...)
	return finalResultPorts
}

// removeDuplicate 函数保持不变，用于去除整数切片中的重复项
func removeDuplicate(old []int) []int {
	if len(old) == 0 {
		return []int{}
	}
	// 预估容量以提高效率
	result := make([]int, 0, len(old))
	temp := make(map[int]struct{}, len(old))
	for _, item := range old {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
