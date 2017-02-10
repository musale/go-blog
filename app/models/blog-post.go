package models

import (
	"time"

	"github.com/revel/revel"
)

// BlogPost struct
type BlogPost struct {
	ID               int64     `db:"id" json:"id"`
	Title            string    `db:"title" json:"title"`
	Body             string    `db:"body" json:"body"`
	Category         string    `db:"category" json:"category"`
	DateOfPublishing time.Time `db:"date_of_publishing" json:"date_of_publishing"`
	Author           string    `db:"author" json:"author"`
}

// Validate a BlogPost
func (b *BlogPost) Validate(v *revel.Validation) {

	v.Check(b.Author,
		revel.ValidRequired())

	v.Check(b.Body,
		revel.ValidRequired())

	v.Check(b.DateOfPublishing,
		revel.ValidRequired())

	v.Check(b.Title,
		revel.ValidRequired())

	v.Check(b.ID,
		revel.ValidRequired())
}
