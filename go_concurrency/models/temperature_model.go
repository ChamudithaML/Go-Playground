package models

import "time"

type Temperature struct {
	SensorID    string    `json:"sensor_id"`
	Timestamp   time.Time `json:"timestamp"`
	Temperature float64   `json:"temperature"`
	Alert       bool      `json:"alert"`
}
