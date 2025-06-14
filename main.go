package main

import (
	"fmt"
	"sistema/models"
	"sistema/services"
)

func main() {
	// Datos de prueba
	nombre := "Justin Salazar"
	email := "jusalazarro@uide.edu.ec"
	contrasena := "pechamen12"

	// Crear nuevo usuario
	usuario, err := models.NuevoUsuario(nombre, email, contrasena)
	if err != nil {
		fmt.Println("❌ Error al crear usuario:", err)
	} else {
		fmt.Println("✅ Usuario creado con éxito")
		fmt.Println("Nombre:", usuario.GetNombre())
		fmt.Println("Email:", usuario.GetEmail())
		fmt.Println("Fecha de registro:", usuario.GetFechaRegistro())

		// Verificar contraseña correcta
		fmt.Println("\nProbando contraseña correcta:")
		if usuario.VerificarPassword(contrasena) {
			fmt.Println("🔐 Contraseña válida")
		} else {
			fmt.Println("❌ Contraseña incorrecta")
		}

		// Verificar contraseña incorrecta
		fmt.Println("\nProbando contraseña incorrecta:")
		if usuario.VerificarPassword("micontraseña123") {
			fmt.Println("⚠️ ¡Esto no debería pasar!")
		} else {
			fmt.Println("✅ Contraseña incorrecta detectada correctamente")
		}

		// Subiendo un video de prueba
		fmt.Println("\n--- Subiendo un video de prueba ---")
		video := models.NuevoVideo(
			"Mi primer video",
			"Aprendizaje Autonomo 2 ",
			"video1.mp4",
			"data/videos/video1.mp4",
		)

		// Mostrar datos del video
		fmt.Println("📼 Video creado:")
		fmt.Println("ID:", video.GetID())
		fmt.Println("Título:", video.GetTitulo())
		fmt.Println("Descripción:", video.GetDescripcion())
		fmt.Println("Archivo:", video.GetNombreArchivo())
		fmt.Println("Ruta:", video.GetRutaArchivo())

		// Crear gestor e implementar interfaz
		gestor := services.NuevoServicioVideo()

		// Subir el video
		gestor.Subir(video)

		// Listar todos los videos
		gestor.Listar()
	}
}
