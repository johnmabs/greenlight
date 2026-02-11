package data

import (
	"time"
)

// Annotate the Movie struct with struct tags to control how the keys appear in the
// JSON-encoded output.
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Use the - directive
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitzero"`           // Add the omitzero directive
	Runtime   int32     `json:"runtime,omitzero,string"` // Add the string directive
	Genres    []string  `json:"genres,omitempty"`        // Use the omitempty directive
	Version   int32     `json:"version"`
}
