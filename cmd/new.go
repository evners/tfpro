/*
Copyright © 2023 EVNERS - info@evners.com
*/
package cmd

import (
	// "bufio"
    "fmt"
    "os"
	"log"
    // "strings"
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new terraform project",
	
	Run: func(cmd *cobra.Command, args []string) {
		createProject(args)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func createProject(args []string) {
	
	projectName := getProjectName(args)
	project, err := scaffoldProjectStructure(projectName)

	if err != nil {
        log.Fatal(err)
		return
    }

	fmt.Println(project, "project created!")
	
}

func getProjectName(args []string) string {

	var projectName string

	if len(args) > 0 {
		projectName = args[0]
		return projectName
	} 

    projectNamePrompt := promptContent{
        "Please provide a a name for the project.",
        "What name would you like to use for the new project?",
    }
    
	projectName = promptGetInput(projectNamePrompt)
	return projectName

}

func scaffoldProjectStructure(projectName string) (string, error) {
return projectName, nil
}

type promptContent struct {
    errorMsg string
    label    string
}

func promptGetInput(pc promptContent) string {
    
	validate := func(input string) error {
        if len(input) <= 0 {
            return errors.New(pc.errorMsg)
        }
        return nil
    }

    templates := &promptui.PromptTemplates{
        Prompt:  "{{ . }} ",
        Valid:   "{{ . | green }} ",
        Invalid: "{{ . | red }} ",
        Success: "{{ . | bold }} ",
    }

    prompt := promptui.Prompt{
        Label:     pc.label,
        Templates: templates,
        Validate:  validate,
    }

    result, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }

    return result

}