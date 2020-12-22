package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	HREF string
	Text string
}

func removeEmpty(ss []string) []string {
	r := make([]string, 0, len(ss))
	for _, s := range ss {
		if strings.TrimSpace(s) == "" {
			continue
		}
		r = append(r, s)
	}
	return r
}
func getText(n *html.Node) string {
	var data []string
	switch n.Type {
	case html.ElementNode:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			data = append(data, getText(c))
		}
	case html.TextNode:
		data = append(data, strings.TrimSpace(n.Data))
	}
	return strings.Join(removeEmpty(data), " ")
}

func getHREF(attrs []html.Attribute) string {
	for _, a := range attrs {
		if a.Key == "href" {
			return strings.TrimSpace(a.Val)
		}
	}
	return ""
}

func Parse(r io.Reader) ([]*Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []*Link
	var collectLinks func(*html.Node)
	collectLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			l := &Link{
				HREF: getHREF(n.Attr),
				Text: getText(n),
			}
			links = append(links, l)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				collectLinks(c)
			}
		}
	}
	collectLinks(doc)
	return links, nil
}
