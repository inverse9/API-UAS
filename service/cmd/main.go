package main

import (
	"fmt"

	"wahyu/e-commerce/core"
	"wahyu/e-commerce/service/router"
)

func main() {
	app := core.NewApp()

	env := app.Env
	gin := app.Web

	router := router.RouterConstructor(gin)
	router.NewRouter()

	gin.Run(fmt.Sprintf(":%s", env.Port))
}
