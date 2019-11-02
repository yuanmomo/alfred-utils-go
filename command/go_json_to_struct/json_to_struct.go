package command

// Package is called aw
import (
	"alfred-utils-go/command"
	"fmt"
	aw "github.com/deanishe/awgo"
	"log"
	"strings"
)

type GoJsonCommand struct{}

func (c *GoJsonCommand) Name() string {
	return "gojson"
}

func (c *GoJsonCommand) Description() command.Description {
	return command.Description{
		Short: "generate the go struct of the input json string.",
		Usage: []string{
			"json jsonString",
			"yaml jsonString",
		},
	}
}

func (c *GoJsonCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	if len(args) <= 0  {
		return wf
	}

	optionType := args[0]
	inputString := args[1]

	log.Printf("Generate type [%s] for string [%s]", optionType, inputString)

	tagList := make([]string, 0)
	tagList = append(tagList, optionType)

	var convertFloats bool
	var parser Parser
	switch optionType {
	case "json":
		parser = ParseJson
		convertFloats = true
	case "yaml":
		parser = ParseYaml
	}

	if output, err := Generate(strings.NewReader(inputString), parser, "structName", tagList, true,convertFloats); err != nil {
		title := fmt.Sprintf("Error [%s]", err.Error())
		wf.NewItem(title).Valid(true).Arg(title);
	} else {
		log.Print(string(output))
		wf.NewItem(string(output)).Valid(true).Arg(string(output)).Subtitle("Generated, press enter to copy!!")
	}
	return  wf
}

func init() {
	command.RegisterCommand(&GoJsonCommand{})
}
