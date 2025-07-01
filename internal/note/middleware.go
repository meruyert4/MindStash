package note

import "errors"

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
