package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return (Fibonacci(n-1) + Fibonacci(n-2))
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/get", func(c *gin.Context) {
		number := c.DefaultQuery("number", "0")
		n, err := strconv.Atoi(number)
		if err != nil {
			c.String(
				http.StatusBadRequest, err.Error())
		} else {
			if n == 0 {
				c.String(
					http.StatusBadRequest, "number should be a positive number, but it is zero")
			} else if n < 0 {
				c.String(
					http.StatusBadRequest, "number should be a positive number, but it is negative")
			} else if n > 9223372036854775807 {
				c.String(
					http.StatusBadRequest, "number should be a number, but it is exceeds the maximum value if int")
			} else {
				c.String(http.StatusOK, fmt.Sprintf("Result is %d", Fibonacci(n)))
			}
		}

	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
	//coloration for logging
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
    //sample logging
	log.WithFields(log.Fields{
		"user": "admin",
	}).Info("Some info")

	log.Warn("This is a warning")
	log.Error("An error occured!")
}
