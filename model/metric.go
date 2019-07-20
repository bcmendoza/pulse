package model

import "github.com/bcmendoza/pulse/utils"

// Metric represents a single numerical value
// It never has any children, always nil
type Metric struct {
	Children map[string]struct{} `json:"-"`
	Stream   Stream              `json:"stream"`
}

func (h *Hospital) AddMetric(department, patient, metric, unitType string) {
	h.Lock()
	defer h.Unlock()

	h.Children[department].Children[patient].Children[metric] = Metric{
		Children: nil,
		Stream: Stream{
			Owner:    metric,
			UnitType: unitType,
			Ratings:  make(map[string]Rating),
			Current: Pulse{
				Score:     0,
				Timestamp: utils.Timestamp(),
			},
			Historical: make([]Pulse, 0),
		},
	}
}
