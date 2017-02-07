package controllers

import (
    "github.com/musale/go-blog/app/models"
    "encoding/json"
)

type BlogPostItem struct {
    GorpController
}

func (c BlogPostItem) parseBlogPost() (models.BlogPost, error) {
    blogitem := models.BlogPost{}
    err := json.NewDecoder(c.Request.Body).Decode(&blogitem)
    return blogitem, err
}


func (c BlogPostItem) Add() revel.Result {
    if blogitem, err := c.parseBidItem(); err != nil {
        return c.RenderText("Unable to parse the BlogPost from JSON.")
    } else {
        // Validate the model
        blogitem.Validate(c.Validation)
        if c.Validation.HasErrors() {
            // Do something better here!
            return c.RenderText("You have error in your BlogPost.")
        } else {
            if err := c.Txn.Insert(&blogitem); err != nil {
                return c.RenderText(
                    "Error inserting record into database!")
            } else {
                return c.RenderJson(biditem)
            }
        }
    }
}

func (c BlogPostItem) Get(id int64) revel.Result {
    blogitem := new(models.BlogPost)
    err := c.Txn.SelectOne(blogitem,
        `SELECT * FROM BlogPost WHERE id = ?`, id)
    if err != nil {
        return c.RenderText("Error. BlogPost probably doesn't exist.")
    }
    return c.RenderJson(blogitem)
}
