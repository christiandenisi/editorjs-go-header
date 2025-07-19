package header

import (
	"fmt"
	"html"

	"github.com/christiandenisi/editorjs-go"
)

type HeaderData struct {
	Text  string `json:"text"`
	Level int    `json:"level"`
}

// RenderHeader is the renderer function for the "header" block.
func RenderHeader(b editorjs.Block[HeaderData], ctx *editorjs.Context) (string, error) {
	level := b.Data.Level
	if level < 1 || level > 6 {
		return "", fmt.Errorf("invalid level: %d", level)
	}
	escapedText := html.EscapeString(b.Data.Text)
	return fmt.Sprintf("<h%d>%s</h%d>", b.Data.Level, escapedText, b.Data.Level), nil
}
