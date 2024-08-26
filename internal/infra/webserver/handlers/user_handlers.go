package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/savioafs/simpleAPIGo/internal/dto"
	"github.com/savioafs/simpleAPIGo/internal/entity"
	"github.com/savioafs/simpleAPIGo/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserStorer
}

func NewUserHandler(db database.UserStorer) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
