package model

import (
	"fmt"

	"github.com/bcmendoza/pulse/utils"
)

type Department struct {
	Name     string             `json:"name"`
	Children map[string]Patient `json:"patients"`
	Stream   Stream             `json:"stream"`
}

func (h *Hospital) AddDepartment(name string) {
	h.Lock()
	defer h.Unlock()

	h.Children[name] = Department{
		Name:     fmt.Sprintf("department-%s", name),
		Children: make(map[string]Patient),
		Stream: Stream{
			Owner:    name,
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
