package parser

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/deblinux123/DevDocs-AI/internal/document"
)

func newDocument(path string, content []byte) *document.Document {
	name := filepath.Base(path)

	fileType := strings.ToLower(filepath.Ext(path))

	doc := &document.Document{
		ID:        document.NewID(),
		Name:      name,
		Type:      fileType,
		Path:      path,
		Content:   string(content),
		CreatedAt: time.Now(),
	}

	return doc
}
