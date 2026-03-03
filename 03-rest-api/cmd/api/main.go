package main

import (
	//"fmt"
	"log"
	//"net/http"
	//
	//"github.com/gin-gonic/gin"

	"github.com/timotiviert/go-learning/03-rest-api/internal/config"
	"github.com/timotiviert/go-learning/03-rest-api/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	_, err = database.New(cfg.DBConnString)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	//r := gin.Default()
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//if err := r.Run(); err != nil {
	//	log.Fatalf("failed to start server: %v", err)
	//}
}
