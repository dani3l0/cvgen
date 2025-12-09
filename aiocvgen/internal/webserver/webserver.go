package webserver

import (
	"aiocvgen/internal/config"
	"embed"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var mux *http.ServeMux

// --- Main Webserver Runner ---

func Run(fs embed.FS, staticDir string) error {
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
	//	mux.Handle("/", http.FileServer(http.Dir(staticDir))) // For front development!
	mux.Handle("/", rootPath(http.FileServer(http.FS(fs)), staticDir)) // Front embedded into binary!
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

func rootPath(h http.Handler, staticDir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// add header(s)
		w.Header().Set("Cache-Control", "no-cache")

		if r.URL.Path == "/" {
			r.URL.Path = fmt.Sprintf("/%s/", staticDir)
		} else {
			b := strings.Split(r.URL.Path, "/")[0]
			if b != staticDir {
				r.URL.Path = fmt.Sprintf("/%s%s", staticDir, r.URL.Path)
			}
		}
		h.ServeHTTP(w, r)
	})
}
