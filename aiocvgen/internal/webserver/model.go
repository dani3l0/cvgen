package webserver

import (
	"aiocvgen/internal/config"
	"aiocvgen/internal/ollama"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type generationStatus struct {
	Stage       string `json:"stage"`
	StagesDone  int    `json:"stages_done"`
	StagesTotal int    `json:"stages_total"`
	BatchDone   int    `json:"batch_done"`
	BatchSize   int    `json:"batch_size"`
	Running     bool   `json:"running"`
}

var running bool
var cancelling bool
var generatorConfig config.Conf

func runGeneration() {
	if running {
		return
	}
	running = true
	GeneratedCVs = []CV{}
	for range generatorConfig.Batch {
		if cancelling {
			break
		}
		prompt := "# Candidate information:\n"
		prompt += generatorConfig.LlmAboutMe + "\n-----\n"
		prompt += "# Job offer:\n"
		prompt += generatorConfig.LlmJobOffer + "\n-----\n"
		prompt += "# Target HTML format and look:\n"
		prompt += generatorConfig.LlmTemplateOptions + "\n-----\n"
		prompt += "# Important notes:\n"
		prompt += generatorConfig.LlmOtherNotes + "\n\n"
		prompt = strings.Trim(prompt, "\n")
		resp := ollama.GenerateResume(prompt, generatorConfig)
		html := strings.SplitN(resp, "<html", 2)

		if len(html) > 1 {
			html = strings.Split(html[1], "</html>")
			timex := time.Now().Format("2006-01-02 15:04:05")

			GeneratedCVs = append(GeneratedCVs, CV{
				Id:        len(GeneratedCVs),
				Name:      fmt.Sprintf("Resume #%d", len(GeneratedCVs)+1),
				Html:      fmt.Sprintf("<html%s</html>", html[0]),
				Generated: timex,
				Modified:  "never",
				Selected:  false,
			})
		}
	}
	cancelling = false
	time.Sleep(5 * time.Second)
	running = false
}

func apiRunGeneration(w http.ResponseWriter, r *http.Request) {
	if running {
		http.Error(w, "already running", http.StatusForbidden)
		return
	}
	var msg config.Conf
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		msg = config.Get()
	}
	generatorConfig = msg
	go runGeneration()
	w.Write([]byte("OK"))
}

func apiCancelGeneration(w http.ResponseWriter, r *http.Request) {
	if cancelling {
		http.Error(w, "already cancelling", http.StatusForbidden)
		return
	}
	if ollama.Stage == ollama.StageIdle {
		http.Error(w, "not running", http.StatusForbidden)
		return
	}
	cancelling = true
	w.Write([]byte("OK"))
}

func apiGenerationStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stagesDone, stagesAll := stags()
	json.NewEncoder(w).Encode(generationStatus{
		Stage:       ollama.Stage,
		StagesDone:  stagesDone,
		StagesTotal: stagesAll,
		BatchDone:   len(GeneratedCVs),
		BatchSize:   generatorConfig.Batch,
		Running:     running,
	})
}

func stags() (int, int) {
	finishedOnes := len(GeneratedCVs)
	multiplier := 1
	if generatorConfig.OllamaThink {
		multiplier = 2
		finishedOnes *= multiplier
		if ollama.Stage == ollama.StageWriting {
			finishedOnes += 1 // thinking is finished
		}
	}
	stagesTotal := generatorConfig.Batch * multiplier
	stagesTotal += 1 // additional stage for model loading
	if ollama.Stage != ollama.StageLoading {
		finishedOnes += 1 // model must been loaded, so first stage is done
	}
	return finishedOnes, stagesTotal
}
