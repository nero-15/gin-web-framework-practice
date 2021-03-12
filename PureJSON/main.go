package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Unicode を返します
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// そのままの文字を返します
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 0.0.0.0:8080 でサーバーを立てます。
	r.Run(":8080")
}
