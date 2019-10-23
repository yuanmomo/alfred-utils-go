package command

// Package is called aw
import (
	aw "github.com/deanishe/awgo"
	"net/url"
	"strings"
)

type UrlCodeCommand struct{}

func (c *UrlCodeCommand) Name() string {
	return "urlcode"
}

func (c *UrlCodeCommand) Description() Description {
	return Description{
		Short: "URL encode/decode input string",
		Usage: []string{
			"urlcode en string",
			"urlcode de string",
		},
	}
}

func (c *UrlCodeCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	if len(args) <= 1 {
		return  wf;
	}

	// encode or decode
	var optionType string
	// input string
	var inputString string

	optionType = args[0]
	inputString = args[1]

	switch strings.TrimSpace(optionType) {
	case "en":
		result := url.PathEscape(inputString)
		wf.NewItem(result).Valid(true).Arg(result).Subtitle(inputString)
	case "de":
		bytes, e := url.QueryUnescape(inputString)
		if e == nil {
			result := string(bytes)
			wf.NewItem(result).Valid(true).Arg(result).Subtitle(inputString)
		} else {
			result := "URL decode failed!!!";
			wf.NewItem(result).Valid(true).Arg(result).Subtitle(inputString)
		}
	}

	return  wf
}

func init() {
	RegisterCommand(&UrlCodeCommand{})
}


