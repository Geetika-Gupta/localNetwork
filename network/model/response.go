package model

//ListDevices : device list response model
type ListDevices struct {
	[]Device
}

//Device : device model
type Device struct {
	Type string `json:"type"`
	Name string `json:"name"`
}