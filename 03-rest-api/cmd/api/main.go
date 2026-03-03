package main

import (
	"fmt"
	"log"
	//"net/http"
	//
	//"github.com/gin-gonic/gin"

	"github.com/timotiviert/go-learning/03-rest-api/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	fmt.Println(cfg.DbConnString)

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
