package mywebsocket

import (
	"log"
	"net/http"

	//"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 抽象出需要的数据结构
// ws连接器  数据  管道 // 已经不是http连接，管道传输数据的事情
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024, // 读的缓冲大小
	WriteBufferSize: 1024, // 写的缓冲大小
	CheckOrigin: func(r *http.Request) bool { // 跨域访问
		return true
	},
}

func WsInit(c *gin.Context) *websocket.Conn {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return nil
	}

	defer ws.Close()

	return ws
}
