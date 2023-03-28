package ws

import (
	http "github.com/fengyeall111/gnet-http"
	"github.com/panjf2000/gnet/v2"
)

type WebSocketServer struct { // 添加一个环形队列
	http.HttpServer
}

func (wss *WebSocketServer) Run(addr string, opts ...gnet.Option) error {
	return wss.HttpServer.Run(addr, opts...)
}
func (wss *WebSocketServer) OnTextMessage() {
}
func (wss *WebSocketServer) OnBinaryMessage() {

}
func (wss *WebSocketServer) OnPingMessage() {
}
func (wss *WebSocketServer) OnPongMessage() {
}

func (wss *WebSocketServer) WriteTextMessageAsync([]byte) {
}
func (wss *WebSocketServer) WriteBinaryMessageAsync([]byte) {
}
func (wss *WebSocketServer) WritePingMessageAsync([]byte) {
}
func (wss *WebSocketServer) WritePongMessageAsync([]byte) {
}
func (wss *WebSocketServer) WriteCloseMessageAsync([]byte) {
}
