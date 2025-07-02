package note

import (
	"errors"
	"strings"
)

func validateNoteInput(title, content string) error {
	if title == "" {
		return errors.New("title is required")
	}
	if len(title) > 255 {
		return errors.New("title is too long")
	}
	if content == "" {
		return errors.New("content is required")
	}
	return nil
}

func sanitizeNoteInput(title, content string) (string, string) { //delete unnessary spaces
	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)

	if len(title) > 200 {
		title = title[:200]
	}
	if len(content) > 10000 {
		content = content[:10000]
	}
	return title, content
}

func isEmptyUpdate(input UpdateNoteInput) bool {
	return input.Title == nil && input.Content == nil
}
