package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"vocal_fusion/internals/models"
	"vocal_fusion/internals/repository"
)

type SchoolEventHandler struct {
	Repo repository.SchoolEventRepository
}

func NewSchoolEventHandler(repo repository.SchoolEventRepository) *SchoolEventHandler {
	return &SchoolEventHandler{Repo: repo}
}

// POST /events/{eventID}/register
func (h *SchoolEventHandler) RegisterSchool(w http.ResponseWriter, r *http.Request) {
	eventID, _ := strconv.Atoi(chi.URLParam(r, "eventID"))

	var body struct {
		SchoolID int    `json:"school_id"`
		Code     string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	reg := &models.SchoolEvent{
		SchoolID: body.SchoolID,
		EventID:  eventID,
		Status:   "Registered",
		Code:     body.Code,
	}

	if err := h.Repo.RegisterSchoolForEvent(reg); err != nil {
		http.Error(w, "Failed to register school", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "School registered successfully",
	})
}

// GET /events/{eventID}/registrations
func (h *SchoolEventHandler) GetEventRegistrations(w http.ResponseWriter, r *http.Request) {
	eventID, _ := strconv.Atoi(chi.URLParam(r, "eventID"))

	regs, err := h.Repo.GetRegistrationsByEvent(eventID)
	if err != nil {
		http.Error(w, "Failed to fetch registrations", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(regs)
}

// GET /schools/{schoolID}/registrations
func (h *SchoolEventHandler) GetSchoolRegistrations(w http.ResponseWriter, r *http.Request) {
	schoolID, _ := strconv.Atoi(chi.URLParam(r, "schoolID"))

	regs, err := h.Repo.GetRegistrationsBySchool(schoolID)
	if err != nil {
		http.Error(w, "Failed to fetch registrations", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(regs)
}

// DELETE /events/{eventID}/unregister/{schoolID}
func (h *SchoolEventHandler) UnregisterSchool(w http.ResponseWriter, r *http.Request) {
	eventID, _ := strconv.Atoi(chi.URLParam(r, "eventID"))
	schoolID, _ := strconv.Atoi(chi.URLParam(r, "schoolID"))

	if err := h.Repo.UnregisterSchool(eventID, schoolID); err != nil {
		http.Error(w, "Failed to unregister school", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Unregistered successfully",
	})
}

// PUT /events/{eventID}/schools/{schoolID}/code
func (h *SchoolEventHandler) UpdateSchoolEventCode(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(chi.URLParam(r, "eventID"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}
	
	schoolID, err := strconv.Atoi(chi.URLParam(r, "schoolID"))
	if err != nil {
		http.Error(w, "Invalid school ID", http.StatusBadRequest)
		return
	}

	var payload struct {
		Code string `json:"code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Repo.UpdateSchoolEventCode(eventID, schoolID, payload.Code); err != nil {
		http.Error(w, "Failed to update code", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "School event code updated successfully",
	})
}
