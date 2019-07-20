package model

import (
	"sync"
)

// 1. Random patient metrics bubbling up
// 2. REST API for getting metrics
// 3. REST APi for measurement fields
// 4. REST API for setting metrics via fields
// 5. Multiple outcomes per department
// 6. Multiple outcomes per patient

type Hospital struct {
	sync.Mutex
	Stream      `json:"stream"`
	Departments map[string]Department `json:"departments"`
}

func New() *Hospital {
	return &Hospital{
		Stream:      make(Stream, 0),
		Departments: make(map[string]Department),
	}
}
