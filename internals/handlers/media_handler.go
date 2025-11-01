package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"vocal_fusion/internals/models"
	"vocal_fusion/internals/repository"
)

type MediaHandler struct {
	Repo repository.MediaRepository
}

func NewMediaHandler(repo repository.MediaRepository) *MediaHandler {
	return &MediaHandler{Repo: repo}
}

// CreateMedia handles POST /api/media
func (h *MediaHandler) CreateMedia(w http.ResponseWriter, r *http.Request) {
	var media models.Media
	if err := json.NewDecoder(r.Body).Decode(&media); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if media.Type == "" {
		http.Error(w, "Type is required", http.StatusBadRequest)
		return
	}

	// Set current date if not provided
	if media.Date.IsZero() {
		media.Date = time.Now()
	}

	if err := h.Repo.CreateMedia(&media); err != nil {
		http.Error(w, "Failed to create media", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(media)
}

// GetAllMedia handles GET /api/media
func (h *MediaHandler) GetAllMedia(w http.ResponseWriter, r *http.Request) {
	mediaList, err := h.Repo.GetAllMedia()
	if err != nil {
		http.Error(w, "Failed to fetch media", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(mediaList)
}

// GetMediaByID handles GET /api/media/{id}
func (h *MediaHandler) GetMediaByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	media, err := h.Repo.GetMediaByID(id)
	if err != nil {
		http.Error(w, "Media not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(media)
}

// UpdateMedia handles PUT /api/media/{id}
func (h *MediaHandler) UpdateMedia(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var media models.Media

	if err := json.NewDecoder(r.Body).Decode(&media); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	media.ID = id
	if err := h.Repo.UpdateMedia(&media); err != nil {
		http.Error(w, "Failed to update media", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(media)
}

// DeleteMedia handles DELETE /api/media/{id}
func (h *MediaHandler) DeleteMedia(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Repo.DeleteMedia(id); err != nil {
		http.Error(w, "Failed to delete media", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
