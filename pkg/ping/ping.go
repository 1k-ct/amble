// aliveパッケージは、サーバーにpingを送信し監視します。
// 主に使う関数は、AliveMonitoring, AliveMonitoringForWindows, AliveMonitoringForOther
package alive

import (
	"runtime"

	"github.com/go-ping/ping"
)

type Monitorer interface {
	monitoring(addr string, pingCount int) (*ping.Statistics, error)
}
type aliveMonitorer struct {
	monitor   Monitorer
	addr      string
	pingCount int
}

// AliveMonitoring(addr string, pingCount int) (*ping.Statistics, error)
// サーバーにpingを打ち、その統計情報を返すための関数です
// addr string -> 送信するIPアドレス ex addr="192.168.11.1", addr="www.example.com"
// pingCount int -> pingする回数 ex pingCount=4
//
// > If the OS is Windows, call `AliveMonitoringForWindows`; otherwise, call `AliveMonitoringForOther`
// A function that is used to ping a server and return the statistics of the ping.
func AliveMonitoring(addr string, pingCount int) (statistics *ping.Statistics, err error) {
	os := runtime.GOOS
	if os == "windows" {
		return AliveMonitoringForWindows(addr, pingCount)
	}
	return AliveMonitoringForOther(addr, pingCount)
}

func (a *aliveMonitorer) aliveMonitoring() (*ping.Statistics, error) {
	status, err := a.monitor.monitoring(a.addr, a.pingCount)
	return status, err
}

type forWindows struct{}

// A function that is used to ping a server and return the statistics of the ping.
func (f forWindows) monitoring(addr string, pingCount int) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return nil, err
	}
	pinger.SetPrivileged(true)
	pinger.Count = pingCount
	if err := pinger.Run(); err != nil {
		return nil, err
	}
	stats := pinger.Statistics()
	return stats, nil
}

// AliveMonitoringForWindows(addr string, pingCount int) (*ping.Statistics, error)
// サーバーにpingを打ち、その統計情報を返すためのWindows用関数です
// この関数は自分のosがwindowsの場合に使用すます。相手のサーバーのosはなんでもいい
// addr string -> 送信するIPアドレス ex addr="192.168.11.1", addr="www.example.com"
// pingCount int -> pingする回数 ex pingCount=4
//
// A function that checks whether the server is alive or not.
// A function for windows that is used to ping a server and return the statistics of the ping.
func AliveMonitoringForWindows(addr string, pingCount int) (*ping.Statistics, error) {
	a := &aliveMonitorer{monitor: &forWindows{}, addr: addr, pingCount: pingCount}
	status, err := a.aliveMonitoring()
	return status, err
}

type forOther struct{}

// A function that is used to ping a server and return the statistics of the ping.
func (f forOther) monitoring(addr string, pingCount int) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return nil, err
	}
	pinger.Count = pingCount
	if err := pinger.Run(); err != nil {
		return nil, err
	}
	stats := pinger.Statistics()
	return stats, nil
}

// AliveMonitoringForOther(addr string, pingCount int) (*ping.Statistics, error)
// サーバーにpingを打ち、その統計情報を返すための関数です
// この関数は自分のosが特定できない場合に使用すます。相手のサーバーのosはなんでもいい
// addr string -> 送信するIPアドレス ex addr="192.168.11.1", addr="www.example.com"
// pingCount int -> pingする回数 ex pingCount=4
//
// A function that checks whether the server is alive or not.
// A function that is used to ping a server and return the statistics of the ping.
func AliveMonitoringForOther(addr string, pingCount int) (*ping.Statistics, error) {
	a := &aliveMonitorer{monitor: &forOther{}, addr: addr, pingCount: pingCount}
	status, err := a.aliveMonitoring()
	return status, err
}
