package handlers

import (
	"backend/internal/models"
	"backend/internal/repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

func (h *CabinHandler) DeleteCabin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid cabin ID", http.StatusBadRequest)
		return
	}

	if err := h.Repo.DeleteCabin(id); err != nil {
		http.Error(w, "Failed to delete cabin ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CabinHandler) CreateCabin(w http.ResponseWriter, r *http.Request) {
	var cabin models.Cabin
	if err := json.NewDecoder(r.Body).Decode(&cabin); err != nil {
		log.Print(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateCabin(&cabin); err != nil {
		log.Print(err)
		http.Error(w, "Failed to create cabin", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cabin)
}
