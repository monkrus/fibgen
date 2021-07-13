package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return (Fibonacci(n-1) + Fibonacci(n-2))
	}
}

func main() {

	var n, i, j int
	j = 0

	fmt.Print("Enter the number of terms: ")
	fmt.Scanln(&n)

	fmt.Println("Fibonacci series: ")
	for i = 1; i <= n; i++ {
		fmt.Print(Fibonacci(j), ",")
		j++
	}
	fmt.Println()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	
	/*func (c *Context) GetInt(key string) (i int) {
		if val, ok := c.Get(key); ok && val != nil {
			i, _ = val.(int)
		}
		return
	}
	*/

}
