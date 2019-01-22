package controller

import (
	"github.com/kataras/iris/mvc"
	"../data/datasource"
)

// IndexController is our sample controller
// it handles GET: /hello and GET: /hello/{name}
type IndexController struct{}

var helloView = mvc.View{
	Name: "template/index.html",
	Data: map[string]interface{}{
		"Title":     "Hello Page",
		"MyMessage": "Welcome to my awesome website",
		"moves":datasource.Movies,
	},
}

// Get will return a predefined view with bind data.
//
// `mvc.Result` is just an interface with a `Dispatch` function.
// `mvc.Response` and `mvc.View` are the built'n result type dispatchers
// you can even create custom response dispatchers by
// implementing the `github.com/kataras/iris/hero#Result` interface.
func (c *IndexController) Get() mvc.Result {
	return helloView
}
