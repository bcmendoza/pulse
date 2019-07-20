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

// Adds metrics for all patients
func (h *Hospital) AddHospitalMetrics(metric, unitType string) {
	h.Lock()
	defer h.Unlock()

	for department := range h.Children {
		for patient := range h.Children[department].Children {
			h.Children[department].Children[patient].Children[metric] = Metric{
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
		}
	}
}

// Adds metrics for patients in a single department
func (h *Hospital) AddDepartmentMetrics(department, metric, unitType string) {
	h.Lock()
	defer h.Unlock()

	for patient := range h.Children[department].Children {
		h.Children[department].Children[patient].Children[metric] = Metric{
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
	}
}

// Adds metrics for a single patient
func (h *Hospital) AddPatientMetric(department, patient, metric, unitType string) {
	h.Lock()
	defer h.Unlock()

	h.Children[department].Children[patient].Children[metric] = Metric{
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
}
