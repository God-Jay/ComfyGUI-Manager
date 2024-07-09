package js_replace

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
)

func ChangeUiJs(r *http.Response, body []byte) {
	log.Println("changeUiJs")
	script := `
window.prompt = (message, defaultValue) => {
    app.graphToPrompt().then(p => {
        const json = JSON.stringify(p.workflow, null, 2);
        console.log(json);
    	window.parent.postMessage({type: 'saveFile', json}, '*');
    });
    return ""
};
`
	modifiedBody := append(body, []byte(script)...)

	r.Body = io.NopCloser(bytes.NewReader(modifiedBody))
	r.ContentLength = int64(len(modifiedBody))
	r.Header.Set("Content-Length", strconv.Itoa(len(modifiedBody)))
}
