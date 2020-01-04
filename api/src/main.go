package main

import (
    "github.com/labstack/echo"

    view "app/view"
)

func main() {
    e := echo.New()

    initRouting(e)

    e.Logger.Fatal(e.Start(":1323"))
}

func initRouting(e *echo.Echo) {
    // posts
    e.GET("/posts", view.GetAllPosts)
    e.GET("/posts/:id", view.GetPost)
    e.POST("/posts", view.CreatePost)
    e.PUT("/posts/:id", view.UpdatePost)
    e.DELETE("/posts/:id", view.DeletePost)

    // categories
    e.GET("/categories", view.GetAllCategories)
    e.GET("/categories/:id", view.GetCategory)
    e.POST("/categories", view.CreateCategory)
    e.PUT("/categories/:id", view.UpdateCategory)
    e.DELETE("/categories/:id", view.DeleteCategory)
}