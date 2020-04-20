package main

import (
	"gin-sam/bootstrap"
	"gin-sam/web/routes"
)

func newApp()  *bootstrap.Bootstrapper{
	app:= bootstrap.New("lottery","yookie")
	app.BootStrap()
	return app
}

func main() {

	app := newApp()
	app.Configure(routes.Configure)
	app.Start(":8080")
}
