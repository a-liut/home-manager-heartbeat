package model

import "fmt"

type Device struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Address      string   `json:"address"`
	Data         []string `json:"data"`
	Online       bool     `json:"online"`
	HeartbeatUrl string   `json:"heartbeat_url"`
}

func (d Device) String() string {
	return fmt.Sprintf("%s (%s)", d.Name, d.Id)
}

type Data struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

func (d Data) String() string {
	return fmt.Sprintf("(%s, %s, %s)", d.Name, d.Value, d.Unit)
}
