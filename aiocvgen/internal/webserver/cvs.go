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
}

var CurrentCV CV
var GeneratedCVs []CV

func apiGetGeneratedCVs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GeneratedCVs)
}

func apiGetCurrentCV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GeneratedCVs)
}
