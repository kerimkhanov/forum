package models

import "time"

type Users struct {
	Id            string
	Email         string
	Login         string
	Password      string
	Session_token string
	TimeSessions  time.Time
}
