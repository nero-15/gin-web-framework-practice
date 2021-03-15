package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// v1 のグループ
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{"login": "login"})
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.JSON(200, gin.H{"submit": "submit"})
		})
		v1.POST("/read", func(c *gin.Context) {
			c.JSON(200, gin.H{"read": "read"})
		})
	}

	// v2 のグループ
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{"login": "login"})
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.JSON(200, gin.H{"submit": "submit"})
		})
		v2.POST("/read", func(c *gin.Context) {
			c.JSON(200, gin.H{"read": "read"})
		})
	}

	router.Run(":8080")
}
