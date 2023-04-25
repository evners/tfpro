package config

import (
    "os"
    "github.com/TwiN/go-color"
	"text/template"
)

func ConfigurePrecommit(project string) {

    path := project + "/.pre-commit-config.yaml"

    // Parse the template
    tmpl, _ := template.ParseFiles("templates/pre-commit-config.tmpl")

    // Create a new file
    file, _ := os.Create(path)
    defer file.Close()

    // Apply the template to the vars map and write the result to file.
    tmpl.Execute(file, "")
	println(color.InGreen("CREATE") + path)

}

