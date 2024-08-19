package comfyUI

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"slices"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"comfygui-manager/backend/comfyUI/js_replace"
)

var (
	proxy      *httputil.ReverseProxy
	wsUpgrader websocket.Upgrader
)

func RunProxy(ctx context.Context) {
	proxy = newProxy()
	wsUpgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有源，您可能需要根据实际情况进行调整
		},
	}
	// 启动代理服务器
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			proxyHandler(ctx, w, r)
		})
		http.ListenAndServe(":8189", nil)
	}()
}

func newProxy() *httputil.ReverseProxy {
	targetURL, _ := url.Parse("http://127.0.0.1:8188")

	proxy = httputil.NewSingleHostReverseProxy(targetURL)

	var findJss []string
	hasIndexJs := false
	indexJs := ""

	proxy.ModifyResponse = func(r *http.Response) error {
		if r.Request.URL.Path == "/" || r.Request.URL.Path == "" {
			log.Println("find index js", r.Request.URL.Path)
			if r.StatusCode != 200 {
				log.Println("error", r.StatusCode)
				return nil
			}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				return err
			}
			r.Body.Close()

			hasIndexJs, indexJs = js_replace.FindIndexJs(r, body)
		}

		if hasIndexJs {
			findJss = []string{indexJs}
		} else {
			findJss = []string{"/scripts/ui.js", "/scripts/app.js"}
		}

		if slices.Contains(findJss, r.Request.URL.Path) {
			log.Println("modify", r.Request.URL.Path)
			if r.StatusCode != 200 {
				log.Println("error", r.StatusCode)
				return nil
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				return err
			}
			r.Body.Close()

			switch r.Request.URL.Path {
			case "/scripts/ui.js":
				js_replace.ChangeUiJs(r, body)
			case "/scripts/app.js":
				js_replace.ChangeAppJs(r, body)
			case indexJs:
				js_replace.ChangeIndexJs(r, body)
			default:
				return nil
			}

			return nil
		}

		return nil
	}
	return proxy
}

func proxyHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if websocket.IsWebSocketUpgrade(r) {
		handleWebSocket(ctx, w, r)
		return
	}
	proxy.ServeHTTP(w, r)
}

func handleWebSocket(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	targetURL := "ws://127.0.0.1:8188" + r.URL.Path
	if r.URL.RawQuery != "" {
		targetURL += "?" + r.URL.RawQuery
	}

	// 连接到目标WebSocket服务器
	targetConn, _, err := websocket.DefaultDialer.Dial(targetURL, nil)
	if err != nil {
		http.Error(w, "Could not connect to WebSocket server", http.StatusInternalServerError)
		return
	}
	defer targetConn.Close()

	// 升级客户端连接为WebSocket
	clientConn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade connection", http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	// 双向转发WebSocket消息
	errc := make(chan error, 2)
	go proxyWebSocket(ctx, targetConn, clientConn, errc)
	go proxyWebSocket(ctx, clientConn, targetConn, errc)

	<-errc
}

func proxyWebSocket(ctx context.Context, dst, src *websocket.Conn, errc chan error) {
	for {
		messageType, message, err := src.ReadMessage()
		if err != nil {
			errc <- err
			return
		}

		var wsMsg WSMessage
		json.Unmarshal(message, &wsMsg)
		if wsMsg.Type == "executed" {
			log.Println("executed ....", string(message))
			if outputFile := gjson.GetBytes(message, "data.output.images.0.filename"); outputFile.Exists() {
				log.Println("send js events")
				runtime.EventsEmit(ctx, "OutputImageFile", outputFile.String())
			}
		}

		err = dst.WriteMessage(messageType, message)
		if err != nil {
			errc <- err
			return
		}
	}
}

type WSMessage struct {
	Type     string `json:"type"`
	Data     any    `json:"data,omitempty"`
	PromptId string `json:"prompt_id"`
}
