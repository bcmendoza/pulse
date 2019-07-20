package model

import "time"

// Stream is a slice of the avg of the Pulses of n Streams
// Hospital Stream <- Departments Streams <- Patients Streams
type Stream struct {
	Label    string            `json:"label"`
	UnitType string            `json:"unitType"`
	Ratings  map[string]Rating `json:"ratings"`
	Values   []Pulse           `json:"history"`
}

// Pulse is a single snapshot of data
type Pulse struct {
	Score     float64   `json:"score"`
	Timestamp time.Time `json:"timestamp"`
	RatingID  string    `json:"ratingID"`
}

func MakePulse(score float64, ratingID string) Pulse {
	return Pulse{
		Score:     score,
		Timestamp: time.Now(),
		RatingID:  ratingID,
	}
}

// Rating associates a range with label and color to eval a Pulse
// Based on Ratings defined, each Pulse should be assigned a Rating
type Rating struct {
	Label string  `json:"label"`
	Color string  `json:"color"`
	LTE   float64 `json:"lte"`
	GTE   float64 `json:"gte"`
}
