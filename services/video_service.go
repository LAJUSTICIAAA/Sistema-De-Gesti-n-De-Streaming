package services

import (
	"fmt"
	"sistema/models"
)

// Interfaz GestorContenido
type GestorContenido interface {
	Subir(video *models.Video)
	Listar()
}

// Implementación del gestor de videos
type ServicioVideo struct {
	videos []*models.Video
}

// Constructor del servicio
func NuevoServicioVideo() *ServicioVideo {
	return &ServicioVideo{
		videos: make([]*models.Video, 0),
	}
}

// Subir implementa la interfaz
func (s *ServicioVideo) Subir(video *models.Video) {
	s.videos = append(s.videos, video)
	fmt.Println("✅ Video subido:", video.GetTitulo())
}

// Listar implementa la interfaz
func (s *ServicioVideo) Listar() {
	fmt.Println("\n📂 Lista de videos:")
	if len(s.videos) == 0 {
		fmt.Println("No hay videos aún.")
	}
	for i, v := range s.videos {
		fmt.Printf("%d. %s (%s)\n", i+1, v.GetTitulo(), v.GetNombreArchivo())
	}
}
