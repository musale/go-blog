package controllers

import (
    "github.com/revel/revel"
    "regexp"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
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
