package parser

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
)

func TestMarkdonwParser(t *testing.T) {
	dir := t.TempDir()

	t.Log("Temporary directory created.")

	path := filepath.Join(dir, "test.md")

	content := []byte(`# Hello gophers`)

	err := os.WriteFile(path, content, 0644)

	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	markdown := &MarkdownParser{}

	doc, err := markdown.Parse(context.Background(), path)

	if err != nil {
		t.Fatalf("Parse() return error: %v", err)
	}

	if doc == nil {
		t.Fatalf("expected document, got nil")
	}

	if doc.Name != "test.md" {
		t.Errorf("expected name test.md, got %s", doc.Name)
	}

	if doc.Content != string(content) {
		t.Errorf("expected content %q, got %q", string(content), doc.Content)
	}

	if doc.Type != ".md" {
		t.Errorf("expected type .md, got %s", doc.Type)
	}

	if _, err := uuid.Parse(doc.ID); err != nil {
		t.Errorf("invalid document UUID: %v", err)
	}
}

func TestMarkdonwParser_unsupportedFileType(t *testing.T) {
	dir := t.TempDir()

	path := filepath.Join(dir, "test.txt")

	content := []byte(`# hello this is the text file for testing`)

	err := os.WriteFile(path, content, 0644)

	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	markdwon := &MarkdownParser{}

	doc, err := markdwon.Parse(
		context.Background(),
		path,
	)

	if err == nil {
		t.Fatalf("expected error for unsupported file type")
	}

	if doc != nil {
		t.Fatalf("expected nil document, got %+v", doc)
	}
}

func TestMarkdownParser_ContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	cancel()

	markdown := &MarkdownParser{}

	doc, err := markdown.Parse(
		ctx,
		"test.md",
	)

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}

	if doc != nil {
		t.Fatalf("expected nil doc, got %+v", doc)
	}
}
