package public

type AuthRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Request struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	LocationID   string `json:"locationId"`
	ZID          string `json:"zId"`
	HistoryLimit int    `json:"historyLimit"`
}

// RingDeviceStatus represents the Device data on Ring Alarm Devices
type RingDeviceStatus struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Faulted bool   `json:"faulted"`
	Mode    string `json:"mode"`
}

type RingLockStatus struct {
	Name   string `json:"name"`
	ZID    string `json:"zid"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

type RingDeviceEvent struct {
	DeviceName string `json:"name"`
	Time       int64  `json:"time"`
	Type       string `json:"type"`
}

type Response struct {
	DeviceStatus []RingDeviceStatus `json:"deviceStatus"`
	Events       []RingDeviceEvent  `json:"events"`
}

type RingMetaData struct {
	LocationID string `json:"locationId"`
	ZID        string `json:"zId"`
}
