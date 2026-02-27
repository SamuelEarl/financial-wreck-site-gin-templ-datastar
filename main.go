package main

import (
	"embed"

	"financialwreck.com/site/internal/assets"
	"financialwreck.com/site/routes"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	// Link the embed to our helper package.
	assets.StaticFiles = staticFiles

	router := routes.SetupRouter()

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run()
}
