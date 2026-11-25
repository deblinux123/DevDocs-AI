package parser

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/deblinux123/DevDocs-AI/internal/document"
)

type Parser interface {
	Parse(
		ctx context.Context,
		path string,
	) (*document.Document, error)
}

func readFile(ctx context.Context, path string) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return content, nil
}
