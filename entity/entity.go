package entity

import "time"


type Customer struct {
	Client     string
	Status     bool
	LogTime    time.Time
	SystemTime time.Time
	DiffTime   time.Duration
}