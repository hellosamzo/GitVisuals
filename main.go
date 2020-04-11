package main

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var email string = viper.GetString("email")
	//var githubRepos = viper.GetStringMap("githubRepos")

	print(email)
	var dir string
	flag.StringVar(&dir, "add", "", "add a new directory to scan for Git repositories")
	flag.StringVar(&email, "email", email, "the email to scan")
	flag.Parse()

	if dir != "" {
		scan(dir)
		return
	}

	stats(email)
}
