package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getAphabet(c *gin.Context) {

	b := make([]rune, 1)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	c.IndentedJSON(http.StatusOK, gin.H{"letter": string(b)})
}
