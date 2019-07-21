package model

type Department struct {
	Type     string             `json:"type"`
	Name     string             `json:"name"`
	Children map[string]Patient `json:"patients"`
	Stream   Stream             `json:"stream"`
}

func (h *Hospital) AddDepartment(name string) {
	h.Lock()
	defer h.Unlock()

	h.Children[name] = Department{
		Type:     "department",
		Name:     name,
		Children: make(map[string]Patient),
		Stream: Stream{
			UnitType:   "%",
			History:    make([]Pulse, 0),
			Lower:      0,
			Upper:      100,
			Thresholds: []float64{33.33, 66.66},
		},
	}
}
