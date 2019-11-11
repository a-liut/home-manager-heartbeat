package main

import (
	"home_manager_heartbeat/pkg/checker"
	"home_manager_heartbeat/pkg/repository"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	CheckerInterval   = 20 * time.Second
	DevicesBaseUrlKey = "DEVICES_BASEURL"
)

func main() {
	log.Println("Starting Home manager Heartbeat service...")

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Start checker
	go startHeartbeatChecker(stopChan)

	<-stopChan

	log.Println("Home manager Heartbeat service ended.")
}

func startHeartbeatChecker(stopChan chan os.Signal) {
	baseUrl := os.Getenv(DevicesBaseUrlKey)
	log.Printf("Using Devices base url: %s", baseUrl)

	ticker := time.NewTicker(CheckerInterval)

	for {
		select {
		case <-ticker.C:
			// Perform check
			go func() {
				// Get a DeviceRepository
				repo := repository.NewDevicesRepository(baseUrl)

				// Start a new Checker
				heartbeatChecker := checker.NewHeartbeatChecker(repo)
				heartbeatChecker.Start()
			}()
		case <-stopChan:
			// Stop

			ticker.Stop()
			break
		}
	}
}
