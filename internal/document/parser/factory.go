package parser

import (
	"fmt"
	"path/filepath"
	"strings"
)

func NewParser(path string) (Parser, error) {
	fileType := strings.ToLower(filepath.Ext(path))

	var parser Parser
	switch fileType {
	case ".md":
		parser = &MarkdownParser{}
	case ".txt":
		parser = &TextParser{}
	default:
		return nil, fmt.Errorf("unsupported file type: %s", fileType)
	}

	return parser, nil
}
