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

	script := `
window.addEventListener('message', function(event) {
	if (event.data.type === 'VUE_SAVE_OUTPUT') {
		app.graphToPrompt().then(p => {
			const tempWorkflowJson = JSON.stringify(p.workflow, null, 2);

			window.parent.postMessage({
				type: 'IFRAME_SAVE_OUTPUT',
				value: tempWorkflowJson
			}, event.origin);
		}).catch(error => {
			console.error("Error in graphToPrompt:", error);
		});
	}

	if (event.data.type === 'LOAD_IMG_WORKFLOW') {
		app.loadGraphData(JSON.parse(event.data.workflow), true, true, event.data.name)
	}
});
`

	js = js + script

	r.Body = io.NopCloser(bytes.NewReader([]byte(js)))
	r.ContentLength = int64(len([]byte(js)))
	r.Header.Set("Content-Length", strconv.Itoa(len([]byte(js))))
}
