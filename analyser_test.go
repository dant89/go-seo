package goseo_test

import (
	"testing"

	goseo "github.com/dant89/go-seo"
)

func TestAnalyser(t *testing.T) {
	h1Short := "This is a test"
	h1ShortResult := goseo.CheckH1Length(h1Short)
	if !h1ShortResult.Passed() {
		for _, err := range h1ShortResult.GetFeedback() {
			if err != goseo.H1ShortError {
				t.Errorf("Incorrect message for short H1")
			}
		}
	}

	h1Long := "This is a test, This is a test, This is a test, This is a test, This is a test, This is a test"
	h1LongResult := goseo.CheckH1Length(h1Long)
	if !h1LongResult.Passed() {
		for _, err := range h1LongResult.GetFeedback() {
			if err != goseo.H1LongError {
				t.Errorf("Incorrect message for long H1")
			}
		}
	}

	h1Perfect := "This is a test, This is a test, This is a test"
	h1PerfectResult := goseo.CheckH1Length(h1Perfect)
	if !h1PerfectResult.Passed() {
		if len(h1PerfectResult.GetFeedback()) > 0 {
			t.Errorf("Incorrect message for perfect H1.")
		}
	}
}
