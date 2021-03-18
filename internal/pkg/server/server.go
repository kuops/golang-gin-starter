package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Run(ctx context.Context) {
	router := gin.Default()
	registerRouter(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Printf("http: web server shutdown complete")
			} else {
				log.Printf("http: web server closed unexpect: %s", err)
			}
		}
	}()

	<-ctx.Done()
	log.Printf("http: shutting down web server")
	err := server.Close()
	if err != nil {
		log.Printf("http: web server shutdown failed: %v", err)
	}
}
