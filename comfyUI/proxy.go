package comfyUI

import (
	"comfygui-manager/comfyUI/js_replace"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"slices"
)

var (
	proxy      *httputil.ReverseProxy
	wsUpgrader websocket.Upgrader
)

func RunProxy() {
	proxy = newProxy()
	wsUpgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有源，您可能需要根据实际情况进行调整
		},
	}
	// 启动代理服务器
	go func() {
		http.HandleFunc("/", proxyHandler)
		http.ListenAndServe(":8189", nil)
	}()
}

func newProxy() *httputil.ReverseProxy {
	targetURL, _ := url.Parse("http://127.0.0.1:8188")

	proxy = httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ModifyResponse = func(r *http.Response) error {
		if slices.Contains([]string{"/scripts/ui.js", "/scripts/app.js"}, r.Request.URL.Path) {
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
			default:
				return nil
			}

			return nil
		}
		return nil
	}
	return proxy
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	if websocket.IsWebSocketUpgrade(r) {
		handleWebSocket(w, r)
		return
	}
	proxy.ServeHTTP(w, r)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
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
	go proxyWebSocket(targetConn, clientConn, errc)
	go proxyWebSocket(clientConn, targetConn, errc)

	<-errc
}

func proxyWebSocket(dst, src *websocket.Conn, errc chan error) {
	for {
		messageType, message, err := src.ReadMessage()
		if err != nil {
			errc <- err
			return
		}

		// 在这里，您可以修改WebSocket消息
		// 例如：message = modifyWebSocketMessage(message)

		err = dst.WriteMessage(messageType, message)
		if err != nil {
			errc <- err
			return
		}
	}
}
