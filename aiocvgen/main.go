package main

import (
	"aiocvgen/internal/config"
	"aiocvgen/internal/ollama"
	"aiocvgen/internal/webserver"
	"embed"
	"fmt"
	"net/http"
	"strings"
)

// --- Embed Static Files ---
//
//go:embed web/*
var embedFs embed.FS

func main() {
	config.LoadConfig()
	ollama.SetupClient()
	fmt.Printf("Available models: \n- %s\n\n", strings.Join(ollama.GetModels(), "\n- "))
	root := http.Dir("web")
	// root := http.FS(embedFs)
	webserver.Run(root)
}
