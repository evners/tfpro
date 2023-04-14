/*
Copyright © 2023 EVNERS - info@evners.com
*/
package cmd

import (
	"bufio"
    "fmt"
    "os"
    "strings"
	"errors"
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
	
	projectName, err := getProjectName(args)
		
	if err != nil {
		fmt.Println(err)
		return
	}
		
	fmt.Println(projectName, "project created!")
	
}

func getProjectName(args []string) (string, error) {

	var projectName string

	if len(args) > 0 {
		projectName = args[0]
		return projectName, nil
	} 

	projectName = InputPrompt("What name would you like to use for the new project?")	
    
	if len(projectName) < 1 {
		return "", errors.New("You must provide a name for the project")
	}

	return projectName, nil

}


// InputPrompt receives a string value using the label
func InputPrompt(label string) string {
    
	var s string
    r := bufio.NewReader(os.Stdin)
    
	for {
        fmt.Fprint(os.Stderr, label+" ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }

    return strings.TrimSpace(s)

}
