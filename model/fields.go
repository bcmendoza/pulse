package model

import (
	"errors"
	"fmt"
)

// Measure is used to track metrics
type Measure struct {
	Label    string          `json:"label"`
	UnitType string          `json:"unitType"`
	Rules    map[string]Rule `json:"rules"`
}

// Rule associates a range with label and color to eval a Pulse score
// Based on Rules defined, each Pulse should be assigned a Rule
type Rule struct {
	Label string `json:"label"`
	Color string `json:"color"`
	LTE   int    `json:"lte"`
	GTE   int    `json:"gte"`
}

func (d *Department) AddMeasure(m Measure) error {
	if _, ok := d.Measures[m.Label]; !ok {
		d.Measures[m.Label] = m
		return nil
	}
	return errors.New(fmt.Sprintf("Measure %s already exists", m.Label))
}
