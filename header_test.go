package header

import (
	"testing"

	"github.com/christiandenisi/editorjs-go"
)

func TestRenderHeader_ValidLevel(t *testing.T) {
	block := editorjs.Block[HeaderData]{
		Data: HeaderData{
			Text:  "My Title",
			Level: 2,
		},
	}

	ctx := &editorjs.Context{}

	out, err := RenderHeader(block, ctx)
	if err != nil {
		t.Fatalf("RenderHeader returned an error: %v", err)
	}

	expected := "<h2>My Title</h2>"
	if out != expected {
		t.Errorf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestRenderHeader_Escaping(t *testing.T) {
	block := editorjs.Block[HeaderData]{
		Data: HeaderData{
			Text:  "Hello <world>",
			Level: 3,
		},
	}

	ctx := &editorjs.Context{}

	out, err := RenderHeader(block, ctx)
	if err != nil {
		t.Fatalf("RenderHeader returned an error: %v", err)
	}

	expected := "<h3>Hello &lt;world&gt;</h3>"
	if out != expected {
		t.Errorf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestRenderHeader_InvalidLevel(t *testing.T) {
	block := editorjs.Block[HeaderData]{
		Data: HeaderData{
			Text:  "Invalid Level",
			Level: 9,
		},
	}

	ctx := &editorjs.Context{}

	_, err := RenderHeader(block, ctx)
	if err == nil {
		t.Fatal("expected error for invalid level, got nil")
	}
}
