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
	mux.HandleFunc("/api/getRuntimeConfig", apiGetRuntimeConfig)
	mux.HandleFunc("/api/sendConfig", apiSendConfig)
	mux.HandleFunc("/api/getCurrentCV", apiGetCurrentCV)
	mux.HandleFunc("/api/setCurrentCV", apiSetCurrentCV)
	mux.HandleFunc("/api/saveCurrentCV", apiSaveCurrentCV)
	mux.HandleFunc("/api/getGeneratedCVs", apiGetGeneratedCVs)
	mux.HandleFunc("/api/runGeneration", apiRunGeneration)
	mux.HandleFunc("/api/cancelGeneration", apiCancelGeneration)
	mux.HandleFunc("/api/generationStatus", apiGenerationStatus)
	mux.Handle("/", http.FileServer(staticFiles))
	generatorConfig = config.Get()

	srv := &http.Server{
		Addr:         config.Get().ListenAddr + ":" + strconv.Itoa(config.Get().ListenPort),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server on http://%s...", srv.Addr)
	return srv.ListenAndServe()
}
