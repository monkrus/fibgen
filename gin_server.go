package main

import (
	"github.com/gin-gonic/gin"
)

type MyGinServer struct {
	*gin.Engine
	MySpecName string
}

//constructor
func NewMyGinServer(specName string) *MyGinServer {
	return &MyGinServer{
		Engine:     gin.Default(),
		MySpecName: specName,
	}
}

// set up initial routes
func (engine *MyGinServer) InitRoutes(routes []RouteDescriptor) {
	for _, route := range routes {
		switch route.MethodType {
		case HTTPGet:
			engine.GET(route.Path, route.Method)
		case HTTPPost:
			engine.POST(route.Path, route.Method)
		case HTTPDelete:
			engine.DELETE(route.Path, route.Method)
		case HTTPPut:
			engine.PUT(route.Path, route.Method)
		}
	}
}
