package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"wiza.core/config"
	"wiza.core/db"
	"wiza.core/handler"
	"wiza.core/repository"
	"wiza.core/service"
)

func main() {
	cfg := config.Load()

	pool, err := db.NewPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer pool.Close()

	repo := repository.NewClientRepository(pool)
	svc := service.NewClientService(repo)
	h := handler.NewClientHandler(svc)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/clients/:iin", h.GetByIIN)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Printf("starting on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server: %v", err)
	}
}
