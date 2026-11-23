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

type MarkdownParser struct{}

func (p *MarkdownParser) Parse(
	ctx context.Context,
	path string,
) (*document.Document, error) {
	// 1. cheack the context cancellation
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:

	}

	if strings.ToLower(filepath.Ext(path)) != ".md" {
		return nil, fmt.Errorf("unsupported file type: expected .md")
	}

	// 2. open file
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("failed to open the file: %w", err)
	}

	// 3. close file with defer
	defer file.Close()

	// 4. read file content
	content, err := io.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("failed to read the file: %w", err)
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
