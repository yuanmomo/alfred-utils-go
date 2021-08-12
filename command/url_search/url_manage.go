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
			"url search groupNameRegex urlRegex",
		},
	}
}

func (c *UrlSearchCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	if len(args) <= 0 {
		return wf
	}

	optionType := args[0]

	//log.Printf("[%s] url", optionType)

	config := Read(wf, fileName)

	switch optionType {
	case "add":
		addParams := args[1:]

		var newUrl Url
		var groupName string
		for index, value := range addParams {
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

	case "search":
		if len(args) == 1 {
			// no search param or search by groupName, show groupList name list
			groupList := ListGroup(config)
			log.Printf("list group result [%s] ", groupList)
			for _, groupName := range groupList {
				wf.NewItem(groupName).Valid(false).Subtitle("Enter to select group").Autocomplete(groupName)
			}
			return wf
		}

		// search url with groupName and urlRegex
		groupNameRegex := fmt.Sprintf(".*%s.*", args[1])
		urlRegex := ".*"
		if len(args) >= 3 {
			urlRegex = fmt.Sprintf(".*%s.*", args[2])
		}
		urlListMap := SearchUrl(config, groupNameRegex, urlRegex)
		//log.Printf("search url result [%s] with param:[%s : %s]", urlListMap, groupNameRegex, urlRegex)
		for groupName, urlList := range urlListMap {
			for _, url := range urlList {
				urlInfo := fmt.Sprintf("%s %s %s", url.URL, url.Username, url.Password)
				wf.NewItem(urlInfo).Valid(true).Arg(urlInfo).Subtitle(groupName)
			}
		}
	}
	return wf
}

func init() {
	command.RegisterCommand(&UrlSearchCommand{})
}
