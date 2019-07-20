package model

import "github.com/bcmendoza/pulse/utils"

// Metric represents a single numerical value
// It never has any children, always nil
type Metric struct {
	Children map[string]struct{} `json:"-"`
	Stream   Stream              `json:"stream"`
}

func (d *Department) AddMetric(owner, unitType string) {
	d.Children[owner] = Patient{
		Children: nil,
		Stream: Stream{
			Owner:    owner,
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
