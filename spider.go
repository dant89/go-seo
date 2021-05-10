package goseo

import (
	"net/url"
	"strings"
	"sync"
	"unicode"
)

type Url struct {
	Url     string
	Links   []string
	Crawled bool
	Depth   int
}

type Spider struct {
	Recursive bool
}

type SpiderWeb interface {
	GetLinks()
}

var urlsFound map[string]*Url = make(map[string]*Url)
var maxWorkers chan struct{} = make(chan struct{}, 10)
var lock sync.Mutex
var wg sync.WaitGroup

func (s Spider) GetLinks(url string, internalOnly bool, depth int) map[string]*Url {
	addFoundUrl(url, false, []string{}, 1)

	for i := 1; i <= depth; i++ {
		for _, value := range urlsFound {
			maxWorkers <- struct{}{}
			wg.Add(1)
			go crawlPage(&wg, maxWorkers, url, value, internalOnly, i)
		}
		wg.Wait()
	}

	return urlsFound
}

func crawlPage(wg *sync.WaitGroup, maxWorkers chan struct{}, originalUrl string, url *Url, internalOnly bool, depth int) {
	defer wg.Done()

	if !url.Crawled && url.Depth == depth {
		crawler := Crawler{}
		response, err := crawler.Crawl(url.Url)
		if err != nil {
			addFoundUrl(url.Url, true, []string{}, depth)
			<-maxWorkers
			return
		}

		var formattedUrls []string
		parser := Parser{}
		results := parser.GetAllLinkHrefs(response)
		for _, result := range results {
			strippedHash, err := formatUrl(result, originalUrl)
			if err != nil {
				addFoundUrl(url.Url, true, []string{}, depth)
				<-maxWorkers
				return
			}
			if internalOnly {
				if strings.Contains(strippedHash, originalUrl) {
					formattedUrls = append(formattedUrls, strippedHash)
				}
			} else {
				formattedUrls = append(formattedUrls, strippedHash)
			}
		}

		addFoundUrl(url.Url, true, formattedUrls, depth)
	}
	<-maxWorkers
}

func addFoundUrl(url string, crawled bool, links []string, depth int) {
	lock.Lock()
	defer lock.Unlock()

	urlInstace := Url{Url: url, Crawled: crawled, Depth: depth, Links: links}
	urlsFound[url] = &urlInstace

	if len(links) > 0 {
		for _, link := range links {
			urlInstace := Url{Url: link, Crawled: false, Depth: depth + 1}
			urlsFound[link] = &urlInstace
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
