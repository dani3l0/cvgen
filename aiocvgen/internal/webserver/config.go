package webserver

import (
	"aiocvgen/internal/config"
	"encoding/json"
	"net/http"
)

func apiGetConfig(w http.ResponseWriter, r *http.Request) {
	resp := config.Get()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func apiSendConfig(w http.ResponseWriter, r *http.Request) {
	var msg config.Conf
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	config.Defaults = msg
	config.SaveConfig()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
