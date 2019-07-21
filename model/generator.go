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
		for i := 0; i < 2; i++ {
			h.AddHospitalMetrics(utils.UUID(), "x", 1, 1000)
		}
	}
	go h.RunGenerator()
}

func (h *Hospital) RunGenerator() {
	for {
		for k := range h.MetricKeys {
			h.AddMetricPulse(
				k.Department,
				k.Patient,
				k.Metric,
				float64(rand.Intn(999)),
			)
			d, err := time.ParseDuration(fmt.Sprintf("%ds", rand.Intn(5)))
			if err != nil {
			}
			time.Sleep(d)
		}
	}
}
