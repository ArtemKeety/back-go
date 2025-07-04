package model

import "time"

type Session struct {
	Id      int
	Refresh string
	Guid    string
	Time    time.Time
	Ip      string
}
