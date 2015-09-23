package main

import (
	"github.com/heroku/go-getting-started/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/wait", func(c *gin.Context) {
		second_str := c.DefaultQuery("seconds", "1")
		seconds, _ := strconv.Atoi(second_str)
		time.Sleep(time.Duration(seconds) * time.Second)
		c.String(http.StatusOK, "Done")
	})

	router.Run(":" + port)
}
