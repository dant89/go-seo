package goseo_test

import (
	"strings"
	"testing"

	goseo "github.com/dant89/go-seo"
)

func TestSpider(t *testing.T) {

	spider := goseo.Spider{}
	links := spider.GetLinks("https://dant.blog", true, 1)
	if len(links) < 1 {
		t.Errorf("no links found")
	}

	found := false
	for _, link := range links {
		if strings.Contains(link.Url, "author/dant") {
			found = true
		}
	}

	if !found {
		t.Errorf("author link not found")
	}
}
