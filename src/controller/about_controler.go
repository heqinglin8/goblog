package controller

import (
	"github.com/kataras/iris"
)

type mypage struct {
	Title   string
	Message string
}


func RenderAbout(ctx iris.Context) {
	ctx.Gzip(true)
	ctx.ViewData("", mypage{"My Page title", "Hello world!"})
	ctx.View("template/mypage.html")
	// Note that: you can pass "layout" : "otherLayout.html" to bypass the config's Layout property
	// or view.NoLayout to disable layout on this render action.
	// third is an optional parameter
}
