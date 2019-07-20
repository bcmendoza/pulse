package model

// Attrs are inconsequential patient data
// They are added to the Patient struct in unstructured JSON
type Attrs map[string]string

type Patient struct {
	Children map[string]Metric `json:"children"`
	Stream   Stream            `json:"stream"`
	Attrs
}

func (d *Department) AddPatient(label string) {
	d.Children[label] = Patient{
		Children: make(map[string]Metric),
		Stream: Stream{
			Label:    label,
			UnitType: "%",
			Ratings:  make(map[string]Rating),
			Values:   make([]Pulse, 0),
		},
	}
}
