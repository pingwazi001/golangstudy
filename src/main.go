package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	r.Use(RequestTime)
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("hello go")
		c.String(http.StatusOK, "hello go")
	})
	r.Run(":8080")
}

func RequestTime(c *gin.Context) {
	start := time.Now()
	c.Next()//执行下一个处理函数
	timeCost := time.Since(start).Milliseconds()
	fmt.Println("请求耗时：", timeCost)
}
