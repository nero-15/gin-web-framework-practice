package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 別の prefix を使うこともできます
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// while(1);["lena","austin","foo"] が出力されます。
		c.SecureJSON(http.StatusOK, names)
	})

	// 0.0.0.0:8080 でサーバーを立てます。
	r.Run(":8080")
}
