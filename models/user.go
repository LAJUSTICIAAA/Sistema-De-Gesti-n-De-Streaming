package models

import (
	"errors"
	"sistema/utils"
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	id            string
	nombre        string
	email         string
	passwordHash  string
	fechaRegistro time.Time
}

// Constructor
func NuevoUsuario(nombre, email, password string) (*Usuario, error) {
	if nombre == "" || email == "" || password == "" {
		return nil, errors.New("todos los campos son obligatorios")
	}
	if len(password) < 6 {
		return nil, errors.New("la contraseña debe tener al menos 6 caracteres")
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &Usuario{
		id:            uuid.NewString(),
		nombre:        nombre,
		email:         email,
		passwordHash:  hash,
		fechaRegistro: time.Now(),
	}, nil
}

// Getters
func (u *Usuario) GetID() string {
	return u.id
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) GetEmail() string {
	return u.email
}

func (u *Usuario) GetFechaRegistro() time.Time {
	return u.fechaRegistro
}

// Verificar contraseña
func (u *Usuario) VerificarPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.passwordHash)
}
