package model

type Department struct {
	Children map[string]Patient `json:"patients"`
	Stream   Stream             `json:"stream"`
}

func (h *Hospital) AddDepartment(label string) {
	h.Lock()
	defer h.Unlock()

	h.Children[label] = Department{
		Children: make(map[string]Patient),
		Stream: Stream{
			Label:    label,
			UnitType: "%",
			Ratings:  make(map[string]Rating),
			Values:   make([]Pulse, 0),
		},
	}
}
