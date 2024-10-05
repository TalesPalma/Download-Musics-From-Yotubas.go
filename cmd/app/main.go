package main

import (
	"fmt"

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
	fmt.Println(r.Run(":8080"))
}
