package webserver

import (
	"encoding/json"
	"net/http"
)

type CV struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Html      string `json:"html"`
	Generated string `json:"generated"`
	Modified  string `json:"modified"`
	Selected  bool   `json:"selected"`
}
type idReq struct {
	Id int `json:"id"`
}

var CurrentCV CV
var GeneratedCVs []CV

func apiGetGeneratedCVs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GeneratedCVs)
}

func apiSetCurrentCV(w http.ResponseWriter, r *http.Request) {
	var msg idReq
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	w.Write([]byte("OK"))
}

func apiGetCurrentCV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GeneratedCVs)
}

func apiSaveCurrentCV(w http.ResponseWriter, r *http.Request) {
	var msg CV
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	w.Write([]byte("OK"))
}
