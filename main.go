package main

import (
	"coolgo/loader"
	"coolgo/routers"
)

var App *loader.App

func main()  {
	App = &loader.App{}
	App.SysInit()
	router := routers.InitRouter(App)
	router.Run()
}

