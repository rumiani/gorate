package lang

import (
	"encoding/json"
	"fmt"
	"os"
)

var translations = make(map[string]map[string]string)

func LoadTranslations() error {
	langs := []string{"en", "fa"}
	for _, lang := range langs {
		data, err := os.ReadFile(fmt.Sprintf("locales/%s.json", lang))
		if err != nil {
			return fmt.Errorf("failed to read %s.json: %w", lang, err)
		}
		var content map[string]string
		if err := json.Unmarshal(data, &content); err != nil {
			return fmt.Errorf("failed to parse %s.json: %w", lang, err)
		}
		translations[lang] = content
	}
	return nil
}

// T returns the translated text for a key and language
func T(lang, key string) string {
	if tr, ok := translations[lang][key]; ok {
		return tr
	}
	// fallback to English if missing
	if tr, ok := translations["en"][key]; ok {
		return tr
	}
	return key
}
