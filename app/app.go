package app

import (
	"L/app/http/middlewares"
	"L/routes"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Application struct {
	Engine *gin.Engine
}

var server *http.Server

func New() *Application {
	// 实例化 gin
	engine := gin.Default()

	// 全局中间件
	middlewares.Setup(engine)

	// 路由加载
	routes.Setup(engine)

	return &Application{
		Engine: engine,
	}
}

func (app *Application) Run() {
	// 配置服务
	server = &http.Server{
		Addr:         ":8080",
		Handler:      app.Engine,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	// 启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Http server listen: %s\n", err)
		}
	}()

	// 欢迎信息
	app.welcome()

	// 关闭服务
	app.terminate()
}

func (app *Application) welcome() {
	// 控制台输出信息
	log.Println("|-----------------------------------|")
	log.Println("|            Welcome to L!          |")
	log.Println("|-----------------------------------|")
}

func (app *Application) terminate() {
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Signal:", sig)
	log.Println("Shutdown Server ...")

	// 关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
