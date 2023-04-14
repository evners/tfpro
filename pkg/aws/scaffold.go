package aws

import (
	"os"
	"log"
)

func Scaffold(project string) {

	createBackend(project)
	createResourcesDirectory(project)

	addGitIgnoreFile(project)
	addReadme(project)

	configureTerraformTools(project)
	initGitRepo(project)

}

func createBackend(project string) {

	backend := project + "/backend"
	backendTemplate := backend + "/templates"

	if err := os.MkdirAll(backendTemplate, os.ModePerm); err != nil {
        log.Fatal(err)
    }

	addTerraformBackendFiles(backend)
	addTerraformBackendTemplates(backendTemplate)

}

func createResourcesDirectory(project string) {

	resources := project + "/resources"

	if err := os.MkdirAll(resources, os.ModePerm); err != nil {
        log.Fatal(err)
    }

	addResourcesModules(resources)
	addTerraformResourcesFiles(resources)

}

func addTerraformBackendFiles(backend string) {

	data, err := os.Create(backend + "/data.tf")
	
	if err != nil {
		log.Fatal(err)
    }
	
	defer data.Close()

	main, err := os.Create(backend + "/main.tf")
	
	if err != nil {
		log.Fatal(err)
    }
	
	defer main.Close()

	variables, err := os.Create(backend + "/variables.tf")
	
	if err != nil {
		log.Fatal(err)
    }
	
	defer variables.Close()

}

func addTerraformBackendTemplates(backendTemplate string) {

	versions, err := os.Create(backendTemplate + "/versions.tf.tpl")
	
	if err != nil {
		log.Fatal(err)
    }
	
	defer versions.Close()


}

func addResourcesModules(resources string) {

	modules := resources + "/modules"

	if err := os.MkdirAll(modules, os.ModePerm); err != nil {
        log.Fatal(err)
    }

	gitkeep, err := os.Create(modules + "/.gitkeep")
	
	if err != nil {
		log.Fatal(err)
    }
	
	defer gitkeep.Close()

}

func addTerraformResourcesFiles(resources string) {
	
	main, err := os.Create(resources + "/main.tf")
	
	if err != nil {
		log.Fatal(err)
    }
	
	defer main.Close()

}

func addGitIgnoreFile(project string) {

	gitignore, err := os.Create(project + "/.gitignore")
	
	if err != nil {
		log.Fatal(err)
    }
	
	defer gitignore.Close()

}

func addReadme(project string) {
	
	reedme, err := os.Create(project + "/README.md")
	
	if err != nil {
		log.Fatal(err)
    }
	
	defer reedme.Close()

}

func configureTerraformTools(project string) {}

func initGitRepo(project string) {}
