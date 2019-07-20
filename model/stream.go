package model

import "github.com/bcmendoza/pulse/utils"

// Stream is a slice of the avg of the Pulses of n Streams
// Hospital Stream <- Departments Streams <- Patients Streams
type Stream struct {
	Owner      string            `json:"owner"`
	UnitType   string            `json:"unitType"`
	Ratings    map[string]Rating `json:"ratings"`
	Current    Pulse             `json:"current"`
	Historical []Pulse           `json:"historical"`
}

// Pulse is a single snapshot of data
type Pulse struct {
	Score       float64 `json:"score"`
	Timestamp   int64   `json:"timestamp"`
	RatingLabel string  `json:"ratingLabel"`
}

func MakePulse(score float64, ratingLabel string) Pulse {
	return Pulse{
		Score:       score,
		Timestamp:   utils.Timestamp(),
		RatingLabel: ratingLabel,
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
