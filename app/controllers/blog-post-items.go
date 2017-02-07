package controllers

import (
    "github.com/musale/go-blog/app/models"
    "github.com/revel/revel"
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
    if blogitem, err := c.parseBlogPost(); err != nil {
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
                return c.RenderJson(blogitem)
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

func (c BlogPostItem) List() revel.Result {
    lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
    limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
    blogitems, err := c.Txn.Select(models.BlogPost{},
        `SELECT * FROM BlogPost WHERE Id > ? LIMIT ?`, lastId, limit)
    if err != nil {
        return c.RenderText(
            "Error trying to get records from DB.")
    }
    return c.RenderJson(blogitems)
}

func (c BlogPostItem) Update(id int64) revel.Result {
    blogitems, err := c.parseBlogPost()
    if err != nil {
        return c.RenderText("Unable to parse the BlogPostItem from JSON.")
    }
    // Ensure the Id is set.
    blogitems.Id = id
    success, err := c.Txn.Update(&blogitems)
    if err != nil || success == 0 {
        return c.RenderText("Unable to update bid item.")
    }
    return c.RenderText("Updated %v", id)
}

func (c BlogPostItem) Delete(id int64) revel.Result {
    success, err := c.Txn.Delete(&models.BlogPost{Id: id})
    if err != nil || success == 0 {
        return c.RenderText("Failed to remove BlogPostItem")
    }
    return c.RenderText("Deleted %v", id)
}
