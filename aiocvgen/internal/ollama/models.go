package ollama

import (
	"context"
	"fmt"
)

// Gets list of available models
func GetModels() []string {
	var models []string
	list, err := client.List(context.Background())
	if err == nil {
		for _, entry := range list.Models {
			models = append(models, fmt.Sprintf("%s (%.2f GB)", entry.Model, float64(entry.Size)/1000/1000/1000))
		}
	}
	return models
}
