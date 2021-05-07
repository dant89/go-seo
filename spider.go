package goseo

import (
	"strings"
	"unicode"
)

type Spider struct {
	Recursive bool
}

type SpiderWeb interface {
	GetLinks()
}

func (s Spider) GetLinks(url string, internalOnly bool) []string {

	crawler := Crawler{}
	response, err := crawler.Crawl(url)
	if err != nil {
		panic(err)
	}

	parser := Parser{}
	results := parser.GetAllLinkHrefs(response)

	var formatted []string
	for _, result := range results {
		strippedHash := stripHash(result)
		strippedSlash := stripTrailingSlash(strippedHash)
		if internalOnly {
			if strings.Index(strippedSlash, "/") == 0 || strings.Contains(strippedSlash, url) {
				formatted = append(formatted, strippedSlash)
			}
		} else {
			formatted = append(formatted, strippedSlash)
		}
	}

	results = uniqueLinks(formatted)
	return results
}

func uniqueLinks(links []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range links {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func stripHash(source string) string {
	if cut := strings.IndexAny(source, "#"); cut >= 0 {
		return strings.TrimRightFunc(source[:cut], unicode.IsSpace)
	}
	return source
}

func stripTrailingSlash(source string) string {
	return strings.TrimSuffix(source, "/")
}
