// Film entity definition

package dto

import (
	"time"
)

// variables beginning with a capital are exported

const Kind = "film"

// using 'datastore' and 'json' annotations
type Actor struct {
	Name string `datastore:"name,noindex" json:"name"`
	Age  int    `datastore:"age,noindex" json:"age"`
}

type Film struct {
	Title       string   `datastore:"title" json:"title"`
	Description string   `datastore:"description,noindex" json:"description"`
	Genres      []string `datastore:"genres" json:"genres"`
	// TODO use Genre struct within Film struct
	//Genres []Genre        `datastore:"genres" json:"genres"`
	Actors      []Actor   `datastore:"actors,noindex" json:"actors"`
	ReleaseDate time.Time `datastore:"releaseDate,noindex" json:"releaseDate"`
}
