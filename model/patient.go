package model

type Patient struct {
	Type       string            `json:"type"`
	Name       string            `json:"name"`
	Children   map[string]Metric `json:"children"`
	Stream     Stream            `json:"stream"`
	ParentChan chan struct{}     `json:"-"`
	UpdateChan chan struct{}     `json:"-"`
}

type PatientKey struct {
	Department string
	Patient    string
}

func (h *Hospital) AddPatient(department, patient string) {
	h.Lock()
	defer h.Unlock()

	p := Patient{
		Type:     "patient",
		Name:     patient,
		Children: make(map[string]Metric),
		Stream: Stream{
			UnitType:   "%",
			History:    make([]Pulse, 0),
			Lower:      0,
			Upper:      100,
			Thresholds: []float64{33.33, 66.66},
		},
		ParentChan: h.Children[department].UpdateChan,
		UpdateChan: make(chan struct{}, 1),
	}
	h.Children[department].Children[patient] = p
	go p.Subscribe()
}

func (p *Patient) Subscribe() {
	for range p.UpdateChan {
		var size, sum float64
		for _, m := range p.Children {
			if m.Percent != 0 {
				size++
				sum += m.Percent
			}
		}
		p.Stream.History = append(
			p.Stream.History,
			MakePulse(sum/float64(size), p.Stream.Thresholds),
		)
		p.ParentChan <- struct{}{}
	}
}
