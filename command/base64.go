package command

// Package is called aw
import (
	"encoding/base64"
	aw "github.com/deanishe/awgo"
	"strings"
)

type Base64Command struct{}

func (c *Base64Command) Name() string {
	return "base64"
}

var subCommand []string = []string{"en","de"}

func (c *Base64Command) Description() Description {
	return Description{
		Short: "BASE64 encode/decode input string",
		Usage: []string{
			"base64 en string",
			"base64 de string",
		},
	}
}

func (c *Base64Command) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
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
		result := base64.StdEncoding.EncodeToString([]byte(inputString))
		wf.NewItem(result).Valid(true).Arg(result).Subtitle(inputString)
	case "de":
		bytes, e := base64.StdEncoding.DecodeString(inputString)
		if e == nil {
			result := string(bytes)
			wf.NewItem(result).Valid(true).Arg(result).Subtitle(inputString)
		} else {
			result := "BASE64 decode failed!!!";
			wf.NewItem(result).Valid(true).Arg(result).Subtitle(inputString)
		}
	}

	return  wf
}

func init() {
	RegisterCommand(&Base64Command{})
}


