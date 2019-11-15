package main

import (
	"fmt"
	"go-gin-base/bootstrap"
	"go-gin-base/common"
	"go-gin-base/config"
	"go-gin-base/web/routes"
	"net/http"
	"runtime"
	"time"
)

func initEnv()  {
	// 配置线程数量
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("基础架构", "xiaolin")
	app.Bootstrap()
	app.Configure(routes.ApiConfigure)

	return app
}

func main ()  {
	initEnv();
	app := newApp()


	startServer(app)
}

func startServer (b *bootstrap.Bootstrapper)  {
	server := &http.Server{
		Addr:           ":" + config.Cfg.Produce.Port,
		Handler:        b,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}




	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	// 平滑退出，先结束所有在执行的任务
	common.GracefulExitWeb(server)
}
