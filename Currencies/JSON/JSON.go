package JSON

import (
	"encoding/json"
	"os"
)

type JSON struct {
	Currency string `json:"currency"`
	IN       string `json:"IN"`
	OUT      string `json:"OUT"`
}

func Save(things []JSON) {
	data, _ := json.MarshalIndent(things, "", "  ")
	os.WriteFile("history.json", data, 0644)
}

func Load() []JSON {
	here, _ := os.ReadFile("history.json")
	var history []JSON
	json.Unmarshal(here, &history)
	return history
}
