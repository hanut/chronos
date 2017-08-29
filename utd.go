package main

import (
	"strconv"
	"time"
)

type Utd struct {
	Phone     string    `json:"phone"`
	Balance   float32   `json:"balance"`
	Cpm       int       `json:"cpm"`
	StartedAt time.Time `json:"startedAt"`
}

func (user Utd) String() string {
	var stringVal string = ""
	stringVal += "Phone:\t" + user.Phone + "\n"
	stringVal += "Balance:\t" + strconv.FormatFloat(float64(user.Balance), 'g', 1, 32) + "\n"
	stringVal += "CPM:\t" + strconv.Itoa(user.Cpm) + "\n"
	stringVal += "Started At:\t" + user.StartedAt.String() + "\n"
	return stringVal
}
