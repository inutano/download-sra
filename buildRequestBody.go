package main

import (
	"bytes"
	"encoding/json"
	"log"
)

func BuildRequestBody(keyword, date, reads, bases string) (buf bytes.Buffer) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "Ohanami",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return
}
