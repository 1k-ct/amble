package alive

import (
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

func (a *aliveMonitorer) aliveMonitoring() (*ping.Statistics, error) {
	status, err := a.monitor.monitoring(a.addr, a.pingCount)
	return status, err
}

type forWindows struct{}

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

// A function for windows that is used to ping a server and return the statistics of the ping.
func AliveMonitoringForWindows(addr string, pingCount int) (*ping.Statistics, error) {
	a := &aliveMonitorer{monitor: &forWindows{}, addr: addr, pingCount: pingCount}
	status, err := a.aliveMonitoring()
	return status, err
}

type forOther struct{}

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

// A function for other that is used to ping a server and return the statistics of the ping.
func AliveMonitoringForOther(addr string, pingCount int) (*ping.Statistics, error) {
	a := &aliveMonitorer{monitor: &forOther{}, addr: addr, pingCount: pingCount}
	status, err := a.aliveMonitoring()
	return status, err
}
