package model

import (
	"fmt"
	"time"

	"github.com/bcmendoza/pulse/utils"
)

type TestData struct {
	Label    string
	UnitType string
	Lower    float64
	Upper    float64
}

var allMetrics = []TestData{
	TestData{
		Label:    "Readmission Rate",
		UnitType: "%",
		Lower:    1,
		Upper:    40,
	},
	TestData{
		Label:    "Patient Wait-Time",
		UnitType: "minutes",
		Lower:    3,
		Upper:    20,
	},
	TestData{
		Label:    "Staff-to-Patient Ratio",
		UnitType: "per patient",
		Lower:    0.01,
		Upper:    1.0,
	},
	TestData{
		Label:    "Total Patients Admitted",
		UnitType: "patients",
		Lower:    100,
		Upper:    5000,
	},
	TestData{
		Label:    "Average Length of Stay",
		UnitType: "day",
		Lower:    1,
		Upper:    8,
	},
}

var deptMetrics = map[string][]TestData{
	"Critical Care: SICU":       []TestData{},
	"Maternal & Fetal Medicine": []TestData{},
	"Neuro/Medical":             []TestData{},
	"Radiation Oncology":        []TestData{},
	"Ortho/Surgical":            []TestData{},
	"Rehab":                     []TestData{},
	"PCU":                       []TestData{},
}

func (h *Hospital) LoadTestSchemas() {
	for k := range deptMetrics {
		h.AddDepartment(k)
		for i := 0; i < 2; i++ {
			h.AddPatient(k, utils.UUID())
		}
	}
	for _, m := range allMetrics {
		h.AddHospitalMetrics(m.Label, m.UnitType, m.Lower, m.Upper)
	}
}

func (h *Hospital) RunGenerator() {
TIMED_LOOP:
	for timeout := time.After(time.Minute * 5); ; {
		select {
		case <-timeout:
			break TIMED_LOOP
		default:
			for k := range h.MetricKeys {
				if m, ok := h.Children[k.Department].Children[k.Patient].Children[k.Metric]; ok {
					h.AddMetricPulse(k.Department, k.Patient, k.Metric, utils.Random(m.Stream.Lower, m.Stream.Upper))
				}
				d, err := time.ParseDuration(fmt.Sprintf("%ds", 2))
				time.Sleep(d)
				if err != nil {
				}
			}
			d, err := time.ParseDuration(fmt.Sprintf("%ds", 5))
			if err != nil {
			}
			time.Sleep(d)
		}
	}
}
