package middlewares

import "github.com/gin-gonic/gin"

var (
	middleware = []gin.HandlerFunc{
		gin.Logger(),
		gin.Recovery(),
	}
)

func Setup(engine *gin.Engine) {
	for _, val := range middleware {
		engine.Use(val)
	}
}
