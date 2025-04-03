package controllers

import (
	"btc_api/fibonacci"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FibonacciController (c *gin.Context) {
	ch := make(chan uint64)
	number := c.Param("number")
	n, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	
	go fibonacci.Fibonacci(ch)
	ch <- uint64(n)
	c.JSON(200, gin.H{
		"number": n,
		"fibonacci": <- ch,
	})
}