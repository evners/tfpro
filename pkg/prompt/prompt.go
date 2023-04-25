package prompt

import (
    "fmt"
    "os"
	"errors"

	"github.com/manifoldco/promptui"
    "github.com/evners/tfpro/pkg/utils"
)

type promptContent struct {
    errorMsg string
    label    string
}

func GetProjectName(args []string) string {

    if len(args) > 0 {
        return args[0]
    }

	projectNamePrompt := promptContent{
        "Please provide a a name for the project.",
        "What name would you like to use for the new project?",
    }

	projectName := getInput(projectNamePrompt)
    return projectName

}

func GetProvider() string {

	providerPromptContent := promptContent{
        "Please provide a cloud provider.",
        fmt.Sprintf("Which cloud provider do you want to use?"),
    }

    provider := promptGetSelect(providerPromptContent)
	return provider

}

func getInput(pc promptContent) string {

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

func promptGetSelect(pc promptContent) string {
    items := []string{providers.AWS}
    index := -1
    var result string
    var err error

    for index < 0 {
        prompt := promptui.Select{
            Label:    pc.label,
            Items:    items,
        }

        index, result, err = prompt.Run()

        if index == -1 {
            items = append(items, result)
        }
    }

    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }

    return result
}
