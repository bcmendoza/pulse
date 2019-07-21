package model

import "github.com/bcmendoza/pulse/utils"

// Pulse is a single snapshot of data
type Pulse struct {
	Score     float64 `json:"score"`
	Timestamp int64   `json:"timestamp"`
	Rating    int     `json:"rating"`
}

func MakePulse(score float64, thresholds []float64) Pulse {
	pulse := Pulse{
		Score:     score,
		Timestamp: utils.Timestamp(),
	}
	if score > thresholds[1] {
		pulse.Rating = 1
	} else if score > thresholds[0] {
		pulse.Rating = 2
	} else {
		pulse.Rating = 3
	}

	return pulse
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

	m := h.Children[department].Children[patient].Children[metric]
	m.Stream.History = append(m.Stream.History, MakePulse(value, m.Stream.Thresholds))
	h.Children[department].Children[patient].Children[metric] = m

	percent := utils.CalcRelativePercent(value, m.Stream.Upper, m.Stream.Lower)
	m.Percent = percent
	m.ParentChan <- struct{}{}
}
