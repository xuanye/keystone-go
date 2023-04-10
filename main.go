package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/xuanye/keystone-go/adapter"
	"github.com/xuanye/keystone-go/application"
	"github.com/xuanye/keystone-go/common"
	"github.com/xuanye/keystone-go/common/config"
	"github.com/xuanye/keystone-go/domain"
	"github.com/xuanye/keystone-go/infrastructure"
)

func setupEnvironment() *common.HostEnvironment {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = os.Getenv("APP_ENV")
	}

	return &common.HostEnvironment{EnvironmentName: env}
}

func setupHttpServer(router *gin.Engine, settings *config.Settings) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Port),
		Handler: router,
	}

	return server
}

func setDependency() *dig.Container {

	container := dig.New()

	container.Provide(setupEnvironment)
	container.Provide(setupHttpServer)
	adapter.InitContainer(container)
	application.InitContainer(container)
	domain.InitContainer(container)
	infrastructure.InitContainer(container)

	return container
}

func main() {

	c := setDependency()

	var server *http.Server

	// 运行容器
	if err := c.Invoke(func(s *http.Server) {
		server = s
	}); err != nil {
		panic(err)
	}

	// 启动 HTTP Server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 监听关闭信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")

	// 创建一个 context.Context 对象，设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭 HTTP Server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
