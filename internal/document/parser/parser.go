package parser

import (
	"context"

	"github.com/deblinux123/DevDocs-AI/internal/document"
)

type Parser interface {
	Parse(
		ctx context.Context,
		path string,
	) (*document.Document, error)
}
