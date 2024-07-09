package js_replace

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ChangeAppJs(r *http.Response, body []byte) {
	toReplaceJs := `
					options.unshift(
						{
							content: "Open Image",
							callback: () => {
								let url = new URL(img.src);
								url.searchParams.delete("preview");
								window.open(url, "_blank");
							},
						},
						...getCopyImageOption(img), 
						{
							content: "Save Image",
							callback: () => {
								const a = document.createElement("a");
								let url = new URL(img.src);
								url.searchParams.delete("preview");
								a.href = url;
								a.setAttribute("download", new URLSearchParams(url.search).get("filename"));
								document.body.append(a);
								a.click();
								requestAnimationFrame(() => a.remove());
							},
						}
					);
`
	newJs := `
					options.unshift(
						{
							content: "Open Image",
							callback: () => {
								let url = new URL(img.src);
								url.searchParams.delete("preview");
								window.open(url, "_blank");
							},
						},
						...getCopyImageOption(img), 
						{
							content: "Save Image",
							callback: () => {
    							window.parent.postMessage({type: 'saveImg', url: img.src}, '*');
							},
						}
					);
`

	js := strings.ReplaceAll(replaceJs(string(body)), replaceJs(toReplaceJs), replaceJs(newJs))

	r.Body = io.NopCloser(bytes.NewReader([]byte(js)))
	r.ContentLength = int64(len([]byte(js)))
	r.Header.Set("Content-Length", strconv.Itoa(len([]byte(js))))
}
