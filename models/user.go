package models

import (
	"time"

	"github.com/LAJUSTICIAAA/Sistema-De-Gesti-n-De-Streaming/utils"
	"github.com/google/uuid"
)

type Usuario struct {
	ID            string
	Nombre        string
	Email         string
	PasswordHash  string
	FechaRegistro time.Time
}

func NewUsuario(nombre, email, password string) (*Usuario, error) {
	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &Usuario{
		ID:            uuid.NewString(),
		Nombre:        nombre,
		Email:         email,
		PasswordHash:  hash,
		FechaRegistro: time.Now(),
	}, nil
}

func (u *Usuario) VerificarPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.PasswordHash)
}
