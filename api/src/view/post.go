package view

import (
    "time"
    "net/http"

    "github.com/labstack/echo"
	
    . "app/controller"
)

type Post struct {
    Id int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
    Title string `json:"title" gorm:"type:varchar(127)"`
    Text string `json:"text"`
    CategoryId int `json:"category_id"`
    Category Category `json:"category" gorm:"foreignkey:Id;association_foreignkey:CategoryId"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func GetAllPosts(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Post{})

    var posts []Post
    db.Find(&posts)
    return c.JSON(http.StatusOK, posts)
}

func GetPost(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Post{})

    if id := c.Param("id"); id != "" {
        var post Post
        db.First(&post, id)
        return c.JSON(http.StatusOK, post)
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}

func CreatePost(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Post{})

    post := new(Post)
    if err := c.Bind(post); err != nil {
        return err
    }

    db.Create(&post)
    return c.JSON(http.StatusOK, post)
}

func UpdatePost(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()

    new_post := new(Post)
    if err := c.Bind(new_post); err != nil {
        return err
    }

    if id := c.Param("id"); id != "" {
        var post Post
        db.First(&post, id).Update(new_post)
        return c.JSON(http.StatusOK, post)
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}

func DeletePost(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()

    if id := c.Param("id"); id != "" {
        var post Post
        db.First(&post, id)
        db.Delete(post)
        return c.JSON(http.StatusOK, post)
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}