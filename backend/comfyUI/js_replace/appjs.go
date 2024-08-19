package js_replace

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
)

var appJsAppendScript = `
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

var appJsReplaceSaveImageScript = `
						{
							content: "Save Image",
							callback: () => {
    							window.parent.postMessage({type: 'saveImg', url: img.src}, '*');
							},
						}
`

func ChangeAppJs(r *http.Response, body []byte) {
	found, toReplaceBlock := findToReplaceBlock(body, "...getCopyImageOption(img)")

	if found {
		body = bytes.ReplaceAll(body, toReplaceBlock, []byte(appJsReplaceSaveImageScript))
	}

	body = append(body, []byte(appJsAppendScript)...)

	r.Body = io.NopCloser(bytes.NewReader(body))
	r.ContentLength = int64(len(body))
	r.Header.Set("Content-Length", strconv.Itoa(len(body)))
}
