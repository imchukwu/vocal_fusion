package handlers

import (
	"encoding/json"
	"net/http"

	"vocal_fusion/internals/models"
	"vocal_fusion/internals/repository"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepo: userRepo}
}

// Register User
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.UserRepo.CreateUser(&user); err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

// Login user (simple placeholder response for now)
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login endpoint - JWT coming soon 🔐"))
}

// func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
// 	var input struct {
// 		Email string `json:"email"`
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := h.UserRepo.GetUserByEmail(input.Email)
// 	if err != nil {
// 		http.Error(w, "User not found", http.StatusUnauthorized)
// 		return
// 	}

// 	token, err := auth.GenerateToken(user.ID)
// 	if err != nil {
// 		http.Error(w, "Could not generate token", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"token": token,
// 	})
// }
