package devs

import (
	"context"
)

type DevUseCase struct {
	Repository *DevRepository
}

func NewDevUseCase(devRepository *DevRepository) *DevUseCase {
	usecase := DevUseCase{
		Repository: devRepository,
	}

	return &usecase
}

func (usecase *DevUseCase) CreateDev(ctx context.Context, dev *Dev) error {
	return usecase.Repository.CreateDev(ctx, dev)
}
