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
// Format: VF-YYYY-XXXXX where XXXXX is a random string.
func GenerateEventCode() string {
	year := time.Now().Year()
	randomPart := randomString(5)
	return fmt.Sprintf("VF-%d-%s", year, randomPart)
}

func randomString(n int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
