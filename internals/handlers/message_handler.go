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

type MessageHandler struct {
	Repo repository.MessageRepository
}

func NewMessageHandler(repo repository.MessageRepository) *MessageHandler {
	return &MessageHandler{Repo: repo}
}

// CreateMessage handles POST /api/messages
func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if msg.SenderName == "" || msg.Email == "" || msg.Content == "" {
		http.Error(w, "Sender name, email and content are required", http.StatusBadRequest)
		return
	}

	msg.Date = time.Now()

	if err := h.Repo.CreateMessage(&msg); err != nil {
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Message sent successfully",
	})
}

// GetAllMessages handles GET /api/messages
func (h *MessageHandler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Repo.GetAllMessages()
	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}

// GetMessageByID handles GET /api/messages/{id}
func (h *MessageHandler) GetMessageByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	msg, err := h.Repo.GetMessageByID(id)
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(msg)
}

// UpdateMessageStatus handles PATCH /api/messages/{id}/status
func (h *MessageHandler) UpdateMessageStatus(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var payload struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if payload.Status == "" {
		http.Error(w, "Status is required", http.StatusBadRequest)
		return
	}

	if err := h.Repo.UpdateMessageStatus(id, payload.Status); err != nil {
		http.Error(w, "Failed to update status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Status updated successfully",
	})
}

// DeleteMessage handles DELETE /api/messages/{id}
func (h *MessageHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Repo.DeleteMessage(id); err != nil {
		http.Error(w, "Failed to delete message", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
