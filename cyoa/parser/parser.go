package parser

import "encoding/json"

// Story type is a map of Chapters
type Story map[string]Chapter

// Chapter type is a struct that represents a chapter of our story
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// JSONParser parses a JSON stream and returns a Story
func JSONParser(j []byte) (Story, error) {
	var ret Story

	if err := json.Unmarshal(j, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
