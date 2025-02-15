package controllers

import (
	"btc_api/fibonacci"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FibonacciController (c *gin.Context) {
	number := c.Param("number")
	n, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	ch := make(chan uint64)
	go fibonacci.Fibonacci(uint64(n), ch)

	c.JSON(200, gin.H{
		"number": n,
		"fibonacci": <- ch,
	})
}