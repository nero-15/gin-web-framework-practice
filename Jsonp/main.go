package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// x({\"foo\":\"bar\"}) が出力されます。
		c.JSONP(http.StatusOK, data)
	})

	// 0.0.0.0:8080 でサーバーを立てます。
	r.Run(":8080")
}
