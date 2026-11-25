package parser

import (
	"strings"
	"testing"
)

func TestNewParser_Markdown(t *testing.T) {
	path := "test.md"

	parser, err := NewParser(path)

	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	_, ok := parser.(*MarkdownParser)

	if !ok {
		t.Fatalf("expected markdown parser, got %v", parser)
	}
}

func TestNewParser_Text(t *testing.T) {
	path := "test.txt"

	parser, err := NewParser(path)

	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	_, ok := parser.(*TextParser)

	if !ok {
		t.Fatalf("expected text parser, got %v", parser)
	}
}

func TestNewParser_Unsupported(t *testing.T) {
	path := "test.jpg"

	parser, err := NewParser(path)

	if !strings.Contains(err.Error(), ".jpg") {
		t.Fatalf("expected error to contain .jpg, got %v", err)
	}

	if err == nil {
		t.Fatalf("expected err, got nil")
	}

	if parser != nil {
		t.Fatalf("expected nil parser, gor %T", parser)
	}

}
