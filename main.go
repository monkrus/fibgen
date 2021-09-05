package main

import (
	"math"

	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func Fibonacci(n int) (retValue int) {
	defer func() {
		if r := recover(); r != nil {
			retValue = -1
		}
	}()
	if (math.MaxInt64 / 2) < n {
		panic("Out of bound access for function!")
	}
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func setupRouter() *MyGinServer {
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
