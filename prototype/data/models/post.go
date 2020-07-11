package models

import (
	"github.com/google/uuid"
)

type Post struct {
	id uuid.UUID

	title       string
	description string

	links      []Link
	categories []Category
	tags       []string
}

func (p *Post) Save() {
	if p.id == uuid.Nil {
		p.id = uuid.New()
	}

}

type Link struct {
	src         string
	description string
	pattern     Pattern
}

type Pattern struct {
	value string
}

type Category struct {
	title string
}
