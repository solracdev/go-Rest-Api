package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID      int64  `bson:"id" json:"id"`
	Title   string `bson:"title" json:"title"`
	Content string `bson:"concent" json:"content"`
	//Author    User      `bson:"user" json:"user"`
	AuthorID  int64     `bson:"author_id" json:"author_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	//p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Valitade() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}

	if p.Content == "" {
		return errors.New("Required Content")
	}

	if p.AuthorID < 1 {
		return errors.New("Required Author")
	}

	return nil
}
