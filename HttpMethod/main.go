package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// デフォルトのミドルウェアで新しい gin ルーターを作成する
	// logger とアプリケーションクラッシュをキャッチする recovery ミドルウェア
	router := gin.Default()

	router.GET("/someGet", func(c *gin.Context) {
		fmt.Println("get")
	})
	router.POST("/somePost", func(c *gin.Context) {
		fmt.Println("post")
	})
	router.PUT("/somePut", func(c *gin.Context) {
		fmt.Println("put")
	})
	router.DELETE("/someDelete", func(c *gin.Context) {
		fmt.Println("delete")
	})
	router.PATCH("/somePatch", func(c *gin.Context) {
		fmt.Println("patch")
	})
	router.HEAD("/someHead", func(c *gin.Context) {
		fmt.Println("head")
	})
	router.OPTIONS("/someOptions", func(c *gin.Context) {
		fmt.Println("options")
	})

	// デフォルトではポート 8080 が利用されるが、
	// 環境変数 PORT を指定していればそちらが優先される。
	router.Run()
	// router.Run(":3000") と書くことでポートをハードコーディングできる
}
