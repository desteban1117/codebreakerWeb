package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/codebreaker/guess/:number", func(c *gin.Context) {
		number := c.Param("number")
    response := codebreaker(number)
		c.String(http.StatusOK, "Hello %s", response)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/codebreaker/code/:number", func(c *gin.Context) {
		number := c.Param("number")
    setSecret(number)
		c.String(http.StatusOK, "Hello")
	})

	router.Run(":8080")
}
