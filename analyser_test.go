package goseo_test

import (
	"testing"

	goseo "github.com/dant89/go-seo"
)

func TestAnalyser(t *testing.T) {
	h1Short := "This is a test"
	h1ShortResult := goseo.CheckH1Length(h1Short)
	if h1ShortResult != "The H1 is too short, aim for 20 characters minimum." {
		t.Errorf("Incorrect message for short H1")
	}

	h1Long := "This is a test, This is a test, This is a test, This is a test, This is a test, This is a test"
	h1LongResult := goseo.CheckH1Length(h1Long)
	if h1LongResult != "The H1 is too long, aim for 70 characters maximum." {
		t.Errorf("Incorrect message for long H1")
	}

	h1Perfect := "This is a test, This is a test, This is a test"
	h1PerfectResult := goseo.CheckH1Length(h1Perfect)
	if h1PerfectResult != "The H1 is perfect." {
		t.Errorf("Incorrect message for prefect H1.")
	}
}
