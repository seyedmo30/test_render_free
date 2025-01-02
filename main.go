package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Struct to hold JSONPlaceholder post data
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Route for GET request
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// Route for GET request with a parameter
	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, "Hello, "+name+"!")
	})

	// Route for POST request
	e.POST("/echo", func(c echo.Context) error {
		body := map[string]string{}
		if err := c.Bind(&body); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, body)
	})

	// Route to fetch a post by ID from JSONPlaceholder
	e.GET("/posts/:id", func(c echo.Context) error {
		postID := c.Param("id")
		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", postID)

		// Make the API call
		resp, err := http.Get(url)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch post",
			})
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return c.JSON(resp.StatusCode, map[string]string{
				"error": "Post not found",
			})
		}

		// Decode the JSON response
		var post Post
		if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to decode response",
			})
		}

		return c.JSON(http.StatusOK, post)
	})

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
