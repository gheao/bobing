package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var liveLists AliveList

func checkOrigin(r *http.Request) bool {
	return true
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func cmdWebSocket(c *gin.Context) {

	wsConn, _ := upGrader.Upgrade(c.Writer, c.Request, nil)

	defer wsConn.Close()

	for {
		_, msg, _ := wsConn.ReadMessage()
		fmt.Printf("%s\n", msg)
	}
}
func ServerInit() *gin.Engine {
	r := gin.Default()

	r.GET("/login", userLogin)
	r.POST("/history", getHistory)
	liveLists.Listen()
	r.GET("/room", cmdWebSocket)
	r.POST("/build", newroom)

	loadStaticFile(r)
	loadStaticFiles(r)

	return r
}
