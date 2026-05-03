package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateEventCode generates a unique code for an event registration.
// Format: PREFIX-YYYY-S-XXX where XXX is a sequence number.
func GenerateEventCode(prefix string, sequence int) string {
	year := time.Now().Year()
	return fmt.Sprintf("%s-%d-S-%03d", prefix, year, sequence)
}
