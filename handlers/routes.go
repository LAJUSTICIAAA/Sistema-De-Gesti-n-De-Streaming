package handlers

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.GET("/users", ListUsers)

	router.GET("/videos", GetVideos)
	router.POST("/videos", PostVideo)
	router.POST("/upload", UploadVideo)
}
