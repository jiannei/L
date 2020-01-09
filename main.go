package main

import (
	"L/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	routes.Setup(engine) // 路由加载

	engine.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
