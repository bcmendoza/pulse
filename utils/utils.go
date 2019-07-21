package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func UUID() string {
	uuid := uuid.New()
	return fmt.Sprintf("%s", uuid)
}

func Timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func CalcThresholds(lower, upper float64) []float64 {
	thirdAmt := (upper - lower) / 3
	thresh1 := lower + thirdAmt
	thresh2 := thresh1 + thirdAmt
	return []float64{thresh1, thresh2}
}
