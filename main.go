package main

import (
	"flag"
)

func main() {
	var dir string
	var email string
	flag.StringVar(&dir, "add", "", "add a new directory to scan for Git repositories")
	flag.StringVar(&email, "email", "email@gmail.com", "the emai to scan")
	flag.Parse()

	if dir != "" {
		scan(dir)
		return
	}

	stats(email)
}
