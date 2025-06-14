package models

import (
	"github.com/google/uuid"
)

// Video representa un archivo de video subido al sistema
type Video struct {
	id            string
	titulo        string
	descripcion   string
	nombreArchivo string
	rutaArchivo   string
}

// NuevoVideo crea y devuelve un nuevo objeto Video
func NuevoVideo(titulo, descripcion, nombreArchivo, rutaArchivo string) *Video {
	return &Video{
		id:            uuid.NewString(),
		titulo:        titulo,
		descripcion:   descripcion,
		nombreArchivo: nombreArchivo,
		rutaArchivo:   rutaArchivo,
	}
}

// Getters para acceder a los campos privados

func (v *Video) GetID() string {
	return v.id
}

func (v *Video) GetTitulo() string {
	return v.titulo
}

func (v *Video) GetDescripcion() string {
	return v.descripcion
}

func (v *Video) GetNombreArchivo() string {
	return v.nombreArchivo
}

func (v *Video) GetRutaArchivo() string {
	return v.rutaArchivo
}
