package main

import (
	"github.com/TalesPalma/internal/ApiServices/configs"
	"github.com/TalesPalma/internal/ApiServices/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	runApi()
}

func runApi() {
	r := gin.Default()
	r.Use(configs.ConfigCors())
	handler.Handler(r)
	r.Run(":8080")
}
