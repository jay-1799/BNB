package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

type BookingHandler struct {
	Repo *repository.Repository
}

func (h *BookingHandler) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	after := r.URL.Query().Get("after")
	before := r.URL.Query().Get("before")

	bookings, err := h.Repo.GetBookings(after, before)
	if err != nil {
		http.Error(w, "Failed to fetch bookings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func (h *BookingHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}
	booking, err := h.Repo.GetBookingByID(id)
	if err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func (h *BookingHandler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	booking.ID = id // Ensure the ID is set

	// Call the repository update function (implement it if not already present)
	if err := h.Repo.UpdateBooking(&booking); err != nil {
		http.Error(w, "Failed to update booking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func (h *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	if err := h.Repo.DeleteBooking(id); err != nil {
		http.Error(w, "Failed to delete booking", http.StatusInternalServerError)
		return
	}

	// Return no content status if deletion was successful.
	w.WriteHeader(http.StatusNoContent)
}

func (h *BookingHandler) GetTodaysActivity(w http.ResponseWriter, r *http.Request) {
	today := time.Now().Format("2006-01-02") // format as YYYY-MM-DD
	// Example query: get bookings where start_date or end_date is today.
	bookings, err := h.Repo.GetBookingsForDate(today)
	if err != nil {
		http.Error(w, "Failed to fetch today's activity", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}
