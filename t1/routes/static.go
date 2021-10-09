package routes

import (
	"github.com/gin-gonic/gin"
)

func loadStaticFile(r *gin.Engine) {
	r.StaticFile("/HYShangWeiShouShuW.ttf", "./static/HYShangWeiShouShuW.ttf")
}
func loadStaticFiles(r *gin.Engine) {
	r.Static("/static", "./static")
}
