package model

import (
	"fmt"

	"github.com/bcmendoza/pulse/utils"
)

type Patient struct {
	Name     string            `json:"name"`
	Children map[string]Metric `json:"children"`
	Stream   Stream            `json:"stream"`
}

type PatientKey struct {
	Department string
	Patient    string
}

func (h *Hospital) AddPatient(department, patient string) {
	h.Lock()
	defer h.Unlock()

	h.Children[department].Children[patient] = Patient{
		Name:     fmt.Sprintf("patient-%s", patient),
		Children: make(map[string]Metric),
		Stream: Stream{
			Owner:    patient,
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
