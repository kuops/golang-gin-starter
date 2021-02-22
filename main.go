package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("hello golang-gin-starter!")

	router := gin.Default()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	stopCh := make(chan os.Signal,1)
	signal.Notify(stopCh,syscall.SIGQUIT,syscall.SIGTERM)
	<-stopCh
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Panic("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}


