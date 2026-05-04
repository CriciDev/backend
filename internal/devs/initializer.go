package devs

import "database/sql"

func InitializeDevController(db *sql.DB) *DevController {
	repository := NewDevRepository(db)
	usecase := NewDevUseCase(repository)
	controller := NewDevController(usecase)

	return controller
}
