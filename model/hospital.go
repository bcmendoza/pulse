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
	UpdateChan chan Update            `json:"-"`
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
		UpdateChan: make(chan Update, 1),
	}
	go h.Subscribe()
	return h
}

func (h *Hospital) Subscribe() {
	for u := range h.UpdateChan {
		if p, ok := h.Children[u.Department].Children[u.Patient]; ok {
			var size, sum float64
			for _, m := range p.Children {
				if m.Percent != 0 {
					size++
					sum += m.Percent
				}
			}
			if size > 0 {
				p.Stream.History = append(p.Stream.History, MakePulse(sum/float64(size), p.Stream.Thresholds))
				h.Children[u.Department].Children[u.Patient] = p
			}
		}
		if d, ok := h.Children[u.Department]; ok {
			var size, sum float64
			for _, p := range d.Children {
				if s := len(p.Stream.History); s > 0 {
					size++
					sum += p.Stream.History[s-1].Score
				}
			}
			if size > 0 {
				d.Stream.History = append(d.Stream.History, MakePulse(sum/float64(size), d.Stream.Thresholds))
				h.Children[u.Department] = d
			}
		}
		var size, sum float64
		for _, d := range h.Children {
			if s := len(d.Stream.History); s > 0 {
				size++
				sum += d.Stream.History[s-1].Score
			}
		}
		if size > 0 {
			h.Stream.History = append(h.Stream.History, MakePulse(sum/float64(size), h.Stream.Thresholds))
		}
	}
}
