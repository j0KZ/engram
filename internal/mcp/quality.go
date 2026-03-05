package mcp

import (
	"fmt"
	"strings"
)

const (
	minTitleChars   = 3
	minContentWords = 5
)

var noiseSet = map[string]bool{
	"ok": true, "done": true, "yes": true, "no": true,
	"test": true, "todo": true, "fixme": true,
	"placeholder": true, "n/a": true, "none": true,
	"null": true, "empty": true, "tbd": true,
}

func checkObservationQuality(title, content string) error {
	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)

	if len([]rune(title)) < minTitleChars {
		return fmt.Errorf("title too short (min %d chars): use a descriptive, searchable title", minTitleChars)
	}

	words := strings.Fields(content)
	if len(words) < minContentWords {
		return fmt.Errorf("content too brief (%d words, min %d): provide context — what, why, and where", len(words), minContentWords)
	}

	if noiseSet[strings.ToLower(content)] {
		return fmt.Errorf("content is a placeholder or noise value: provide meaningful information")
	}

	return nil
}
