package handlers

import "errors"

var (
	ErrNameNotEmpty      = errors.New("Campo Nome deve ser preenchido!")
	ErrEmailNotEmpty     = errors.New("Campo Email deve ser preenchido!")
	ErrEmailAlreadyInUse = errors.New("Email já está em uso")
	ErrBioNotEmpty       = errors.New("Campo Biografia deve ser preenchido")
	ErrPasswordToLong    = errors.New("O Campo de Senha não pode ultrapassar 72 caracteres")
)
