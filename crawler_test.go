package goseo_test

import (
	"testing"

	goseo "github.com/dant89/go-seo"
)

func TestFetcher(t *testing.T) {
	crawler := goseo.Crawler{}

	body, _ := crawler.Crawl("https://bbc.co.uk")
	if body == "" {
		t.Errorf("unexpected empty body")
	}
}
