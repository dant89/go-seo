package goseo

import (
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/grokify/html-strip-tags-go"
	"golang.org/x/net/html"
)

func GetH1(rawHtml string) (string, error) {
	var h1 *html.Node
	doc, _ := html.Parse(strings.NewReader(rawHtml))

	// recursively parse the html until we find a H1
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "h1" {
			h1 = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if h1 != nil {
		formatted := renderNode(h1)
		stripTags := stripTags(formatted)
		return stripTags, nil
	}
	return "", errors.New("missing <h1> in the node tree")
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func stripTags(h1 string) string {
	stripped := strip.StripTags(h1)
	trimmed := strings.TrimSpace(stripped)
	return trimmed
}
