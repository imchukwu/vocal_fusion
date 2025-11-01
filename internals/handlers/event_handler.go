package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"vocal_fusion/internals/models"
	"vocal_fusion/internals/repository"
)

type EventHandler struct {
	EventRepo repository.EventRepository
}

func NewEventHandler(eventRepo repository.EventRepository) *EventHandler {
	return &EventHandler{EventRepo: eventRepo}
}

// Create Event
func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.EventRepo.CreateEvent(&event); err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

// Fetch all events
func (h *EventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.EventRepo.GetAllEvents()
	if err != nil {
		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}

// Fetch single event by ID
func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	event, err := h.EventRepo.GetEventByID(uint(id))
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(event)
}

// ✅ Update Event
func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var updatedEvent models.Event
	if err := json.NewDecoder(r.Body).Decode(&updatedEvent); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event, err := h.EventRepo.GetEventByID(uint(id))
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	// Update fields
	event.Title = updatedEvent.Title
	event.Type = updatedEvent.Type
	event.Date = updatedEvent.Date
	event.Time = updatedEvent.Time
	event.Location = updatedEvent.Location

	if err := h.EventRepo.UpdateEvent(event); err != nil {
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(event)
}

// ✅ Delete Event
func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if err := h.EventRepo.DeleteEvent(uint(id)); err != nil {
		http.Error(w, "Failed to delete event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
