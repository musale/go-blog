package models
import (
    "github.com/revel/revel"
    "time"
)

type BlogPost struct {
    Id                 int64        `db:"id" json:"id"`
    Title              string       `db:"title" json:"title"`
    Category           string       `db:"category" json:"category"`
    DateOfPublishing   time.Time    `db:"date_of_publishing" json:"date_of_publishing"`
    Author             string       `db:"author" json:"author"`
}


func (b *BlogPost) Validate(v *revel.Validation) {

    v.Check(b.Author,
        revel.ValidRequired())

    v.Check(b.DateOfPublishing,
        revel.ValidRequired())

    v.Check(b.Title,
        revel.ValidRequired())

    v.Check(b.Id,
        revel.ValidRequired())
}
