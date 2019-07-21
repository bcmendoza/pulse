package model

import (
	"sync"
)

type Hospital struct {
	sync.Mutex
	Type       string                 `json:"type"`
	Children   map[string]Department  `json:"children"`
	Stream     Stream                 `json:"stream"`
	MetricKeys map[MetricKey]struct{} `json:"-"`
	UpdateChan chan struct{}          `json:"-"`
}

func New() *Hospital {
	h := &Hospital{
		Type:     "hospital",
		Children: make(map[string]Department),
		Stream: Stream{
			UnitType:   "%",
			History:    make([]Pulse, 0),
			Lower:      0,
			Upper:      100,
			Thresholds: []float64{33.33, 66.66},
		},
		MetricKeys: make(map[MetricKey]struct{}),
		UpdateChan: make(chan struct{}, 1),
	}
	go h.Subscribe()
	return h
}

func (h *Hospital) Subscribe() {
	for range h.UpdateChan {
		var size, sum float64
		for _, d := range h.Children {
			if s := len(d.Stream.History); s > 0 {
				size++
				sum += d.Stream.History[s-1].Score
			}
		}
		if size > 0 {
			h.Stream.History = append(
				h.Stream.History,
				MakePulse(sum/float64(size), h.Stream.Thresholds),
			)
		}
	}
}
