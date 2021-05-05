package goseo_test

import (
	"testing"

	goseo "github.com/dant89/go-seo"
)

func TestParser(t *testing.T) {
	html := "<html><body><header><h1>Test H1</h1></header></body></html>"
	result, _ := goseo.GetH1(html)
	if result != "Test H1" {
		t.Errorf("unexpected H1 result")
	}
}
