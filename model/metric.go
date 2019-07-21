package model

// Metric represents a single numerical value
// It never has any children, always nil
type Metric struct {
	Type     string              `json:"type"`
	Name     string              `json:"name"`
	Children map[string]struct{} `json:"-"`
	Stream   Stream              `json:"stream"`
}

type MetricKey struct {
	Department string
	Patient    string
	Metric     string
}

// Adds metrics for all patients
func (h *Hospital) AddHospitalMetrics(label, unitType string, lower, upper float64) {
	h.Lock()
	defer h.Unlock()

	for d := range h.Children {
		for p := range h.Children[d].Children {
			h.Children[d].Children[p].Children[label] = Metric{
				Type:   "metric",
				Name:   label,
				Stream: MakeMetricStream(label, unitType, lower, upper),
			}
			h.MetricKeys[MetricKey{
				Department: d,
				Patient:    p,
				Metric:     label,
			}] = struct{}{}
		}
	}
}

// Adds metrics for patients in a single department
func (h *Hospital) AddDepartmentMetrics(department, label, unitType string, lower, upper float64) {
	h.Lock()
	defer h.Unlock()

	for p := range h.Children[department].Children {
		h.Children[department].Children[p].Children[label] = Metric{
			Type:     "metric",
			Name:     label,
			Children: nil,
			Stream:   MakeMetricStream(label, unitType, lower, upper),
		}
		h.MetricKeys[MetricKey{
			Department: department,
			Patient:    p,
			Metric:     label,
		}] = struct{}{}
	}
}

// Adds metrics for a single patient
func (h *Hospital) AddPatientMetric(department, patient, label, unitType string, lower, upper float64) {
	h.Lock()
	defer h.Unlock()

	h.Children[department].Children[patient].Children[label] = Metric{
		Type:     "metric",
		Name:     label,
		Children: nil,
		Stream:   MakeMetricStream(label, unitType, lower, upper),
	}
	h.MetricKeys[MetricKey{
		Department: department,
		Patient:    patient,
		Metric:     label,
	}] = struct{}{}
}

// Adds a new Pulse to a given Metric's stream
func (h *Hospital) AddMetricPulse(department, patient, metric string, value float64) {
	h.Lock()
	defer h.Unlock()

	if _, ok := h.MetricKeys[MetricKey{
		Department: department,
		Patient:    patient,
		Metric:     metric,
	}]; !ok {
		return
	}

	// test
	m := h.Children[department].Children[patient].Children[metric]
	h.Children[department].Children[patient].Children[metric] = m
}
