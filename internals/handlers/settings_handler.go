package handlers

import (
	"encoding/json"
	"net/http"
	"vocal_fusion/internals/models"
	"vocal_fusion/internals/repository"
)

type SettingsHandler struct {
	Repo repository.SettingsRepository
}

func NewSettingsHandler(repo repository.SettingsRepository) *SettingsHandler {
	return &SettingsHandler{Repo: repo}
}

// GET /settings
func (h *SettingsHandler) GetSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := h.Repo.GetSettings()
	if err != nil {
		http.Error(w, "Failed to fetch settings", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(settings)
}

// PUT /settings
func (h *SettingsHandler) UpdateSettings(w http.ResponseWriter, r *http.Request) {
	var settings models.Settings
	if err := json.NewDecoder(r.Body).Decode(&settings); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Repo.UpdateSettings(&settings); err != nil {
		http.Error(w, "Failed to update settings", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Settings updated successfully",
	})
}
