package webserver

import (
	"aiocvgen/internal/config"
	"encoding/json"
	"net/http"
)

func apiGetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.Get())
}

func apiGetRuntimeConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(generatorConfig)
}

func apiSendConfig(w http.ResponseWriter, r *http.Request) {
	var msg config.Conf
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	config.Defaults = msg
	config.SaveConfig()
	generatorConfig = config.Get()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
