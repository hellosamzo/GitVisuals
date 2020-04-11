package main

import (
	"flag"
)

// scan a path and its children directories
// search for git repos
func scan(path string) {
	print("scan func")
}

// stats of commits etc
func stats(email string) {
	print("stats")
}

func main() {
	var dir string
	var email string
	flag.StringVar(&dir, "add", "", "add a new directory to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the emai to scan")
	flag.Parse()

	if dir != "" {
		scan(dir)
		return
	}

	stats(email)
}
