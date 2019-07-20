package model

import (
	"errors"
	"fmt"
)

// Attrs are inconsequential patient data
// They are added to the Patient struct in unstructured JSON
type Attrs map[string]string

type Patient struct {
	ID     string `json:"id"`
	Stream `json:"stream"`

	Metrics map[string]float64 `json:"metrics"`
	Attrs
}

func (d *Department) AddPatient(p Patient) error {
	if _, ok := d.Patients[p.ID]; !ok {
		d.Patients[p.ID] = p
		return nil
	}
	return errors.New(fmt.Sprintf("Patient %s already exists", p.ID))
}
