package controllers

import (
	"github.com/kataras/iris"
)

func Index(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	ctx.View("shop/index.html")
}
