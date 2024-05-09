package main

import (
	"fmt"

	"github.com/Bonk123123/notes_server/internal/database"
	"github.com/Bonk123123/notes_server/internal/middlewares"
	"github.com/Bonk123123/notes_server/internal/models"
	"github.com/Bonk123123/notes_server/internal/services"
	"github.com/gin-gonic/gin"
)






func main() {
	router := gin.Default()
	database.ConnectDB()

	database.Db.AutoMigrate(&models.User{})
	database.Db.AutoMigrate(&models.Note{})
	database.Db.AutoMigrate(&models.NoteElement{})

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("", services.Auth)
	publicRoutes.POST("/register", services.Registration)
	publicRoutes.POST("/refresh", services.Refresh)

	protectedRoutes := router.Group("/note")
	protectedRoutes.Use(middlewares.JWTAuthMiddleware())

	protectedRoutes.GET("", services.GetNotes)
	protectedRoutes.GET("/:id", services.GetNote)
	protectedRoutes.POST("", services.CreateNote)
	protectedRoutes.PUT("/:id", services.UpdateNote)
	protectedRoutes.DELETE("", services.DeleteNote)

	protectedRoutes.GET("/element", services.GetNoteElements)
	protectedRoutes.GET("/element/:id", services.GetNoteElement)
	protectedRoutes.POST("/element", services.CreateNoteElement)
	protectedRoutes.PUT("/element/:id", services.UpdateNoteElement)
	protectedRoutes.DELETE("/element", services.DeleteNoteElement)

	router.Run(":8000")
    fmt.Println("Server running on port 8000")
}


// /auth {login, password} -> {access, refresh}
// /auth/register {login, password, password_repeat} -> {access, refresh}
// /auth/refresh {refresh} -> {access, refresh}

// /notes headers: {Authorization: access} -> notes
// ...