package model

// Metric represents a single numerical value
// It never has any children, always nil
type Metric struct {
	Children map[string]struct{} `json:"-"`
	Stream   Stream              `json:"stream"`
}

func (d *Department) AddMetric(label, unitType string) {
	d.Children[label] = Patient{
		Children: nil,
		Stream: Stream{
			Label:    label,
			UnitType: unitType,
			Ratings:  make(map[string]Rating),
			Values:   make([]Pulse, 0),
		},
	}
}
