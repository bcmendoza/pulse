package model

import "github.com/bcmendoza/pulse/utils"

// Stream is a slice of the avg of the Pulses of n Streams
// Hospital Stream <- Departments Streams <- Patients Streams
type Stream struct {
	Owner      string    `json:"owner"`
	UnitType   string    `json:"unitType"`
	History    []Pulse   `json:"history"`
	Lower      float64   `json:"lower"`
	Upper      float64   `json:"upper"`
	Thresholds []float64 `json:"-"`
}

func MakeMetricStream(label, unitType string, lower, upper float64) Stream {
	stream := Stream{
		Owner:    label,
		UnitType: unitType,
		History:  make([]Pulse, 0),
		Lower:    lower,
		Upper:    upper,
	}
	thirdAmt := (upper - lower) / 3
	thresh1 := lower + thirdAmt
	thresh2 := thresh1 + thirdAmt
	stream.Thresholds = []float64{thresh1, thresh2}
	return stream
}

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
