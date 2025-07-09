
// backend/database/db.go
package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "sykell-url-analyzer/models"
)

func InitDB() *gorm.DB {
    dsn := "root:password@tcp(mysql:3306)/sykell?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&models.Analysis{})
    return db
}
