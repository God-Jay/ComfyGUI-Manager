package js_replace

import (
	"log"
	"strings"
)

var uiJsAppendScript = `
window.prompt = (message, defaultValue) => {
    app.graphToPrompt().then(p => {
        const json = JSON.stringify(p.workflow, null, 2);
        console.log(json);
    	window.parent.postMessage({type: 'saveFile', json}, '*');
    });
    return ""
};
`

func ChangeUiJs(body []byte, isIndex bool) []byte {
	log.Println("changeUiJs")

	if isIndex {
		uiJsAppendScript = strings.ReplaceAll(uiJsAppendScript, "app.", "window.comfyAPI.app.app.")
	}
	return append(body, []byte(uiJsAppendScript)...)
}
