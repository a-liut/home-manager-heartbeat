package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"home_manager_heartbeat/pkg/model"
	"net/http"
	"strconv"
)

type DevicesRepository struct {
	baseUrl string
}

/**
 * Get all registered devices
 */
func (repository DevicesRepository) GetAll() ([]*model.Device, error) {
	url := fmt.Sprintf("%s/devices", repository.baseUrl)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var devices []*model.Device
	if err := json.NewDecoder(resp.Body).Decode(&devices); err != nil {
		return nil, err
	}

	return devices, nil
}

/**
 * Get a device by ID
 */
func (repository DevicesRepository) GetById(id string) (*model.Device, error) {
	url := fmt.Sprintf("%s/devices/%s", repository.baseUrl, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var device model.Device
	if err := json.NewDecoder(resp.Body).Decode(device); err != nil {
		return nil, err
	}

	return &device, nil
}

/**
 * Update a device
 */
func (repository DevicesRepository) Update(device *model.Device) error {
	url := fmt.Sprintf("%s/devices/%s", repository.baseUrl, device.Id)

	body, err := json.Marshal(map[string]string{
		"online": strconv.FormatBool(device.Online),
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return err
}

func NewDevicesRepository(url string) *DevicesRepository {
	return &DevicesRepository{baseUrl: url}
}
