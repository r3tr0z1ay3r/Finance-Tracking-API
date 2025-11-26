package Model

import (
	"encoding/json"
	"time"
)

// customTime wraps time.Time -> JSON format
type customTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05" //Time format which is to be followed

func (ct *customTime) unMarshalJSON(b []byte) error {

	s := string(b)
	s = s[1 : len(b)-1]
	t, err := time.Parse(ctLayout, s)
	if err != nil {

		return err

	}
	ct.Time = t
	return nil

}

func (ct customTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Format(ctLayout))
}

type Transaction struct {
	time customTime `json:"time"`
	amt  float64    `json:"user"`
	flow string     `json:"flow"`
	mode string     `json:"mode"`
}
