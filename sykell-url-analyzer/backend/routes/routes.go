
// backend/routes/routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "sykell-url-analyzer/controllers"
    "sykell-url-analyzer/middleware"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    r.Use(middleware.AuthMiddleware())

    api := r.Group("/api")
    {
        api.POST("/analyze", controllers.AnalyzeURL(db))
        api.GET("/urls", controllers.GetURLs(db))
        api.GET("/urls/:id", controllers.GetURLDetail(db))
    }
}
