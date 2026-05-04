package devs

import "database/sql"

func InitializeDevController(db *sql.DB) *DevController {
	repository := New(db)
	usecase := NewDevUseCase(repository)
	controller := NewDevController(usecase)

	return controller
}
