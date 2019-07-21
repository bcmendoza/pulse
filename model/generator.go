package model

import (
	"fmt"
	"math/rand"
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
}

var deptMetrics = map[string][]TestData{
	"Critical Care: SICU": []TestData{
		TestData{},
		TestData{},
		TestData{},
	},
	"Maternal & Fetal Medicine": []TestData{
		TestData{},
		TestData{},
		TestData{},
	},
	"Level III NICU": []TestData{
		TestData{},
		TestData{},
		TestData{},
	},
	"Radiation Oncology": []TestData{
		TestData{},
		TestData{},
		TestData{},
	},
	"Cardiology Group": []TestData{
		TestData{},
		TestData{},
		TestData{},
	},
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
	t, err := time.ParseDuration("5m")
	if err != nil {
	}
	ticker := time.NewTicker(t)
INFINITE_LOOP:
	for {
		for k := range h.MetricKeys {
			if m, ok := h.Children[k.Department].Children[k.Patient].Children[k.Metric]; ok {
				h.AddMetricPulse(k.Department, k.Patient, k.Metric, utils.Random(m.Stream.Lower, m.Stream.Upper))
			}
		}
		d, err := time.ParseDuration(fmt.Sprintf("%ds", rand.Intn(5)))
		if err != nil {
		}
		time.Sleep(d)
		for range ticker.C {
			break INFINITE_LOOP
		}
	}
}
