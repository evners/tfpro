/*
Copyright © 2023 EVNERS - info@evners.com
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/TwiN/go-color"
	"github.com/evners/tfpro/pkg/aws"
	"github.com/evners/tfpro/pkg/gcp"
	"github.com/evners/tfpro/pkg/azure"
	"github.com/evners/tfpro/pkg/prompt"
	"github.com/evners/tfpro/pkg/utils"
)

// Represents the new command
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
	
	projectName  := prompt.GetProjectName(args)
	provider 	 := prompt.GetProvider()
	
	switch provider {
		case providers.AWS:
			aws.Scaffold(projectName)
		case providers.GCP:
			gcp.Scaffold(projectName)
		case providers.AZURE:
			azure.Scaffold(projectName)
    }

	println()
	println("Project " + projectName + " created! " + color.InGreen("OK"))
	println()
	
}