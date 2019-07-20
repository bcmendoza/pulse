package model

import "time"

// Stream is the average of the last Pulses of n Streams
// Hospital Stream <- Departments Streams <- Patients Streams
type Stream []Pulse

// Pulse is a single snapshot of aggregated metric data
// Its score is calculated on interval via a channel
// TODO: It should be matched to a specific Rule
type Pulse struct {
	Timestamp time.Time `json:"timestamp"`
	Score     int       `json:"score"`
	RuleName  string    `json:"-"`
}

func MakePulse(score int) Pulse {
	return Pulse{
		Timestamp: time.Now(),
		Score:     score,
	}
}

// TODO: Outcome is a uniquely identified goal based on a set of metrics
type Outcome struct {
	Name   string
	Stream []Pulse
}
