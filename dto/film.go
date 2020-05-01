// Film entity definition

package dto

import (
	"time"
)

// variables beginning with a capital are exported
const Kind = "film"

type Film struct {
	Title string          `datastore:"title,noindex" json:"title"`
	Description string    `datastore:"description,noindex" json:"description"`
	Genres []Genre        `datastore:"genres,noindex" json:"genres"`
	ReleaseDate time.Time `datastore:"releaseDate,noindex" json:"releaseDate"`
}