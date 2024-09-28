package main

import (
	"time"

	api "github.com/TalesPalma/internal/ApiServices"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	runApi()
}

func runApi() {
	r := gin.Default()
	r.Use(configCors())
	api.Handler(r)
	r.Run(":8080")
}

func configCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
