package config

import (
	"encoding/json"
	"os"
)

type Conf struct {
	ListenAddr        string  `json:"listen_addr"`
	ListenPort        int     `json:"listen_port"`
	OllamaUrl         string  `json:"ollama_url"`
	OllamaModel       string  `json:"ollama_model"`
	OllamaThink       bool    `json:"ollama_think"`
	OllamaTemperature float64 `json:"ollama_temperature"`
	OllamaStream      bool    `json:"ollama_stream"`
	Batch             int     `json:"batch"`
}

var Defaults = Conf{
	ListenAddr:        "0.0.0.0",
	ListenPort:        54321,
	OllamaUrl:         "127.0.0.1:11434",
	OllamaModel:       "gpt-oss:20b",
	OllamaThink:       false,
	OllamaTemperature: 0.5,
	OllamaStream:      true,
	Batch:             3,
}

const File = "config.json"

func LoadConfig() {
	contents, err := os.ReadFile(File)
	if err == nil {
		json.Unmarshal(contents, &Defaults)
	}
	SaveConfig()
}
func SaveConfig() error {
	j, err := json.MarshalIndent(Defaults, "", "\t")
	os.WriteFile(File, j, 0640)
	return err
}

func Get() Conf {
	return Defaults
}
