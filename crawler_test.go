package goseo_test

import (
	"testing"

	goseo "github.com/dant89/go-seo"
)

func TestFetcher(t *testing.T) {
	webpage := goseo.Webpage{"https://bbc.co.uk"}

	body, _ := webpage.Crawl()
	if body == "" {
		t.Errorf("unexpected empty body")
	}
}
