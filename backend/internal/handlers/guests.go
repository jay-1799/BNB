package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/repository"
	"github.com/gorilla/mux"
)

type GuestHandler struct {
	Repo *repository.Repository
}

func (h *GuestHandler) GetAllGuests(w http.ResponseWriter, r *http.Request) {
	guests, err := h.Repo.GetAllGuests()
	if err != nil {
		http.Error(w, "Failed to fetch guests", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guests)
}

func (h *GuestHandler) GetGuestByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid guest ID", http.StatusBadRequest)
		return
	}
	guest, err := h.Repo.GetGuestByID(id)
	if err != nil {
		http.Error(w, "Guest not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guest)
}
