package model

import (
	"sync"

	"github.com/bcmendoza/pulse/utils"
)

// 1. Random patient metrics bubbling up
// 2. REST API for getting metrics
// 3. REST APi for measurement fields
// 4. REST API for setting metrics via fields
// 5. Multiple outcomes per department
// 6. Multiple outcomes per patient

type Hospital struct {
	sync.Mutex
	Name        string                  `json:"name"`
	Children    map[string]Department   `json:"children"`
	Stream      Stream                  `json:"stream"`
	PatientKeys map[PatientKey]struct{} `json:"-"`
	MetricKeys  map[MetricKey]struct{}  `json:"-"`
}

func New() *Hospital {
	return &Hospital{
		Name:     "hospital",
		Children: make(map[string]Department),
		Stream: Stream{
			Owner:    "hospital",
			UnitType: "%",
			Ratings:  make(map[string]Rating),
			Current: Pulse{
				Score:     0,
				Timestamp: utils.Timestamp(),
			},
			Historical: make([]Pulse, 0),
		},
	}
}
