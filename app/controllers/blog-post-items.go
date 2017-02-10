package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/musale/go-blog/app/models"
	"github.com/musale/go-blog/app/routes"
	"github.com/revel/revel"
)

// BlogPostItem structure
type BlogPostItem struct {
	GorpController
}

func (c BlogPostItem) parseBlogPost() (models.BlogPost, error) {
	c.Request.ParseForm()
	fmt.Println("body:: ", c.Request.Body)
	blogitem := models.BlogPost{}
	err := json.NewDecoder(c.Request.Body).Decode(&blogitem)
	return blogitem, err
}

// NewPost is used to render view for adding a new blog item
func (c BlogPostItem) NewPost() revel.Result {
	return c.Render(c)
}

// Add is the endpoint to receive POST requests to add a new blogitem
func (c BlogPostItem) Add() revel.Result {
	if blogitem, err := c.parseBlogPost(); err != nil {
		fmt.Println("JSON ERROR: ", err)
		return c.RenderText("Unable to parse the BlogPost from JSON.")
	} else {
		// Validate the model
		blogitem.Validate(c.Validation)
		if c.Validation.HasErrors() {
			// Do something better here!
			fmt.Println("VALIDATION ERROR: ", c.Validation.ErrorMap())
			return c.RenderText("You have error in your BlogPost.")
		} else {
			if err := c.Txn.Insert(&blogitem); err != nil {
				fmt.Print("DB INSERT ERROR: ", err)
				return c.RenderText(
					"Error inserting record into database!")
			} else {
				return c.Redirect(routes.BlogPostItem.List())
			}
		}
	}
}

// Get a BlogPost using id
func (c BlogPostItem) Get(id int64) revel.Result {
	blogitem := new(models.BlogPost)
	err := c.Txn.SelectOne(blogitem,
		`SELECT * FROM BlogPost WHERE id = ?`, id)
	if err != nil {
		return c.RenderText("Error. BlogPost probably doesn't exist.")
	}
	return c.RenderJson(blogitem)
}

// List all blogitems
func (c BlogPostItem) List() revel.Result {
	lastID := parseIntOrDefault(c.Params.Get("lid"), -1)
	limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
	blogitems, err := c.Txn.Select(models.BlogPost{},
		`SELECT * FROM BlogPost WHERE ID > ? LIMIT ?`, lastID, limit)
	if err != nil {
		fmt.Println("FETCH LIST ERROR: ", err)
		return c.RenderText(
			"Error trying to get records from DB.")
	}
	return c.Render(blogitems)
}

// Update a BlogPostItem
func (c BlogPostItem) Update(id int64) revel.Result {
	blogitems, err := c.parseBlogPost()
	if err != nil {
		return c.RenderText("Unable to parse the BlogPostItem from JSON.")
	}
	// Ensure the Id is set.
	blogitems.ID = id
	success, err := c.Txn.Update(&blogitems)
	if err != nil || success == 0 {
		return c.RenderText("Unable to update bid item.")
	}
	return c.RenderText("Updated %v", id)
}

// Delete a BlogPost
func (c BlogPostItem) Delete(id int64) revel.Result {
	success, err := c.Txn.Delete(&models.BlogPost{ID: id})
	if err != nil || success == 0 {
		return c.RenderText("Failed to remove BlogPostItem")
	}
	return c.RenderText("Deleted %v", id)
}
