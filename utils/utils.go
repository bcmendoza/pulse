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
