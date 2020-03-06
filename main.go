package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/zsais/go-gin-prometheus"
	"net/http"
	"os"
	"os/signal"
	"goexample/handler"
	"time"
)

func main() {
	var (
		httpAddr  = flag.String("http.addr", ":8080", "HTTP listen address")
		debugAddr = flag.String("debug.addr", ":23333", "metrics listen address")
	)
	flag.Parse()
	svc := handler.NewService()
	gin.DisableConsoleColor()
	router := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.SetListenAddress(*debugAddr)
	p.Use(router)

	router.GET("/ping", svc.Ping)
	router.POST("/count", svc.Count)
	server := &http.Server{
		Addr:    *httpAddr,
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				glog.Warning("Server closed under request")
			} else {
				glog.Fatal("Server closed unexpect")
			}
		}
	}()

	<-quit
	glog.Info("receive interrupt signal")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		glog.Fatal("Server Shutdown:", err)
	}

	glog.Info("Server exiting")
}
