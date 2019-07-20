package model

import "github.com/bcmendoza/pulse/utils"

type Department struct {
	Children map[string]Patient `json:"patients"`
	Stream   Stream             `json:"stream"`
}

func (h *Hospital) AddDepartment(owner string) {
	h.Lock()
	defer h.Unlock()

	h.Children[owner] = Department{
		Children: make(map[string]Patient),
		Stream: Stream{
			Owner:    owner,
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
