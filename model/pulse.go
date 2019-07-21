package model

import "github.com/bcmendoza/pulse/utils"

// Pulse is a single snapshot of data
type Pulse struct {
	Score     float64 `json:"score"`
	Timestamp int64   `json:"timestamp"`
	Rating    int     `json:"rating"`
}

func MakePulse(score float64, thresholds []float64) Pulse {
	pulse := Pulse{
		Score:     score,
		Timestamp: utils.Timestamp(),
	}
	if score > thresholds[1] {
		pulse.Rating = 1
	} else if score > thresholds[0] {
		pulse.Rating = 2
	} else {
		pulse.Rating = 3
	}

	return pulse
}
