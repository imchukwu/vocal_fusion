package utils

import (
	"regexp"
	"testing"
)

func TestGenerateEventCode(t *testing.T) {
	code := GenerateEventCode()

	// Format: VF-YYYY-XXXXX
	pattern := `^VF-\d{4}-[A-Z0-0]{5}$`
	matched, err := regexp.MatchString(pattern, code)
	if err != nil {
		t.Fatalf("Regex error: %v", err)
	}

	if !matched {
		t.Errorf("Code %s does not match expected pattern %s", code, pattern)
	}
}

func TestGenerateEventCodeUniqueness(t *testing.T) {
	codes := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		code := GenerateEventCode()
		if codes[code] {
			t.Errorf("Duplicate code generated: %s", code)
		}
		codes[code] = true
	}
}
