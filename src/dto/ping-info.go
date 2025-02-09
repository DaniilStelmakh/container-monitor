package dto

import "time"

// PingInfo структура для хранения информации о пинге
type PingInfo struct {
	Ip       string    `json:"ip_address"`
	PingTime float64   `json:"ping_time"`
	LastSeen time.Time `json:"last_seen"`
}
