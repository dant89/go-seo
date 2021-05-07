package goseo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Crawler struct{}

type CrawlerWork interface {
	Crawl() (body string, err error)
}

func (c Crawler) Crawl(url string) (string, error) {

	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("not found: %s", url)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response body: %s", url)
	}

	return string(body), nil
}
