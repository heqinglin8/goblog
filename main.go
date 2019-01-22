package main

import (
	"./src/controller"
	"./src/data/api"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.Default()
	// 从 "./views" 文件夹加载所以的模板
	// 其中扩展名为“.html”并解析它们
	// 使用标准的`html / template`包。
	//app.RegisterView(iris.HTML("./src/views/template", ".html"))
	app.RegisterView(iris.HTML("./src/views/", ".html"))
	mvc.New(app.Party("/index")).Handle(new(controller.IndexController))
	mvc.New(app.Party("/")).Handle(new(controller.IndexController))
	app.Get("/about",controller.RenderAbout)
	app.Get("/login", api.Login)
	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8080"))
}
