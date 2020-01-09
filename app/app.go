package app

import (
	"L/app/http/middlewares"
	"L/app/providers"
	"L/routes"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Application struct {
	Engine *gin.Engine
	Config *providers.Config
	Logger *zap.Logger
}

var server *http.Server

func New() *Application {
	// 加载配置
	cfg := providers.NewConfig()

	// 应用模式
	gin.SetMode(cfg.App.Mode)

	// 实例化 gin
	engine := gin.Default()

	// 全局中间件
	middlewares.Setup(engine)

	// 路由加载
	routes.Setup(engine)

	// 收集日志
	logger := providers.NewLogger()

	return &Application{
		Engine: engine,
		Config: cfg,
		Logger: logger,
	}
}

func (app *Application) Run() {
	// 配置服务
	server = &http.Server{
		Addr:         ":8080",
		Handler:      app.Engine,
		ReadTimeout:  app.Config.Http.ReadTimeout * time.Second,
		WriteTimeout: app.Config.Http.WriteTimeout * time.Second,
	}

	// 启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Http server listen: %s\n", err)
		}
	}()

	// 记录日志
	app.Logger.Info("L is running here")

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
	ctx, cancel := context.WithTimeout(context.Background(), app.Config.App.CancelTimeout*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
