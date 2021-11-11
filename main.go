package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// covers big and negative number scenarios
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if n >= 25 {
		return -1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func setupRouter() *MyGinServer {
	gin.SetMode(gin.ReleaseMode)
	r := NewMyGinServer("test name")

	r.InitRoutes([]RouteDescriptor{
		{Path: UrlPong, Method: Pong, MethodType: HTTPGet},
		{Path: UrlFib, Method: Fib, MethodType: HTTPGet},
	})

	return r
}

func main() {
	err := setupRouter().Run(":8080")
	if err != nil {
		log.Fatalf("Internal error: %v", err)

	}
}
