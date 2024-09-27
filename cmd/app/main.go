package main

import (
	api "github.com/TalesPalma/internal/ApiServices"
	"github.com/gin-gonic/gin"
)

func main() {
	runApi()
}

func runApi() {
	r := gin.Default()
	api.Handler(r)
	r.Run(":8080")
}
