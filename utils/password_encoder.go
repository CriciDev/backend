package utils

import (
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/CriciumaDevJobs/backend/handlers"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, *handlers.ErrorResponse) {

	//bcrypt encoder trunkate strings with more than 72 characters
	if utf8.RuneCountInString(password) > 72 {
		return "", handlers.ErrPasswordToLong
	}

	hashedpassword, bcryptErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if bcryptErr != nil {
		log.Printf("ERRO: Falha ao encriptar senha! Message: %s", bcryptErr.Error())
		return "", handlers.NewError(http.StatusInternalServerError, "Erro interno!")
	}

	return string(hashedpassword), nil
}
