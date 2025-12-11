package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGenerateMessage(t *testing.T) {
	expected := "--- Cloud Build Sandbox Demo ---"
	result := generateMessage()

	if result != expected {
		t.Errorf("generateMessage() = %q, want %q", result, expected)
	}
}

func TestGetHostname(t *testing.T) {
	result := getHostname()

	if result == "" {
		t.Error("getHostname() returned empty string")
	}

	// Hostname should either be a valid string or "unknown"
	if len(result) == 0 {
		t.Error("getHostname() should return a non-empty string")
	}
}

func TestGetHostnameWithGetter_Success(t *testing.T) {
	mockGetter := func() (string, error) {
		return "test-hostname", nil
	}

	result := getHostnameWithGetter(mockGetter)
	expected := "test-hostname"

	if result != expected {
		t.Errorf("getHostnameWithGetter() = %q, want %q", result, expected)
	}
}

func TestGetHostnameWithGetter_Error(t *testing.T) {
	mockGetter := func() (string, error) {
		return "", os.ErrInvalid
	}

	result := getHostnameWithGetter(mockGetter)
	expected := "unknown"

	if result != expected {
		t.Errorf("getHostnameWithGetter() on error = %q, want %q", result, expected)
	}
}

func TestFormatTimestamp(t *testing.T) {
	// Test with a known time
	testTime := time.Date(2025, 12, 11, 10, 30, 0, 0, time.UTC)
	expected := "2025-12-11T10:30:00Z"
	result := formatTimestamp(testTime)

	if result != expected {
		t.Errorf("formatTimestamp() = %q, want %q", result, expected)
	}
}

func TestFormatTimestampWithDifferentTimezone(t *testing.T) {
	// Test with a different timezone
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Skip("Could not load America/New_York timezone")
	}

	testTime := time.Date(2025, 1, 1, 12, 0, 0, 0, loc)
	result := formatTimestamp(testTime)

	// Should be in RFC3339 format
	_, err = time.Parse(time.RFC3339, result)
	if err != nil {
		t.Errorf("formatTimestamp() did not return valid RFC3339 format: %q", result)
	}
}

func BenchmarkGenerateMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateMessage()
	}
}

func BenchmarkFormatTimestamp(b *testing.B) {
	testTime := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatTimestamp(testTime)
	}
}

func TestMain(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run main
	main()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify output contains expected strings
	expectedStrings := []string{
		"--- Cloud Build Sandbox Demo ---",
		"Current time:",
		"Hostname:",
		"Build test successful!",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("main() output missing expected string %q\nGot output:\n%s", expected, output)
		}
	}

	// Verify timestamp format (should contain RFC3339-like format)
	if !strings.Contains(output, "T") || !strings.Contains(output, ":") {
		t.Error("main() output does not contain a properly formatted timestamp")
	}
}
