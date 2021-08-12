package command

// Package is called aw
import (
	"alfred-utils-go/command"
	"fmt"
	aw "github.com/deanishe/awgo"
	"log"
)

const fileName = "url.json"

type UrlSearchCommand struct{}

func (c *UrlSearchCommand) Name() string {
	return "url"
}

func (c *UrlSearchCommand) Description() command.Description {
	return command.Description{
		Short: "search url from configurations.",
		Usage: []string{
			"url add groupName url userName password",
			"url search type keyword",
		},
	}
}

func (c *UrlSearchCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	if len(args) <= 0 {
		return wf
	}

	optionType := args[0]
	params := args[1:]

	log.Printf("[%s] url [%s]", optionType, params)

	config := Read(wf, fileName)

	switch optionType {
	case "add":
		var newUrl Url
		var groupName string
		for index, value := range params {
			switch index {
			case 0:
				groupName = value
			case 1:
				newUrl.URL = value
			case 2:
				newUrl.Username = value
			case 3:
				newUrl.Password = value
			default:
				newUrl.Extra = fmt.Sprintf("%s %s", newUrl.Extra, value)
			}
		}
		Add(config, groupName, newUrl)
		config.Write(wf, fileName)
		fmt.Printf("Add success!!!\n", )
	case "search":

		log.Print(string(output))
		wf.NewItem(string(output)).Valid(true).Arg(string(output)).Subtitle("Generated, press enter to copy!!")
	}

	return wf
}

func init() {
	command.RegisterCommand(&UrlSearchCommand{})
}
