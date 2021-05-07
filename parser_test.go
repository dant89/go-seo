package goseo_test

import (
	"testing"

	goseo "github.com/dant89/go-seo"
)

func TestParserGetFirstElement(t *testing.T) {
	parser := goseo.Parser{}

	html := "<html><body><header><h1>Test H1</h1></header></body></html>"
	result, _ := parser.GetFirstElement(html, "h1", false)
	if result != "Test H1" {
		t.Errorf("unexpected H1 result")
	}
}
