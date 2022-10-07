package logs

import "strings"

func Application(log string) string {
	for _, r := range log {
		switch r {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

func Replace(log string, old, new rune) string {
	return strings.ReplaceAll(log, string(old), string(new))
}

func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}