# NOTES

Whenever possible, I have installed Go packages (like Templ and Air) as Go tools (i.e. locally in my project instead of globally). Since these packages are not installed globally, I won't have access to their commands in the terminal. For example, I can't run any `templ ...` or `air ...` commands. However, since Go is install globally, I can navigate to my project directory and prefix my commands with `go tool` to access my local packages that are installed as Go tools. For example, in order to use Air to run and live reload the Gin web server, I have to run `go tool air` instead of `air`.

Pros of the Project-Local (Go Modules) Approach:

* Version Control: Project A can use Gin v1.9, while Project B uses Gin v1.10. They won't conflict.
* Reproducibility: Anyone who clones your code runs go mod download and gets the exact same versions you used.
* No "GOPATH" Mess: You can keep your code anywhere on your computer, not just in one specific src folder.
