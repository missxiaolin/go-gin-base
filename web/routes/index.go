package routes

import (
	"go-gin-base/bootstrap"
	"go-gin-base/web/controller"
)

func ApiConfigure(b *bootstrap.Bootstrapper)  {
	d := b.Group("/api")

	d.GET("/index", new(controller.Index).Welcome)
}