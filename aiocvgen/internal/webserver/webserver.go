package webserver

import (
	"aiocvgen/internal/config"
	"log"
	"net/http"
	"strconv"
	"time"
)

var mux *http.ServeMux

// --- Main Webserver Runner ---

func Run(staticFiles http.FileSystem) error {
	mux = http.NewServeMux()

	mux.HandleFunc("/api/getConfig", apiGetConfig)
	mux.HandleFunc("/api/postConfig", apiSendConfig)
	mux.HandleFunc("/api/getCurrentCV", apiGetCurrentCV)
	mux.HandleFunc("/api/getGeneratedCVs", apiGetGeneratedCVs)
	mux.Handle("/", http.FileServer(staticFiles))

	srv := &http.Server{
		Addr:         config.Get().ListenAddr + ":" + strconv.Itoa(config.Get().ListenPort),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server on http://%s...", srv.Addr)
	return srv.ListenAndServe()
}
