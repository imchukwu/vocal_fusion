package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"vocal_fusion/internals/models"
	"vocal_fusion/internals/repository"
)

type FAQHandler struct {
	Repo repository.FAQRepository
}

func NewFAQHandler(repo repository.FAQRepository) *FAQHandler {
	return &FAQHandler{Repo: repo}
}

func (h *FAQHandler) CreateFAQ(w http.ResponseWriter, r *http.Request) {
	var faq models.FAQ
	if err := json.NewDecoder(r.Body).Decode(&faq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if faq.Subject == "" {
		http.Error(w, "Subject is required", http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateFAQ(&faq); err != nil {
		http.Error(w, "Failed to create FAQ", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(faq)
}

func (h *FAQHandler) GetAllFAQs(w http.ResponseWriter, r *http.Request) {
	faqs, err := h.Repo.GetAllFAQs()
	if err != nil {
		http.Error(w, "Failed to fetch FAQs", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(faqs)
}

func (h *FAQHandler) GetFAQByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	faq, err := h.Repo.GetFAQByID(id)
	if err != nil {
		http.Error(w, "FAQ not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(faq)
}

func (h *FAQHandler) UpdateFAQ(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var faq models.FAQ
	if err := json.NewDecoder(r.Body).Decode(&faq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	faq.ID = id
	if err := h.Repo.UpdateFAQ(&faq); err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(faq)
}

func (h *FAQHandler) DeleteFAQ(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Repo.DeleteFAQ(id); err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
