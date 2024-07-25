package common

import (
	"net"
	"time"
)

func WrapperTcpWithTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := &net.Dialer{Timeout: timeout}
	return WrapperTCP(network, address, d)
}

func WrapperTCP(network, address string, forward *net.Dialer) (net.Conn, error) {
	// 使用 forward Dialer 直接创建连接
	conn, err := forward.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
