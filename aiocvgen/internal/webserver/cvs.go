package webserver

import (
	"encoding/json"
	"net/http"
	"time"
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
	for i := range len(GeneratedCVs) {
		sel := i == msg.Id
		GeneratedCVs[i].Selected = sel
		if sel {
			CurrentCV = GeneratedCVs[i]
		}
	}
	w.Write([]byte("OK"))
}

func apiGetCurrentCV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CurrentCV)
}

func apiSaveCurrentCV(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&CurrentCV); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	CurrentCV.Modified = time.Now().Format("2006-01-02 15:04:05")
	w.Write([]byte("OK"))
}
