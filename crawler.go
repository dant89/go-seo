package goseo

import (
	"errors"
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
		return "", err
	}
	if res.StatusCode != 200 {
		err := errors.New("none 200 response")
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	res.Body.Close()

	return string(body), nil
}
