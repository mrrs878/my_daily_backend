package main

import (
	"context"
	"demo_1/src/config"
	"demo_1/src/controller"
	"demo_1/src/database"
	"demo_1/src/service/msg"
	"demo_1/src/util/ws"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	gin.SetMode(config.AppMode)
	engine := gin.New()

	database.SetUpDatabase()
	controller.SetupRouter(engine)

	server := &http.Server{
		Addr:         config.AppPort,
		Handler:      engine,
		ReadTimeout:  config.AppReadTimeout * time.Second,
		WriteTimeout: config.AppWriteTimeout * time.Second,
	}

	log.Println("|  Go Http Server Start Successful  |")
	log.Println("|    Port" + config.AppPort + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")

	go ws.WebSocketManager.Start()
	go ws.WebSocketManager.SendService()
	go ws.WebSocketManager.SendService()
	go ws.WebSocketManager.SendGroupService()
	go ws.WebSocketManager.SendGroupService()
	go ws.WebSocketManager.SendAllService()
	go ws.WebSocketManager.SendAllService()
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server listen: %s\n", err)
		}
	}()
	go msg.Setup()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("get signal:", sig)
	log.Println("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("server shutdown: ", err)
	}
	log.Println("server exiting")
}
