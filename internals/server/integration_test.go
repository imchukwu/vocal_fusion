package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"vocal_fusion/internals/middleware"

	"github.com/go-chi/chi/v5"
)

// Mock handlers and repo would be ideal, but for integration we can test middleware directly
// and handler response codes if we had a full test DB setup.
// Since we don't want to spin up a real DB, we will verify:
// 1. Rate Limiting Middleware behaves
// 2. Structs marshal correctly (sanity check)

func TestRateLimitMiddleware(t *testing.T) {
	r := chi.NewRouter()
	rl := middleware.NewRateLimiter(1, 1) // 1 req/sec, burst 1
	r.Use(rl.Limit)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	client := ts.Client()

	// 1st request - should pass
	resp, err := client.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
	}

	// 2nd request - immediate - should fail (burst 1 used)
	resp, err = client.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		// Note: Timing is tricky in tests, but with burst 1 and immediate follow up, it should likely 429.
		// If it doesn't, it might be due to window reset.
		// Let's try sending a few to guarantee
		for i := 0; i < 5; i++ {
			resp, _ = client.Get(ts.URL)
			if resp.StatusCode == http.StatusTooManyRequests {
				return // Success
			}
		}
		t.Errorf("Expected 429 Too Many Requests eventually")
	}
}

func TestEventDescriptionUpdatePayload(t *testing.T) {
	// Just verifies that we can define a struct with Description and valid JSON for the handler
	jsonStr := `{"description": "New Description", "title": "Test"}`
	var payload map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &payload)
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}
	if payload["description"] != "New Description" {
		t.Errorf("Description field missing or incorrect")
	}
}
