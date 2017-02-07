package models
import time

type BlogPost struct {
    Id                 int64        `db:"id" json:"id"`
    Title              string       `db:"title" json:"title"`
    Category           string       `db:"category" json:"category"`
    DateOfPublishing   time.Time    `db:"date_of_publishing" json:"date_of_publishing"`
    Author             string       `db:"author" json:"author"`
}
