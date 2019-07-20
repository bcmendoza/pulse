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
	Children map[string]Department `json:"children"`
	Stream   Stream                `json:"stream"`
}

func New() *Hospital {
	return &Hospital{
		Children: make(map[string]Department),
		Stream: Stream{
			Label:    "hospital",
			UnitType: "%",
			Ratings:  make(map[string]Rating),
			Values:   make([]Pulse, 0),
		},
	}
}
