package handlers

import (
	"encoding/json"
	"net/http"

	"backend/internal/models"
	"backend/internal/repository" // update module path
)

// SettingsHandler handles HTTP requests for settings.
type SettingsHandler struct {
	Repo *repository.Repository
}

// GetSettings handles GET /api/settings.
func (h *SettingsHandler) GetSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := h.Repo.GetSettings()
	if err != nil {
		http.Error(w, "Failed to fetch settings", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settings)
}

func (h *SettingsHandler) UpdateSettings(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON payload.
	var newSettings models.Settings
	if err := json.NewDecoder(r.Body).Decode(&newSettings); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	currentSettings, err := h.Repo.GetSettings()
	if err != nil {
		http.Error(w, "Settings not found", http.StatusNotFound)
		return
	}

	newSettings.ID = currentSettings.ID

	if err := h.Repo.UpdateSettings(&newSettings); err != nil {
		http.Error(w, "Failed to update settings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSettings)
}
