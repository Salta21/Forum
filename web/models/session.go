package models

import "time"

type Session struct {
	Session_id int
	User_id    int
	Token      string
	Expiry     time.Time
}
