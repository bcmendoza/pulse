package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func UUID() string {
	uuid := uuid.New()
	return fmt.Sprintf("%s", uuid)
}
