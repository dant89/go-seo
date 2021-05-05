package goseo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Webpage struct {
	Url string
}

type Crawler interface {
	Crawl() (body string, err error)
}

func (w Webpage) Crawl() (string, error) {

	res, err := http.Get(w.Url)
	if err != nil {
		return "", fmt.Errorf("not found: %s", w.Url)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body), nil
}
