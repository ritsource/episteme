package models

import (
	"github.com/ritsource/episteme/prototype/data/protobuf/dst/postpb"
)

type Post struct {
	postpb.Post
}

func (p *Post) Example() string {
	return ""
}

// func (p *Post) GetMsg() postpb.Post {
//	return &p.msg
// }

func (p *Post) Save() {
	// msg := p.GetMsg()

	// if msg.Id ==  {
	// 	p.id = uuid.New()
	// }

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
