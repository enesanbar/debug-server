package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDebugHandlerIncludesEnvVars(t *testing.T) {
	const envKey = "DEBUG_SERVER_TEST_ENV"
	const envValue = "present"

	t.Setenv(envKey, envValue)

	req := httptest.NewRequest(http.MethodGet, "/?foo=bar", nil)
	rec := httptest.NewRecorder()

	debugHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var response DebugInfo
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got := response.EnvVars[envKey]; got != envValue {
		t.Fatalf("expected env var %q to be %q, got %q", envKey, envValue, got)
	}

	if got := response.Method; got != http.MethodGet {
		t.Fatalf("expected method %q, got %q", http.MethodGet, got)
	}

	if got := response.QueryParams["foo"]; got != "bar" {
		t.Fatalf("expected query param foo=bar, got %q", got)
	}

	if _, ok := response.EnvVars["PATH"]; !ok && os.Getenv("PATH") != "" {
		t.Fatal("expected PATH to be included in env vars")
	}
}
