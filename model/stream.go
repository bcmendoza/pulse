package model

import "github.com/bcmendoza/pulse/utils"

// Stream is a slice of the avg of the Pulses of n Streams
// Hospital Stream <- Departments Streams <- Patients Streams
type Stream struct {
	UnitType   string    `json:"unitType"`
	History    []Pulse   `json:"history"`
	Lower      float64   `json:"lower"`
	Upper      float64   `json:"upper"`
	Thresholds []float64 `json:"-"`
}

func MakeMetricStream(label, unitType string, lower, upper float64) Stream {
	stream := Stream{
		UnitType:   unitType,
		History:    make([]Pulse, 0),
		Lower:      lower,
		Upper:      upper,
		Thresholds: utils.CalcThresholds(lower, upper),
	}
	return stream
}
