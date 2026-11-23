package parser

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/deblinux123/DevDocs-AI/internal/document"
)

type TextParser struct{}

var _ Parser = (*TextParser)(nil)

func (p *TextParser) Parse(
	ctx context.Context,
	path string,
) (*document.Document, error) {
	// check context cancellation
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if strings.ToLower(filepath.Ext(path)) != ".txt" {
		return nil, fmt.Errorf("unsupported file type, expected .txt")
	}

	// open the file
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	// read the content
	content, err := io.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	name := filepath.Base(path)
	fileType := strings.ToLower(filepath.Ext(path))

	doc := &document.Document{
		ID:        document.NewID(),
		Name:      name,
		Path:      path,
		Type:      fileType,
		Content:   string(content),
		CreatedAt: time.Now(),
	}

	// return doc

	return doc, nil
}
