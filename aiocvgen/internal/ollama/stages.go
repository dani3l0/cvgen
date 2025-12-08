package ollama

type typeStages struct {
	Idle     string
	Loading  string
	Thinking string
	Writing  string
	Done     string
}

var Stage string
var Stages = typeStages{
	Idle:     "idle",
	Loading:  "loading",
	Thinking: "thinking",
	Writing:  "writing",
	Done:     "done",
}
