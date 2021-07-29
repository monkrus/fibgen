package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	r.GET("/get", func(c *gin.Context) {
		//number := c.DefaultQuery("number", "0")
		//c.String(http.StatusOK, number)
		number := c.DefaultQuery("number", "0")
		n, err := strconv.Atoi(number)
		if err != nil {
			c.String(
				http.StatusBadRequest, "number should be a number")
		} else {
			c.String(http.StatusOK, fmt.Sprintf("Result is %d", Fibonacci(n)))
		}
		if err := r.Run(":8080"); err != nil {
			log.Fatal(err.Error())
		}

	})
	r.GET("/neg", func(c *gin.Context) {
		negnumber := c.DefaultQuery("negnumber", "unknown")
		n, err := strconv.Atoi(negnumber)
		if err != nil {
			c.String(http.StatusBadRequest, "number should be a  negative number")
		} else {
			c.String(http.StatusOK, fmt.Sprintf("Result is %d", Fibonacci(n)))
		}
		if err := r.Run(":8080"); err != nil {
			log.Fatal(err.Error())
		}
	})

	r.GET("/get", func(c *gin.Context) {

		maxnumber := c.DefaultQuery("maxnumber", "0")
		n, err := strconv.Atoi(maxnumber)
		if err != nil {
			c.String(
				http.StatusBadRequest, "maxnumber should be equalt to 2147483647")
		} else {
			c.String(http.StatusOK, fmt.Sprintf("Result is %d", Fibonacci(n)))
		}
		if err := r.Run(":8080"); err != nil {
			log.Fatal(err.Error())
		}
	})

}
