package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"vocal_fusion/internals/models"
	"vocal_fusion/internals/repository"

	"github.com/go-chi/chi/v5"
)

type SchoolHandler struct {
	Repo repository.SchoolRepository
}

func NewSchoolHandler(repo repository.SchoolRepository) *SchoolHandler {
	return &SchoolHandler{Repo: repo}
}

// CreateSchool handles POST /schools
func (h *SchoolHandler) CreateSchool(w http.ResponseWriter, r *http.Request) {
	var school models.School
	if err := json.NewDecoder(r.Body).Decode(&school); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if school.Name == "" {
		http.Error(w, "School name is required", http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateSchool(&school); err != nil {
		http.Error(w, "Failed to create school", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "School created successfully",
	})
}

// GetAllSchools handles GET /schools
func (h *SchoolHandler) GetAllSchools(w http.ResponseWriter, r *http.Request) {
	schools, err := h.Repo.GetAllSchools()
	if err != nil {
		http.Error(w, "Failed to fetch schools", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(schools)
}

// GetSchoolByID handles GET /schools/{id}
func (h *SchoolHandler) GetSchoolByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	school, err := h.Repo.GetSchoolByID(id)
	if err != nil {
		http.Error(w, "School not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(school)
}

// UpdateSchool handles PUT /schools/{id}
func (h *SchoolHandler) UpdateSchool(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var updated models.School
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	school, err := h.Repo.GetSchoolByID(id)
	if err != nil {
		http.Error(w, "School not found", http.StatusNotFound)
		return
	}

	school.Name = updated.Name
	school.Email = updated.Email
	school.Address = updated.Address
	school.State = updated.State
	school.City = updated.City
	school.PrincipalName = updated.PrincipalName
	school.CoordinationName = updated.CoordinationName
	school.PaymentStatus = updated.PaymentStatus
	school.MediaList = updated.MediaList
	school.ConfirmStatus = updated.ConfirmStatus

	if err := h.Repo.UpdateSchool(school); err != nil {
		http.Error(w, "Failed to update school", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "School updated successfully",
	})
}

// DeleteSchool handles DELETE /schools/{id}
func (h *SchoolHandler) DeleteSchool(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Repo.DeleteSchool(id); err != nil {
		http.Error(w, "Failed to delete school", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
