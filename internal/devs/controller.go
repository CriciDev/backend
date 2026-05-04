package devs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type DevController struct {
	Usecase *DevUseCase
}

func NewDevController(usecase *DevUseCase) *DevController {
	controller := DevController{
		Usecase: usecase,
	}

	return &controller
}

func (controller *DevController) CreateDev(ctx context.Context, writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dev = Dev{}

	err := json.NewDecoder(request.Body).Decode(&dev)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = controller.Usecase.CreateDev(ctx, &dev)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(writer, "Desenvolvedor salvo com sucesso")
}
