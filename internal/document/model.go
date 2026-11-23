package document

import "time"

type Document struct {
	ID        string
	Name      string
	Path      string
	Type      string
	Content   string
	CreatedAt time.Time
}
