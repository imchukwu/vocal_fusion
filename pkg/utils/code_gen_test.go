package utils

import (
	"regexp"
	"testing"
)

func TestGenerateEventCode(t *testing.T) {
	prefix := "VFMF-SME"
	sequence := 1
	code := GenerateEventCode(prefix, sequence)

	// Format: PREFIX-YYYY-S-XXX
	pattern := `^VFMF-SME-\d{4}-S-001$`
	matched, err := regexp.MatchString(pattern, code)
	if err != nil {
		t.Fatalf("Regex error: %v", err)
	}

	if !matched {
		t.Errorf("Code %s does not match expected pattern %s", code, pattern)
	}
}

func TestGenerateEventCodeSequential(t *testing.T) {
	prefix := "CHC-SCC"
	code1 := GenerateEventCode(prefix, 1)
	code2 := GenerateEventCode(prefix, 2)

	if code1 == code2 {
		t.Errorf("Codes should be unique for different sequences: %s == %s", code1, code2)
	}

	if code1 != "CHC-SCC-2026-S-001" {
		t.Errorf("Unexpected code1: %s", code1)
	}
	if code2 != "CHC-SCC-2026-S-002" {
		t.Errorf("Unexpected code2: %s", code2)
	}
}
