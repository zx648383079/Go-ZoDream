package gzo

import (
	"zodream/modules/gzo/controllers"

	"github.com/kataras/iris/v12"
)

func Register(app iris.Party) {
	app.Get("/", controllers.Index)
}
