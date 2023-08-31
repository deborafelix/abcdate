package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = map[string][]string{"A": []string{"Abigail", "Acampar", "Academia"}}

func getPlace(c *gin.Context) {
	letter := c.Param("letter")

	places := db[letter]

	c.IndentedJSON(http.StatusOK, gin.H{"place": places})
}
