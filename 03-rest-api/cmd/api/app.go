package main

import (
	"log"
	//"net/http"

	"github.com/gin-gonic/gin"

	"github.com/timotiviert/go-learning/03-rest-api/internal/config"
)

type application struct {
	Config config.Config
	Logger *log.Logger
}

func (app *application) routes() *gin.Engine {
	r := gin.Default()
	return r
}
