package parser

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/deblinux123/DevDocs-AI/internal/document"
)

type TextParser struct{}

var _ Parser = (*TextParser)(nil)

func (p *TextParser) Parse(
	ctx context.Context,
	path string,
) (*document.Document, error) {
	if strings.ToLower(filepath.Ext(path)) != ".txt" {
		return nil, fmt.Errorf("unsupported file type, expected .txt")
	}

	// read the content
	content, err := readFile(ctx, path)

	if err != nil {
		return nil, err
	}

	return newDocument(path, content), nil
}
