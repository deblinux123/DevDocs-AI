package parser

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

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

	// 5. extracte file name
	name := filepath.Base(path)
	fileType := strings.ToLower(filepath.Ext(path))

	// 6. create document
	doc := &document.Document{
		ID:        document.NewID(),
		Name:      name,
		Path:      path,
		Type:      fileType,
		Content:   string(content),
		CreatedAt: time.Now(),
	}

	// 7. return document
	return doc, nil
}
