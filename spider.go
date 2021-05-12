package goseo

import (
	"net/url"
	"strings"
	"sync"
	"unicode"
)

type Url struct {
	Url      string
	Links    []string
	Crawled  bool
	Depth    int
	Internal bool
}

type UrlList struct {
	urls map[string]*Url
	mux  sync.Mutex
}

type Spider struct {
	Recursive bool
}

type SpiderWeb interface {
	GetLinks()
}

var urlsFound UrlList = UrlList{urls: make(map[string]*Url)}
var wg sync.WaitGroup

func (s Spider) GetLinks(url string, depth int) map[string]*Url {
	maxWorkers := make(chan struct{}, 1000)
	addFoundUrl(url, false, []string{}, 1, url)

	for i := 1; i <= depth; i++ {
		for _, urlFound := range urlsFound.urls {
			if !urlFound.Crawled && urlFound.Depth == i && urlFound.Internal {
				maxWorkers <- struct{}{}
				wg.Add(1)
				go crawlPage(maxWorkers, url, urlFound, i)
			}
		}
		wg.Wait()
	}

	return urlsFound.urls
}

func crawlPage(maxWorkers chan struct{}, originalUrl string, url *Url, depth int) {
	defer wg.Done()

	formattedUrls := []string{}
	crawler := Crawler{}
	response, err := crawler.Crawl(url.Url)
	if err != nil {
		addFoundUrl(url.Url, true, formattedUrls, depth, originalUrl)
		<-maxWorkers
	}

	parser := Parser{}
	results := parser.GetAllLinkHrefs(response)
	for _, result := range results {
		formattedUrl, err := formatUrl(result, originalUrl)
		if err == nil {
			formattedUrls = append(formattedUrls, formattedUrl)
		}
	}

	addFoundUrl(url.Url, true, formattedUrls, depth, originalUrl)
	<-maxWorkers
}

func addFoundUrl(url string, crawled bool, links []string, depth int, originalUrl string) {
	urlsFound.mux.Lock()
	defer urlsFound.mux.Unlock()

	internalUrl := strings.Contains(url, originalUrl)
	newUrl := Url{Url: url, Crawled: crawled, Depth: depth, Links: links, Internal: internalUrl}
	urlsFound.urls[url] = &newUrl

	for _, link := range links {
		if _, ok := urlsFound.urls[link]; !ok {
			internalUrl := strings.Contains(link, originalUrl)
			newChildUrl := Url{Url: link, Crawled: false, Depth: depth + 1, Internal: internalUrl}
			urlsFound.urls[link] = &newChildUrl
		}
	}
}

func formatUrl(rawUrl string, originalUrl string) (string, error) {
	if rawUrl == "/" {
		fixed, _ := url.Parse(originalUrl + "/")
		return fixed.String(), nil
	}

	formatttedUrl := stripTrailingSlash(stripHash(rawUrl))

	var parsedUrl *url.URL
	var err error
	if strings.HasPrefix(formatttedUrl, "http://") || strings.HasPrefix(formatttedUrl, "https://") {
		parsedUrl, err = url.Parse(formatttedUrl)
	} else if strings.HasPrefix(formatttedUrl, "//") {
		parsedUrl, err = url.Parse(originalUrl + strings.TrimPrefix(formatttedUrl, "/"))
	} else if strings.HasPrefix(formatttedUrl, "/") {
		parsedUrl, err = url.Parse(originalUrl + formatttedUrl)
	} else {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return parsedUrl.String(), nil
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
