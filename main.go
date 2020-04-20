package main

import (
	"flag"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func main() {
	prompt := promptui.Select{
		Label: "Select Data Type",
		Items: []string{"Day Count", "ASCII Graph"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("Displaying: %q\n", result)

	if result == "Day Count" {

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		errz := viper.ReadInConfig() // Find and read the config file

		if errz != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s", errz))
		}

		var email string = viper.GetString("email")
		var dir string = viper.GetString("githubRepos")

		flag.StringVar(&dir, "add", "", "add a new directory to scan for Git repositories")
		flag.StringVar(&email, "email", email, "the email to scan")
		flag.Parse()

		if dir != "" {
			scan(dir)
			return
		}

		stats(email)
	} else if result == "ASCII Graph" {
		fmt.Printf("insert ASCII Graph here")
	}
}
