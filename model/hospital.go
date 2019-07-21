package model

import (
	"sync"
)

// 1. Random patient metrics bubbling up
// 2. REST API for getting metrics
// 3. REST APi for measurement fields
// 4. REST API for setting metrics via fields
// 5. Multiple outcomes per department
// 6. Multiple outcomes per patient

type Hospital struct {
	sync.Mutex
	Type        string                  `json:"type"`
	Children    map[string]Department   `json:"children"`
	Stream      Stream                  `json:"stream"`
	PatientKeys map[PatientKey]struct{} `json:"-"`
	MetricKeys  map[MetricKey]struct{}  `json:"-"`
}

func New() *Hospital {
	return &Hospital{
		Type:     "hospital",
		Children: make(map[string]Department),
		Stream: Stream{
			UnitType:   "%",
			History:    make([]Pulse, 0),
			Lower:      0,
			Upper:      100,
			Thresholds: []float64{33.33, 66.66},
		},
		PatientKeys: make(map[PatientKey]struct{}),
		MetricKeys:  make(map[MetricKey]struct{}),
	}
}
