package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.QueryMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	router.Run(":8080")
}
