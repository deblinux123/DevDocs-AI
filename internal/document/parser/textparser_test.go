package parser

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
)

func TestTextParser(t *testing.T) {
	dir := t.TempDir()

	path := filepath.Join(dir, "text.txt")

	content := []byte(`# hello this is the text data for text.txt file`)

	err := os.WriteFile(path, content, 0644)

	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	textParser := &TextParser{}

	doc, err := textParser.Parse(
		context.Background(),
		path,
	)

	if err != nil {
		t.Fatalf("Parse() return error: %v", err)
	}

	if doc == nil {
		t.Fatalf("expected document, got nil")
	}

	if doc.Name != "text.txt" {
		t.Fatalf("expected name text.txt, got %s", doc.Name)
	}

	if doc.Content != string(content) {
		t.Errorf("expected content %q, got %q", string(content), doc.Content)
	}

	if doc.Type != ".txt" {
		t.Errorf("expected .txt, got %s", doc.Type)
	}

	if _, err := uuid.Parse(doc.ID); err != nil {
		t.Errorf("invlaid document UUID: %v", err)
	}
}

func TestTextParser_UnsupportedFileType(t *testing.T) {
	dir := t.TempDir()

	path := filepath.Join(dir, "text.md")

	content := []byte(`test unsupported file format`)

	err := os.WriteFile(path, content, 0644)

	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	textParser := &TextParser{}

	doc, err := textParser.Parse(
		context.Background(),
		path,
	)

	if err == nil {
		t.Fatalf("expected error, got nil for unsupported file type.")
	}

	if doc != nil {
		t.Fatalf("expected nil docmuent, got %+v", doc)
	}
}

func TestTextParser_ContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	cancel()

	textParser := &TextParser{}

	doc, err := textParser.Parse(
		ctx,
		"test.txt",
	)

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled, got %v", err)
	}

	if doc != nil {
		t.Fatalf("expected nil doc, %+v", doc)
	}
}
