package controller

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func OpenSQLiteConnection() *gorm.DB {
    //db, err := gorm.Open("sqlite3", ":memory:")
    db, err := gorm.Open("sqlite3", "a.sqlite3")
    if err != nil {
        panic("failed to connect database.")
    }
    
    db.LogMode(true)
    return db
}
