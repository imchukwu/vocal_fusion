package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vocal_fusion/internals/models"
	"vocal_fusion/internals/repository"
	"vocal_fusion/pkg/email"

	"github.com/go-chi/chi/v5"
)

type MessageHandler struct {
	Repo  repository.MessageRepository
	Email email.EmailService
}

func NewMessageHandler(repo repository.MessageRepository, email email.EmailService) *MessageHandler {
	return &MessageHandler{Repo: repo, Email: email}
}

// CreateMessage handles POST /messages
func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if msg.Content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	if msg.ReplyToID == nil && (msg.SenderName == "" || msg.Email == "") {
		http.Error(w, "Sender name and email are required for new messages", http.StatusBadRequest)
		return
	}

	// Handle defaults for replies from Admin
	if msg.ReplyToID != nil {
		if msg.SenderName == "" {
			msg.SenderName = "Vocal Fusion Admin"
		}
		if msg.Email == "" {
			msg.Email = "admin@vocalfusion.com"
		}
	}

	// Handle Reply Logic
	if msg.ReplyToID != nil {
		parentID := *msg.ReplyToID

		// 1. Check if parent message exists
		parent, err := h.Repo.GetMessageByID(parentID)
		if err != nil {
			http.Error(w, "Parent message not found", http.StatusNotFound)
			return
		}

		// 2. Enforce One-to-One Reply
		alreadyReplied, err := h.Repo.HasReply(parentID)
		if err != nil {
			http.Error(w, "Error checking reply status", http.StatusInternalServerError)
			return
		}
		if alreadyReplied {
			http.Error(w, "This message has already been replied to", http.StatusBadRequest)
			return
		}

		// Let's enforce that a reply has a subject based on parent if missed
		if msg.Subject == "" {
			msg.Subject = "RE: " + parent.Subject
		}

		// 3. Send Email Notification
		subject := "RE: " + parent.Subject
		if err := h.Email.SendEmail(parent.Email, subject, msg.Content); err != nil {
			// We log the error but optionally allow the message to be saved
			// Or we can fail the request. Let's fail it for consistency if email is critical.
			http.Error(w, "Failed to send email reply: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := h.Repo.CreateMessage(&msg); err != nil {
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}

	// Update Parent Status to "replied"
	if msg.ReplyToID != nil {
		if err := h.Repo.UpdateMessageStatus(*msg.ReplyToID, models.MessageStatusReplied); err != nil {
			// We don't fail here because the message is already saved and email sent,
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Message sent successfully",
	})
}

// GetAllMessages handles GET /messages
func (h *MessageHandler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Repo.GetAllMessages()
	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}

// GetMessageByID handles GET /messages/{id}
func (h *MessageHandler) GetMessageByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	msg, err := h.Repo.GetMessageByID(id)
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(msg)
}

// UpdateMessageStatus handles PATCH /messages/{id}/status
func (h *MessageHandler) UpdateMessageStatus(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var payload struct {
		Status models.MessageStatus `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !payload.Status.IsValid() {
		http.Error(w, "Invalid status. Allowed: unread, read, replied", http.StatusBadRequest)
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

// DeleteMessage handles DELETE /messages/{id}
func (h *MessageHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Repo.DeleteMessage(id); err != nil {
		http.Error(w, "Failed to delete message", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
