package router

import (
	"net/http"

	controllers "github.com/backend-services/controllers"
)

// Route defines a single route, e.g. a human readable name, HTTP method, pattern the function that will execute when the route is called.
type Route struct {
	Protected   bool
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

// List, Post By ID, Create New Post, Update Post, Delete Post
var routes = Routes{
	Route{
		Protected:   false,
		Name:        "Healthcheck API",
		Method:      "GET",
		Pattern:     "/healthcheck",
		HandlerFunc: controllers.HealthCheck,
	},
	Route{
		Protected:   false,
		Name:        "Create Blog",
		Method:      "POST",
		Pattern:     "/blog",
		HandlerFunc: controllers.CreateBlog,
	},
	Route{
		Protected:   false,
		Name:        "GET Blog By ID",
		Method:      "GET",
		Pattern:     "/blog/{id}",
		HandlerFunc: controllers.GetBlogByID,
	},
	Route{
		Protected:   false,
		Name:        "Update Blog By ID",
		Method:      "PUT",
		Pattern:     "/blog/{id}",
		HandlerFunc: controllers.UpdateBlog,
	},
	Route{
		Protected:   false,
		Name:        "Delete Blog By ID",
		Method:      "DELETE",
		Pattern:     "/blog/{id}",
		HandlerFunc: controllers.DeleteBlogByID,
	},
	Route{
		Protected:   false,
		Name:        "List Of Blogs",
		Method:      "GET",
		Pattern:     "/blogs",
		HandlerFunc: controllers.GetBlogs,
	},
}
