package model

import "time"

// Kind represents the type of a Crumb, either a File or a Directory.
type Kind int

// Idiomatic Enumeration of possible Kind values.
const (
	File Kind = iota
	Directory
)

// Determines what kind of location a Crumb points to, either a File or a Directory
type Location struct {
	Kind Kind
	Path string
}

type Crumb struct {
	Name        string
	CreatedBy   string
	Comment     string
	Link        Location
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}
