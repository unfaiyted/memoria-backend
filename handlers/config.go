package handlers

import (
	"encoding/json"
	"memoria-backend/models"
	"memoria-backend/services"
	"net/http"
)

// ConfigHandler handles configuration API endpoints
type ConfigHandler struct {
	configService services.ConfigService
}

// NewConfigHandler creates a new configuration handler
func NewConfigHandler(configService services.ConfigService) *ConfigHandler {
	return &ConfigHandler{
		configService: configService,
	}
}

// GetConfig returns the current configuration
func (h *ConfigHandler) GetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.configService.GetConfig())
}

// UpdateConfig handles configuration updates
func (h *ConfigHandler) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	var cfg models.Configuration
	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.configService.SaveConfig(cfg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
