package routes

import (
	"io/fs"
	"net/http"

	"financialwreck.com/site/internal/assets"
	"financialwreck.com/site/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	datastar "github.com/starfederation/datastar/sdk/go"
)

func SetupRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Serve physical files from the /static folder.
	// If your file is at static/styles/main.css in your project:
	// * Browser requests /static/styles/main.css?v=123.
	// * Gin strips the query string: /static/styles/main.css.
	// * Gin strips the route prefix (/static): styles/main.css.
	// * Gin looks in your embed.FS for styles/main.css.
	// * Result: 404 because the file is actually at static/styles/main.css inside the embed.
	// This will "re-root" the filesystem using fs.Sub so that Gin can find the files without the `static/` prefix.
	// Create a "Sub-Filesystem" that starts INSIDE the static folder.
	// This turns "static/styles/main.css" into "styles/main.css".
	subFS, err := fs.Sub(assets.StaticFiles, "static")
	if err != nil {
		panic("Critical: static folder not found in embed: " + err.Error())
	}
	// Serve this sub-filesystem at the /static route
	router.StaticFS("/static", http.FS(subFS))

	// Serve virtual CSS from Templ's built-in `css` functions by creating a middleware that gathers all the Templ CSS (from the `css` functions in your templates).
	// NOTE: I think this allows Templ to render a <style> tag that contains only the CSS for the components used on the currently displayed page.
	// Always use a non-nil mux to prevent a 500 crash.
	emptyMux := http.NewServeMux()
	cssHandler := templ.NewCSSMiddleware(emptyMux)
	// This will serve all your "css" blocks at /styles/templ.css
	// Wrap the Templ CSS Middleware so Gin can use it.
	router.GET("/styles/templ.css", gin.WrapH(cssHandler))

	// Define a simple GET endpoint
	router.GET("/", home)
	router.GET("/ping", ping)
	router.GET("/hello", hello)
	router.GET("/counter", counter)
	router.POST("/increment", increment)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	return router
}

func home(c *gin.Context) {
	RenderPage(c, 200, "Home Page", views.Home())
}

func ping(c *gin.Context) {
	// Return JSON response
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func hello(c *gin.Context) {
	RenderPage(c, http.StatusOK, "Hello", views.Hello("John"))
}

func counter(c *gin.Context) {
	RenderPage(c, 200, "Counter", views.Counter(0))
}

// Create a struct that matches the keys you defined in your data-signals attribute in the counter.templ component.
type CounterSignals struct {
	Count int `json:"count"`
}

func increment(c *gin.Context) {
	// 1. Parse the incoming signals from the frontend
	var signals CounterSignals
	err := c.ShouldBindJSON(&signals)
	if err != nil {
		c.Status(400)
		return
	}

	// 2. Perform your logic
	signals.Count++

	// 3. Setup the Datastar Server-Sent Event (SSE) stream
	sse := datastar.NewSSE(c.Writer, c.Request)

	// 4. Render the Templ component to a string and send it as a fragment. (i.e. Push the updated fragment back to the browser.)
	// This replaces the HTML and updates the frontend signals. (i.e. Datastar will look for the ID "container" and replace it.)
	sse.MergeFragmentTempl(views.Counter(signals.Count))

	// // 4. (Alternative) Advanced: Sending "Only" a Signal Update
	// // Sometimes you don't want to re-render the HTML; you just want to update a value in the browser's memory (the signal). Datastar allows you to send a signal update without a fragment using the MergeSignals() method.

	// // Marshal the struct directly
	// payload, _ := json.Marshal(signals)

	// // This updates the 'count' variable in the browser
	// // without changing any HTML elements!
	// sse.MergeSignals(payload)
}
