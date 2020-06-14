package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"task-engine/api/routers"
	"task-engine/core"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger())

	app.Use(gin.Recovery())
	//app.Use(middleware.CustomMiddleware)

	c, err := core.Conf.Map("web")
	if err != nil {
		panic(err)
	}

	gin.SetMode(c["RunMode"].(string))

	router := routers.InitRouter(app)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c["HttPort"]),
		Handler:      router,
		ReadTimeout:  time.Duration(c["ReadTimeout"].(int)) * time.Second,
		WriteTimeout: time.Duration(c["WriteTimeout"].(int)) * time.Second,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
