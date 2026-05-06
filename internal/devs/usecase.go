package devs

import (
	"context"
	"log"
	"net/http"
	"net/mail"

	"github.com/CriciumaDevJobs/backend/handlers"
	"github.com/CriciumaDevJobs/backend/utils"
)

type DevUseCase struct {
	Repository *Queries
}

func NewDevUseCase(devRepository *Queries) *DevUseCase {
	usecase := DevUseCase{
		Repository: devRepository,
	}

	return &usecase
}

func (usecase *DevUseCase) CreateDev(ctx context.Context, dev *Dev) (*CreateDevRow, *handlers.ErrorResponse) {

	httpErr := dev.validate()

	if httpErr != nil {
		return nil, httpErr
	}

	row_count, err := usecase.DevExistsByEmail(ctx, dev.Email)

	if err != nil {
		return nil, err
	}

	if row_count > 0 {
		return nil, handlers.ErrEmailAlreadyInUse
	}

	hashedPassword, httpErr := utils.EncryptPassword(dev.Password)

	if httpErr != nil {
		return nil, httpErr
	}

	var devParams = CreateDevParams{
		Name:         dev.Name,
		Email:        dev.Email,
		Password:     hashedPassword,
		Skills:       dev.Skills,
		Bio:          dev.Bio,
		Availability: dev.Availability,
		Socials:      dev.Socials,
	}

	resp, dbErr := usecase.Repository.CreateDev(ctx, devParams)

	if dbErr != nil {
		log.Printf("ERRO: Falha no banco de dados ao salvar novo usuário! Message: %s", dbErr.Error())
		return nil, handlers.NewError(http.StatusInternalServerError, "Erro Interno!")
	}

	return &resp, nil
}

func (usecase *DevUseCase) FindDevByEmail(ctx context.Context, email string) (*FindDevByEmailRow, *handlers.ErrorResponse) {
	dev, err := usecase.Repository.FindDevByEmail(ctx, email)

	if err != nil {
		return nil, handlers.ErrInvalidEmailOrPassword
	}

	return &dev, nil
}

func (usecase *DevUseCase) FindDevByID(ctx context.Context, id int32) (*FindDevByIDRow, *handlers.ErrorResponse) {
	dev, err := usecase.Repository.FindDevByID(ctx, id)

	if err != nil {
		return nil, handlers.ErrProfileNotFound
	}

	return &dev, nil
}

func (usecase *DevUseCase) DevExistsByEmail(ctx context.Context, email string) (int64, *handlers.ErrorResponse) {
	row_count, err := usecase.Repository.EmailAlreadyRegistered(ctx, email)

	if err != nil {
		log.Printf("ERRO: Falha ao executar busca no banco de dados! Message: %s", err.Error())
		return 0, handlers.NewError(http.StatusInternalServerError, "Erro Interno!")
	}

	return row_count, nil
}

func (dev *Dev) validate() *handlers.ErrorResponse {
	if dev.Name == "" {
		return handlers.ErrNameNotEmpty
	}

	if dev.Email == "" {
		return handlers.ErrEmailNotEmpty
	}

	_, err := mail.ParseAddress(dev.Email)

	if err != nil {
		return handlers.ErrEmailAddressNotValid
	}

	if dev.Bio == "" {
		return handlers.ErrBioNotEmpty
	}

	return nil
}
