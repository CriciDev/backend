package devs

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/CriciumaDevJobs/backend/handlers"
)

func InitializeDevContext(db *sql.DB) {
	controller := initializeDevController(db)
	initializeDevEndpoints(controller)
}

func initializeDevController(db *sql.DB) *DevController {
	repository := New(db)
	usecase := NewDevUseCase(repository)
	controller := NewDevController(usecase)

	return controller
}

func initializeDevEndpoints(controller *DevController) {
	http.HandleFunc("/devs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.CreateDev(context.Background(), w, r)

		default:
			handlers.ResponseWithHttpError(w, http.StatusMethodNotAllowed, "Method not Allowed")
		}
	})
}
