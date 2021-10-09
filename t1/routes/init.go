package routes

import (
	"github.com/gin-gonic/gin"
)

func ServerInit() *gin.Engine {
	r := gin.Default()

	r.GET("/login", userLogin)
	r.POST("/history", getHistory)

	r.GET("/room", cmdWebSocket)
	r.POST("/build", newroom)

	go al.Listen()
	loadStaticFile(r)
	loadStaticFiles(r)

	return r
}
