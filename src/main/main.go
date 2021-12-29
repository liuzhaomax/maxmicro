/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/12/25 2:23
 * @version     v1.0
 * @filename    main.go
 * @description
 ***************************************************************************/
package main

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/web"
)

func main() {
	router := gin.Default()
	router.Handle("GET", "/user", func(context *gin.Context) {
		context.String(200, "user api")
	})
	router.Handle("GET", "/news", func(context *gin.Context) {
		context.String(200, "news api")
	})
	server := web.NewService(
		web.Address(":8000"),
		web.Handler(router),
	)
	server.Run()
}
