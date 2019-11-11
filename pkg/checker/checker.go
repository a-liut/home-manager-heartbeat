package checker

import (
	"home_manager_heartbeat/pkg/model"
	"home_manager_heartbeat/pkg/repository"
	"log"
	"net/http"
	"sync"
)

type HeartbeatChecker struct {
	devicesRepository *repository.DevicesRepository
}

/**
 * This method retrieves all registered devices and tries to connect to each of them
 * to check if they are available or not. Then, it updates the DB document with their current status.
 */
func (checker *HeartbeatChecker) Start() {
	log.Println("Heartbeat checker started")

	devices, err := checker.devicesRepository.GetAll()
	if err != nil {
		log.Println("Cannot get devices: ", err)
		return
	}

	log.Printf("Checking %d devices...", len(devices))

	var wg sync.WaitGroup
	wg.Add(len(devices))

	for _, device := range devices {
		go func() {
			if err := checker.checkDevice(device); err != nil {
				log.Printf("Cannot check device %s: %s\n", device, err)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	defer log.Println("Heartbeat checker ended")
}

/**
 * Checks the status of a Device.
 */
func (checker *HeartbeatChecker) checkDevice(device *model.Device) error {
	log.Printf("Checking device %s...\n", device)

	res, err := http.Get(device.HeartbeatUrl)
	if err != nil {
		log.Printf("Error checking device %s: %s", device, err)
	}

	online := err == nil && res.StatusCode == http.StatusOK
	log.Printf("Device %s: %t -> %t\n", device, device.Online, online)

	if device.Online == online {
		log.Println("Skipping update: same online status")
		return nil
	}

	// Update online status
	device.Online = online

	return checker.devicesRepository.Update(device)
}

func NewHeartbeatChecker(devicesRepository *repository.DevicesRepository) *HeartbeatChecker {
	return &HeartbeatChecker{devicesRepository: devicesRepository}
}
