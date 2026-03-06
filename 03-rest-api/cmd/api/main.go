package main

import (
	//"fmt"
	"log"
	"net/http"

	//"net/http"
	//
	//"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin"
	"github.com/timotiviert/go-learning/03-rest-api/internal/config"
	"github.com/timotiviert/go-learning/03-rest-api/internal/database"
	"github.com/timotiviert/go-learning/03-rest-api/internal/handlers"
	"github.com/timotiviert/go-learning/03-rest-api/internal/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.New(cfg.DBConnString)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	userRepo := database.NewUserRepository(db)
	userService := services.New(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", userHandler.Register)
	r.GET("/users/:email", userHandler.GetByEmail)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
