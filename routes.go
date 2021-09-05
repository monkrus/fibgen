package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	UrlPong = "/pong"
	UrlFib  = "/fib"
)

const (
	NonPositiveErr = "number should be a positive number, but it is negative"
	MaxValueErr    = "number should be a number, but it exceeds the maximum value if int"
)

const (
	HTTPGet HTTPType = iota
	HTTPPost
	HTTPPut
	HTTPDelete
)

type HTTPType int

type RouteDescriptor struct {
	Path       string
	Method     func(C *gin.Context)
	MethodType HTTPType
}

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Fib(c *gin.Context) {
	number := c.DefaultQuery("number", "0")
	n, err := strconv.Atoi(number)
	if err != nil {
		c.String(
			http.StatusBadRequest, err.Error())
		return
	}
	switch {
	case n <= 0:
		c.String(
			http.StatusBadRequest, NonPositiveErr)
	case n >= 9223372036854775807:
		c.String(
			http.StatusBadRequest, MaxValueErr)
	default:
		if result := Fibonacci(n); result == -1 {
			c.String(http.StatusInternalServerError, "current value exceeds the limit")
		}
		c.String(http.StatusOK, fmt.Sprintf("Result is %d", Fibonacci(n)))
	}

}
