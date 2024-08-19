package js_replace

import (
	"log"
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

func ChangeUiJs(body []byte) []byte {
	log.Println("changeUiJs")

	return append(body, []byte(uiJsAppendScript)...)
}
