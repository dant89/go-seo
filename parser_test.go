package goseo_test

import (
	"testing"

	goseo "github.com/dant89/go-seo"
)

func TestParserGetFirstElement(t *testing.T) {
	html := "<html><body><header><h1>Test H1</h1></header></body></html>"
	result, _ := goseo.GetFirstElement(html, "h1")
	if result != "Test H1" {
		t.Errorf("unexpected H1 result")
	}
}
