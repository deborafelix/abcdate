package main

import "github.com/gin-gonic/gin"

func startRouter() {
	router := gin.Default()
	router.GET("/random-letter", getAphabet)
	router.GET("/random-letter/:letter", getPlace)

	router.Run("localhost:8080")
}
