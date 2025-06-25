package handlers

import (
	"net/http"

	"github.com/LAJUSTICIAAA/Sistema-De-Gesti-n-De-Streaming/models"
	"github.com/gin-gonic/gin"
)

var usuarios []models.Usuario

func Register(c *gin.Context) {
	var datos struct {
		Nombre   string `json:"nombre"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&datos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	usuario, err := models.NewUsuario(datos.Nombre, datos.Email, datos.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar"})
		return
	}

	usuarios = append(usuarios, *usuario)
	c.JSON(http.StatusCreated, gin.H{"mensaje": "Usuario registrado correctamente"})
}

func Login(c *gin.Context) {
	var datos struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&datos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	for _, u := range usuarios {
		if u.Email == datos.Email && u.VerificarPassword(datos.Password) {
			c.JSON(http.StatusOK, gin.H{"mensaje": "Inicio de sesión exitoso"})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
}

func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, usuarios)
}
