package api

import "github.com/kataras/iris"

/**
首页第一个方法
 */
func Index(ctx iris.Context){
	ctx.WriteString("这是首页 index")
}