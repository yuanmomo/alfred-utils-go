package command

// Package is called aw
import (
	aw "github.com/deanishe/awgo"
	"golang.design/x/clipboard"
	"strings"
)

type SystemTestLogParseCommand struct{}

func (d *SystemTestLogParseCommand) Name() string {
	return "token-log"
}

func (d *SystemTestLogParseCommand) Description() Description {
	return Description{
		Short: "Parse the user reference and the token from the system-test's log line.",
		Usage: []string{
			"token-log string ",
		},
	}
}

func (c *SystemTestLogParseCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	logLineFromClipboard := string(clipboard.Read(clipboard.FmtText)[:])
	if !strings.Contains(logLineFromClipboard, "----") {
		wf.NewItem("No log line found in clipboard.").Valid(false).Subtitle("Valid log line should contains '----'")
		return wf
	}

	itemAndValuesPart := strings.Split(logLineFromClipboard, "----")[1]

	itemAndValueArray := strings.Split(strings.TrimSpace(itemAndValuesPart), "|")

	for _, itemAndValue := range itemAndValueArray {
		valueArray := strings.Split(strings.TrimSpace(itemAndValue), ":")
		wf.NewItem(valueArray[1]).Valid(true).Arg(valueArray[1]).Subtitle(valueArray[0])
	}

	return wf
}

func init() {
	RegisterCommand(&SystemTestLogParseCommand{})

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

}
