package parser

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/deblinux123/DevDocs-AI/internal/document"
)

type MarkdownParser struct{}

var _ Parser = (*MarkdownParser)(nil)

func (p *MarkdownParser) Parse(
	ctx context.Context,
	path string,
) (*document.Document, error) {

	if strings.ToLower(filepath.Ext(path)) != ".md" {
		return nil, fmt.Errorf("unsupported file type: expected .md")
	}

	// 4. read file content
	content, err := readFile(ctx, path)

	if err != nil {
		return nil, err
	}

	return newDocument(path, content), nil
}
