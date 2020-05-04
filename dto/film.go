// Film entity definition

package dto

import (
	"time"
)

// variables beginning with a capital are exported
const Kind = "film"

type Film struct {
	Title string          `datastore:"title" json:"title"`
	Description string    `datastore:"description,noindex" json:"description"`
	Genres []string        `datastore:"genres" json:"genres"`
	//Genres []Genre        `datastore:"genres" json:"genres"`
	ReleaseDate time.Time `datastore:"releaseDate,noindex" json:"releaseDate"`
}