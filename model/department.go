package model

import (
	"errors"
	"fmt"
)

type Department struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Stream `json:"stream"`

	Patients map[string]Patient `json:"patients"`
	Measures map[string]Measure `json:"measures"`
}

func (h *Hospital) AddDepartment(d Department) error {
	h.Lock()
	defer h.Unlock()

	if _, ok := h.Departments[d.ID]; !ok {
		h.Departments[d.ID] = d
		return nil
	}
	return errors.New(fmt.Sprintf("Department %s already exists", d.ID))
}
