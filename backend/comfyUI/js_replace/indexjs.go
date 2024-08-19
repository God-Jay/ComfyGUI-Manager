package js_replace

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func FindIndexJs(r *http.Response, body []byte) (bool, string) {
	log.Println("index html", string(body))

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

func ChangeIndexJs(r *http.Response, body []byte) {
	log.Println("changeIndexJs")

	found, toReplaceBlock := findToReplaceBlock(body, "...getCopyImageOption(img)")
	if found {
		body = bytes.ReplaceAll(body, toReplaceBlock, []byte(appJsReplaceSaveImageScript))
	}

	modifiedBody := append(body, []byte(appJsAppendScript)...)
	modifiedBody = append(modifiedBody, []byte(uiJsAppendScript)...)

	r.Body = io.NopCloser(bytes.NewReader(modifiedBody))
	r.ContentLength = int64(len(modifiedBody))
	r.Header.Set("Content-Length", strconv.Itoa(len(modifiedBody)))
}
