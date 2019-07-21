package model

import (
	"fmt"

	"github.com/bcmendoza/pulse/utils"
)

// Metric represents a single numerical value
// It never has any children, always nil
type Metric struct {
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
func (h *Hospital) AddHospitalMetrics(metric, unitType string) {
	h.Lock()
	defer h.Unlock()

	for d, dept := range h.Children {
		for p, pat := range dept.Children {
			pat.Children[metric] = Metric{
				Name:     fmt.Sprintf("metric-%s", metric),
				Children: nil,
				Stream: Stream{
					Owner:    metric,
					UnitType: unitType,
					Ratings:  make(map[string]Rating),
					Current: Pulse{
						Score:     0,
						Timestamp: utils.Timestamp(),
					},
					Historical: make([]Pulse, 0),
				},
			}
			h.MetricKeys[MetricKey{
				Department: d,
				Patient:    p,
				Metric:     metric,
			}] = struct{}{}
			dept.Children[p] = pat
		}
		h.Children[d] = dept
	}
}

// Adds metrics for patients in a single department
func (h *Hospital) AddDepartmentMetrics(department, metric, unitType string) {
	h.Lock()
	defer h.Unlock()

	if dept, ok := h.Children[department]; ok {
		for p, pat := range dept.Children {
			pat.Children[metric] = Metric{
				Name:     fmt.Sprintf("metric-%s", metric),
				Children: nil,
				Stream: Stream{
					Owner:    metric,
					UnitType: unitType,
					Ratings:  make(map[string]Rating),
					Current: Pulse{
						Score:     0,
						Timestamp: utils.Timestamp(),
					},
					Historical: make([]Pulse, 0),
				},
			}
			h.MetricKeys[MetricKey{
				Department: department,
				Patient:    p,
				Metric:     metric,
			}] = struct{}{}
			dept.Children[p] = pat
		}
		h.Children[department] = dept
	}
}

// Adds metrics for a single patient
func (h *Hospital) AddPatientMetric(department, patient, metric, unitType string) {
	h.Lock()
	defer h.Unlock()

	if dept, ok := h.Children[department]; ok {
		if pat, ok := dept.Children[patient]; ok {
			pat.Children[metric] = Metric{
				Name:     fmt.Sprintf("metric-%s", metric),
				Children: nil,
				Stream: Stream{
					Owner:    metric,
					UnitType: unitType,
					Ratings:  make(map[string]Rating),
					Current: Pulse{
						Score:     0,
						Timestamp: utils.Timestamp(),
					},
					Historical: make([]Pulse, 0),
				},
			}
			h.MetricKeys[MetricKey{
				Department: department,
				Patient:    patient,
				Metric:     metric,
			}] = struct{}{}
			dept.Children[patient] = pat
		}
		h.Children[department] = dept
	}
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
