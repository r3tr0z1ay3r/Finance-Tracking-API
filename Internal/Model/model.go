package Model

import (
	"time"
)

type Transaction struct {
	ID   int       `json:"id"`
	Time time.Time `json:"time"`
	Amt  float64   `json:"amt"`
	Flow string    `json:"flow"`
	Mode string    `json:"mode"`
}
