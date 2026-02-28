# The .DEFAULT_GOAL defines which target is run when no target is specified. In this case, the default is the build target. So when you run `make` (without specifying a target), you will see output like this:
# $ make
# go clean ./...
# go fmt ./...
# go vet ./...
# go build
.DEFAULT_GOAL := build

# The .PHONY line keeps `make` from getting confused if a directory or file in your project has the same name as one of the listed targets.
.PHONY: kill dev clean fmt vet build

# Remove binary files and any generated _templ.go files.
clean:
	@echo "🧹 Cleaning up..."
	@rm -rf ./tmp
	@rm -rf ./bin
	@find . -name "*_templ.go" -delete
	@echo "✅ Clean complete!"

# Kill any orphan processes before starting the app.
kill:
	-kill -9 $(lsof -t -i:8080) || true  # Main application
	-kill -9 $(lsof -t -i:4040) || true  # Proxy server

# Run the full development suite
# 	go run main.go
# 	go tool air
#   (go tool templ generate --watch) & (go tool air)
# dev: kill
dev: clean
	@clear
	@echo "🚀 Starting Air with Datastar support..."
	@(sleep 5 && echo "\n💡 TIP: \n💡 Use localhost:4040 (the Proxy Port) instead of localhost:8080 (the App Port). \n💡 Air's proxy injects a tiny script that can automatically refresh the browser when the server restarts. \n💡 This will help keep your frontend state in sync with your backend changes without manual refreshes.\n") &
	@go tool air

# TODO: Update the target commands in this file so I use the following `build` and `clean` targets instead of the ones at the bottom of this file.

# # Manual build
# build:
# 	go tool templ generate
# 	go build -o bin/main main.go

# This has been automated in the .air.toml file >> [build] >> cmd config.
# templgen:
# 	go tool templ generate

# Each possible operation is called a target and the following definitions are the target definitions. 
# The word before the colon (:) is the name of the target. 
# Any words after the target (like `vet` in the line `build: vet`) are the other targets that must be run before the specified target runs. 
# The tasks that are performed by the target are on the indented lines after the target.
# Using ./... tells a Go tool to apply the command to all the files in the current directory and all subdirectories.
# clean:
# 	go clean ./...

# fmt: clean
# 	go fmt ./...

# vet: fmt
# 	go vet ./...

# build: vet
# 	go build


# ----------------------------------------------------------------
# Node.js processes for working with CSS icons during development
# ----------------------------------------------------------------
install-node-packages:
	cd generate-css-icons && npm install

generate-css-icons:
	cd generate-css-icons && node index.js
