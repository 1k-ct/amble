package example

import (
	"log"

	alive "github.com/1k-ct/amble/pkg/ping"
)

func ExampleAliveMonitoring() {
	addr := "www.example.com"
	pingCount := 5
	statistics, err := alive.AliveMonitoring(addr, pingCount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(statistics.PacketLoss)
}

func ExampleAliveMonitoringForWindows() {
	addr := "www.example.com"
	pingCount := 5
	statistics, err := alive.AliveMonitoringForWindows(addr, pingCount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(statistics.PacketLoss)
}

func ExampleAliveMonitoringForOther() {
	addr := "www.example.com"
	pingCount := 5
	statistics, err := alive.AliveMonitoringForOther(addr, pingCount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(statistics.PacketLoss)
}
