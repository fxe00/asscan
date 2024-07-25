package lib

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/Fxe-h/asscan/common"
)

var (
	Client           *http.Client
	ClientNoRedirect *http.Client
	dialTimout       = 5 * time.Second
	keepAlive        = 5 * time.Second
)

func Inithttp() {
	if common.WebTimeout == 0 {
		common.WebTimeout = 5
	}
	err := InitHttpClient(time.Duration(common.WebTimeout) * time.Second)
	if err != nil {
		panic(err)
	}
}

func InitHttpClient(Timeout time.Duration) error {
	dialer := &net.Dialer{
		Timeout:   dialTimout,
		KeepAlive: keepAlive,
	}

	tr := &http.Transport{
		DialContext:         dialer.DialContext,
		MaxConnsPerHost:     5,
		MaxIdleConns:        0,
		IdleConnTimeout:     keepAlive,
		TLSClientConfig:     &tls.Config{MinVersion: tls.VersionTLS10, InsecureSkipVerify: true},
		TLSHandshakeTimeout: 5 * time.Second,
		DisableKeepAlives:   false,
	}
	Client = &http.Client{
		Transport: tr,
		Timeout:   Timeout,
	}
	ClientNoRedirect = &http.Client{
		Transport:     tr,
		Timeout:       Timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}
	return nil
}
