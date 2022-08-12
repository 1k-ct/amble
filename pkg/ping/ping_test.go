package alive

import (
	"net"
	"testing"
	"time"

	"github.com/go-ping/ping"
)

type forWindowsMock struct{}

func (m *forWindowsMock) monitoring(addr string, pingCount int) (*ping.Statistics, error) {
	status := &ping.Statistics{
		PacketsRecv:           0,
		PacketsSent:           0,
		PacketsRecvDuplicates: 0,
		PacketLoss:            0,
		IPAddr: &net.IPAddr{
			IP:   []byte{},
			Zone: "",
		},
		Addr:      addr,
		Rtts:      []time.Duration{},
		MinRtt:    0,
		MaxRtt:    0,
		AvgRtt:    0,
		StdDevRtt: 0,
	}
	return status, nil
}
func TestAliveMonitoring(t *testing.T) {
	address := "000,00,0,1"
	// address := "www.google.com"
	pingCount := 4
	a := &aliveMonitorer{monitor: &forWindowsMock{}, addr: address, pingCount: pingCount}
	status, err := a.aliveMonitoring()
	if status.Addr != address {
		t.Fatal("aliveMonitoring() mock error, and monitoring() error")
	}
	if err != nil {
		t.Fatal(err)
	}
}

// func TestAliveMonitoringWindows(t *testing.T) {
// 	// address := "192.168.11.6"
// 	address := "www.google.com"
// 	pingCount := 4
// 	status, err := AliveMonitoringWindows(address, pingCount)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	log.Println(status)
// }
