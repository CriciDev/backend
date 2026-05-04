package devs

import (
	"time"

	"github.com/CriciumaDevJobs/backend/handlers"
)

type Dev struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Skills       string   `json:"skills"`
	Bio          string   `json:"bio"`
	Availability bool     `json:"availability"`
	Socials      []string `json:"socials"`
	CreatedAt    string   `json:"created_at"`
}

func NewDev(dev *Dev) (*Dev, error) {
	newDev := Dev{
		Name:         dev.Name,
		Email:        dev.Email,
		Skills:       dev.Skills,
		Bio:          dev.Bio,
		Availability: dev.Availability,
		Socials:      dev.Socials,
		CreatedAt:    time.Now().String(),
	}

	err := newDev.validate()

	if err != nil {
		return nil, err
	}

	return &newDev, nil
}

func (dev *Dev) validate() error {
	if dev.Name == "" {
		return handlers.ErrNameNotEmpty
	}

	if dev.Email == "" {
		return handlers.ErrEmailNotEmpty
	}

	return nil
}
