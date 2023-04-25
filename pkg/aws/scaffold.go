package aws

import (
	"os"
	// "fmt"
	"log"

	"github.com/TwiN/go-color"
	"text/template"
)

const create = "CREATE"

func Scaffold(project string) {

	createProjectFolder(project)

	addReadme(project)
	addGitIgnoreFile(project)

	createBackend(project)
	createResourcesDirectory(project)

	configureTerraformTools(project)
	initGitRepo(project)

}

func createProjectFolder(project string) {

	if err := os.MkdirAll(project, os.ModePerm); err != nil {
        log.Fatal(err)
    }

}

func addReadme(project string) {

	path := project + "/README.md"

	// Variables
	vars := make(map[string]interface{})
    vars["Project"] = project

    // Parse the template
    tmpl, _ := template.ParseFiles("templates/reedme.tmpl")

    // Create a new file
    file, _ := os.Create(path)
    defer file.Close()

    // Apply the template to the vars map and write the result to file.
    tmpl.Execute(file, vars)

	println()
	println(color.InGreen(create) + path)

}

func addGitIgnoreFile(project string) {

	path := project + "/.gitignore"

    // Parse the template
    tmpl, _ := template.ParseFiles("templates/gitignore.tmpl")

    // Create a new file
    file, _ := os.Create(path)
    defer file.Close()

    // Apply the template to the vars map and write the result to file.
    tmpl.Execute(file,"")

	println(color.InGreen(create) + path)

}

func createBackend(project string) {

	backend := project + "/backend"
	backendTemplate := backend + "/templates"

	if err := os.MkdirAll(backendTemplate, os.ModePerm); err != nil {
        log.Fatal(err)
    }

	addTerraformBackendFiles(backend, project)
	addTerraformBackendTemplates(backendTemplate)

}

func createResourcesDirectory(project string) {

	resources := project + "/resources"

	if err := os.MkdirAll(resources, os.ModePerm); err != nil {
        log.Fatal(err)
    }

	addResourcesDirectories(resources)
	addTerraformResourcesFiles(resources)

}

func addTerraformBackendFiles(backend string, project string) {

	path := backend + "/main.tf"

    // Parse the template
    mainTmpl, _ := template.ParseFiles("templates/aws/backend/main.tmpl")

    // Create a new file
    main, _ := os.Create(path)
    defer main.Close()

    // Apply the template to the vars map and write the result to file.
    mainTmpl.Execute(main, "")

	println(color.InGreen(create) + path)

    // Variables
	vars := make(map[string]interface{})
    vars["Project"] = project

	path = backend + "/variables.tf"

    // Parse the template
    variablesTmpl, _ := template.ParseFiles("templates/aws/backend/variables.tmpl")

    // Create a new file
    variables, _ := os.Create(path)
    defer variables.Close()

    // Apply the template to the vars map and write the result to file.
    variablesTmpl.Execute(variables, vars)

	println(color.InGreen(create) + path)

	path = backend + "/data.tf"

    // Parse the template
    dataTmpl, _ := template.ParseFiles("templates/aws/backend/data.tmpl")

    // Create a new file
    data, _ := os.Create(path)
    defer data.Close()

    // Apply the template to the vars map and write the result to file.
    dataTmpl.Execute(data, "")

	println(color.InGreen(create) + path)

}

func addTerraformBackendTemplates(backendTemplate string) {

	path := backendTemplate + "/versions.tf.tpl"

    // Parse the template
    versionsTmpl, _ := template.ParseFiles("templates/aws/resources/versions.tmpl")

    // Create a new file
    versions, _ := os.Create(path)
    defer versions.Close()

    // Apply the template to the vars map and write the result to file.
    versionsTmpl.Execute(versions, "")

	println(color.InGreen(create) + path)

}

func addResourcesDirectories(resources string) {

	modules := resources + "/modules"
	environments := resources + "/env"

	if err := os.MkdirAll(modules, os.ModePerm); err != nil {
        log.Fatal(err)
    }

	if err := os.MkdirAll(environments, os.ModePerm); err != nil {
        log.Fatal(err)
    }

	path := modules + "/.gitkeep"

	// Variables
	vars := make(map[string]interface{})
    vars["Message"] = "Add your custom modules in this directory."

    // Parse the template
    tmpl, _ := template.ParseFiles("templates/gitkeep.tmpl")

    // Create a new file
    gitkeepModules, _ := os.Create(path)
    defer gitkeepModules.Close()

    // Apply the template to the vars map and write the result to file.
    tmpl.Execute(gitkeepModules, vars)

	println(color.InGreen(create) + path)

    path = environments + "/.gitkeep"

	// Variables
    vars["Message"] = "Add the tfvars files depending on the workspaces you are going to use."

    // Create a new file
    file, _ := os.Create(path)
    defer file.Close()

    // Apply the template to the vars map and write the result to file.
    tmpl.Execute(file, vars)

	println(color.InGreen(create) + path)

}

func addTerraformResourcesFiles(resources string) {

	path := resources + "/main.tf"

    // Parse the template
    versionsTmpl, _ := template.ParseFiles("templates/aws/resources/main.tmpl")

    // Create a new file
    versions, _ := os.Create(path)
    defer versions.Close()

    // Apply the template to the vars map and write the result to file.
    versionsTmpl.Execute(versions, "")

	println(color.InGreen(create) + path)

}

func configureTerraformTools(project string) {}

func initGitRepo(project string) {}
