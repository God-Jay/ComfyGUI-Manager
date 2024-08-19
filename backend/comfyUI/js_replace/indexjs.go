package js_replace

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func FindIndexJs(r *http.Response, body []byte) (bool, string) {
	// Parse the HTML
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	// Traverse the HTML to find the script tag with src attribute
	var scriptSrc string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			for _, attr := range n.Attr {
				if attr.Key == "src" && strings.Contains(attr.Val, "index-") {
					scriptSrc = attr.Val
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	r.Body = io.NopCloser(bytes.NewReader(body))

	if scriptSrc != "" {
		log.Println("found script src:", scriptSrc)
		return true, scriptSrc
	} else {
		log.Println("script src not found")
		return false, scriptSrc
	}
}

func ChangeIndexJs(body []byte) []byte {
	log.Println("changeIndexJs")

	modifiedBody := ChangeAppJs(body, true)
	modifiedBody = ChangeUiJs(modifiedBody, true)

	return modifiedBody
}
