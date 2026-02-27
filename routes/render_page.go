package routes

import (
	"financialwreck.com/site/views"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

// This function will render the templ component into a gin context's Response Writer.
// You can call this function using either an integer for the status code or the status code from the "net/http" library.
// Examples:
// views.RenderPage(c, 200, "Home Page", views.Home())
// views.RenderPage(c, http.StatusOK, "Home Page", views.Home())
func RenderPage(c *gin.Context, status int, title string, template templ.Component) {
	c.Status(status)
	views.Layout(title, template).Render(c.Request.Context(), c.Writer)
}
