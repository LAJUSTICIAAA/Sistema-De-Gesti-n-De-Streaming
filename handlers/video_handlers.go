package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/LAJUSTICIAAA/Sistema-De-Gesti-n-De-Streaming/models"
	"github.com/LAJUSTICIAAA/Sistema-De-Gesti-n-De-Streaming/services"
	"github.com/gin-gonic/gin"
)

func GetVideos(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetVideos())
}

func PostVideo(c *gin.Context) {
	var video models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	services.AddVideo(video)
	c.JSON(http.StatusCreated, video)
}

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Archivo no recibido"})
		return
	}
	os.MkdirAll("static", os.ModePerm)
	path := filepath.Join("static", filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Video subido correctamente", "ruta": path})
}
