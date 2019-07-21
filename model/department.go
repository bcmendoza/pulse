package model

type Department struct {
	Type       string             `json:"type"`
	Name       string             `json:"name"`
	Children   map[string]Patient `json:"patients"`
	Stream     Stream             `json:"stream"`
	ParentChan chan struct{}      `json:"-"`
	UpdateChan chan struct{}      `json:"-"`
}

func (h *Hospital) AddDepartment(name string) {
	h.Lock()
	defer h.Unlock()

	d := Department{
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
		ParentChan: h.UpdateChan,
		UpdateChan: make(chan struct{}, 1),
	}
	h.Children[name] = d
	go d.Subscribe()
}

func (d *Department) Subscribe() {
	for range d.UpdateChan {
		var size, sum float64
		for _, p := range d.Children {
			if s := len(p.Stream.History); s > 0 {
				size++
				sum += p.Stream.History[s-1].Score
			}
		}
		if size > 0 {
			d.Stream.History = append(
				d.Stream.History,
				MakePulse(sum/float64(size), d.Stream.Thresholds),
			)
			d.ParentChan <- struct{}{}
		}
	}
}
