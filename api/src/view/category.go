package view

import (
    "time"
    "net/http"

    "github.com/labstack/echo"
	
    . "app/controller"
)

type Category struct {
    Id int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
    Name string `json:"name" gorm:"type:varchar(127)"`
    Description string `json:"description" gorm:"type:varchar(255)"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func GetAllCategories(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Category{})

    var categories []Category
    db.Find(&categories)
    return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Category{})

    if id := c.Param("id"); id != "" {
        var category Category
        db.First(&category, id)
        return c.JSON(http.StatusOK, category)
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}

func CreateCategory(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Category{})

    category := new(Category)
    if err := c.Bind(category); err != nil {
        return err
    }
    db.Create(&category)
    return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()

    new_category := new(Category)
    if err := c.Bind(new_category); err != nil {
        return err
    }

    if id := c.Param("id"); id != "" {
        var category Category
        db.First(&category, id).Update(new_category)
        return c.JSON(http.StatusOK, category)
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}

func DeleteCategory(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()

    if id := c.Param("id"); id != "" {
        var category Category
        db.First(&category, id)
        db.Delete(category)
        return c.JSON(http.StatusOK, category)
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}