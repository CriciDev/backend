package auth

import (
	"context"
	"time"

	"github.com/CriciumaDevJobs/backend/internal/devs"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationUseCase struct {
	DevUseCase *devs.DevUseCase
}

func NewAuthenticationUseCase(devUseCase *devs.DevUseCase) *AuthenticationUseCase {
	auth := AuthenticationUseCase{
		DevUseCase: devUseCase,
	}

	return &auth
}

func (usecase *AuthenticationUseCase) AuthenticateUser(ctx context.Context, email string, password string) (*AuthenticationResponse, error) {
	dev, err := usecase.DevUseCase.FindDevByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	bcrypterr := bcrypt.CompareHashAndPassword([]byte(dev.Password), []byte(password))

	if bcrypterr != nil {
		return nil, devs.ErrInvalidEmailOrPassword
	}

	expiration := time.Now().Add(time.Hour * 4)

	token, err := GenerateJwtToken(dev, expiration)

	if err != nil {
		return nil, err
	}

	authResponse := AuthenticationResponse{
		Token:     token,
		ExpiresAt: expiration,
	}

	return &authResponse, nil
}
