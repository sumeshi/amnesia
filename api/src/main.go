package main

import (
    "github.com/labstack/echo"

    view "app/view"
)

func main() {
    e := echo.New()

    // routings
    initRouting(e)

    e.Logger.Fatal(e.Start(":1323"))
}

func initRouting(e *echo.Echo) {
    e.GET("/", view.Hello)
}