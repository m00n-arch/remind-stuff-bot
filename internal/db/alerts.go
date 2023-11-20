package db

import (
	"time"
)

type Alert struct {
	date      time.Time
	content   string
	userID    string
	messageID int
	period    time.Duration
}

type alertsDB struct {
}

func accurate() {
	time.Parse("01 2006 / January", "")
}
