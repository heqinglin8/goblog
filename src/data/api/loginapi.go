package api

import "github.com/kataras/iris"

/**
首页第一个方法
 */
func Login(ctx iris.Context){
	ctx.JSON(iris.Map{
		"message": "pong",
	})
}
