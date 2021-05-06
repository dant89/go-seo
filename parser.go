package goseo

import (
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/grokify/html-strip-tags-go"
	"golang.org/x/net/html"
)

func GetFirstElement(rawHtml string, element string) (string, error) {
	var foundElement *html.Node
	doc, _ := html.Parse(strings.NewReader(rawHtml))

	// recursively parse the html until we find a H1
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == element {
			foundElement = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if foundElement != nil {
		formatted := renderNode(foundElement)
		stripped := stripHtml(formatted)
		return stripped, nil
	}

	return "", errors.New("missing <" + element + "> in the node tree")
}

func GetAllElements(rawHtml string, element string) []string {
	var foundElements []*html.Node
	var formattedElements []string
	doc, _ := html.Parse(strings.NewReader(rawHtml))

	// recursively parse the html until we find a H1
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == element {
			foundElements = append(foundElements, node)
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)

	if len(foundElements) > 0 {
		for _, v := range foundElements {
			formatted := renderNode(v)
			formattedElements = append(formattedElements, stripHtml(formatted))
		}
	}

	return formattedElements
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func stripHtml(h1 string) string {
	stripped := strip.StripTags(h1)
	trimmed := strings.TrimSpace(stripped)
	return trimmed
}
