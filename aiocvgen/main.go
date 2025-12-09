package main

import (
	"aiocvgen/internal/config"
	"aiocvgen/internal/ollama"
	"aiocvgen/internal/webserver"
	"embed"
)

// --- Embed Static Files ---
//
//go:embed web/*
var embedFs embed.FS

func main() {
	config.LoadConfig()
	ollama.SetupClient()
	webserver.Run(embedFs, "web")
}
