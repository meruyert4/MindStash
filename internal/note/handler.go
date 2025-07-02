package note

import (
	"encoding/json"
	"mindstash/pkg/logger"
	"net/http"
)

type NoteHandler struct {
	service Service
	logger  logger.Logger
}

func NewHandler(s Service, log logger.Logger) *NoteHandler {
	return &NoteHandler{
		service: s,
		logger:  log,
	}
}

func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input CreateNoteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Error(ctx, "Invalid JSON Body", "error", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	created, err := h.service.CreateNote(ctx, input)
	if err != nil {
		h.logger.Error(ctx, err.Error(), "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info(ctx, "Note successfully created!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *NoteHandler) GetNoteById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	note, err := h.service.GetNoteById(ctx, id)
	if err != nil {
		h.logger.Error(ctx, err.Error(), "id=", id)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	h.logger.Info(ctx, "Note successfully returned!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notes, err := h.service.GetNotes(ctx)
	if err != nil {
		h.logger.Error(ctx, err.Error(), "error")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	h.logger.Info(ctx, "Notes successfully created!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	var input UpdateNoteInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Error(ctx, "Invalid JSON Body", "error", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	updatedNote, err := h.service.UpdateNote(ctx, id, input)
	if err != nil {
		h.logger.Error(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info(ctx, "Note successfully updated!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedNote)
}

func (h *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	err := h.service.DeleteNote(ctx, id)
	if err != nil {
		h.logger.Error(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.logger.Info(ctx, "Note successfully deleted!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
