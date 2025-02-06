package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/repository"

	"github.com/gorilla/mux"
)

type CabinHandler struct {
	Repo *repository.Repository
}

func (h *CabinHandler) GetAllCabins(w http.ResponseWriter, r *http.Request) {
	cabins, err := h.Repo.GetAllCabins()
	if err != nil {
		http.Error(w, "Failed to fetch cabins", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cabins)
}

func (h *CabinHandler) GetCabinByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid cabin ID", http.StatusBadRequest)
		return
	}
	cabin, err := h.Repo.GetCabinByID(id)
	if err != nil {
		http.Error(w, "Cabin not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cabin)
}
