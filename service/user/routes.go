package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sakhimpungose/ecom/service/auth"
	"github.com/sakhimpungose/ecom/types"
	"github.com/sakhimpungose/ecom/utils"
)

type Handler struct{
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var data types.RegisterUserDto
	if err := utils.ParseJSON(r, data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check if the user exists
	if _, err := h.store.GetUserByEmailAddress(data.EmailAddress); err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email address %s already exists", data.EmailAddress))
		return
	}

	hashedPassword, err := auth.HashedPassword(data.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		EmailAddress: data.EmailAddress,
		Password:     hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created successfully"})
}
