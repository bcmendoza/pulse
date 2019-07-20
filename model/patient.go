package model

import "github.com/bcmendoza/pulse/utils"

// Attrs are inconsequential patient data
// They are added to the Patient struct in unstructured JSON
type Attrs map[string]string

type Patient struct {
	Children map[string]Metric `json:"children"`
	Stream   Stream            `json:"stream"`
	Attrs
}

func (h *Hospital) AddPatient(department, patient string) {
	h.Lock()
	defer h.Unlock()

	h.Children[department].Children[patient] = Patient{
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
