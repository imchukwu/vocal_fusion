package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "vocal_fusion/internals/models"
    "vocal_fusion/internals/repository"

    "github.com/go-chi/chi/v5"
)

type WinnerSaysHandler struct {
    repo *repository.WinnerSaysRepository
}

func NewWinnerSaysHandler(repo *repository.WinnerSaysRepository) *WinnerSaysHandler {
    return &WinnerSaysHandler{repo: repo}
}

func (h *WinnerSaysHandler) CreateWinnerSays(w http.ResponseWriter, r *http.Request) {
    var winner models.WinnerSays
    if err := json.NewDecoder(r.Body).Decode(&winner); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.repo.Create(&winner); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(winner)
}

func (h *WinnerSaysHandler) GetAllWinnerSays(w http.ResponseWriter, r *http.Request) {
    winners, err := h.repo.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(winners)
}

func (h *WinnerSaysHandler) GetWinnerSaysByID(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    winner, err := h.repo.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(winner)
}

func (h *WinnerSaysHandler) DeleteWinnerSays(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if err := h.repo.Delete(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
