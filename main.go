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
	if n <= 1 {
		return n
	}
	return (Fibonacci(n-1) + Fibonacci(n-2))
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/get", func(c *gin.Context) {
		number := c.DefaultQuery("number", "0")
		n, err := strconv.Atoi(number)
		if err != nil {
			c.String(
				http.StatusBadRequest, err.Error())
			return
		}
		switch {
		case n == 0:
			c.String(
				http.StatusBadRequest, "number should be a positive number, but it is zero")
		case n < 0:
			c.String(
				http.StatusBadRequest, "number should be a positive number, but it is negative")
		case n >= 922337203685477580:
			c.String(
				http.StatusBadRequest, "number should be a number, but it is exceeds the maximum value if int")
		default:
			c.String(http.StatusOK, fmt.Sprintf("Result is %d", Fibonacci(n)))
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
