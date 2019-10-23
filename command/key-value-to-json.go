package command

// Package is called aw
import (
	"alfred-utils-go/common"
	aw "github.com/deanishe/awgo"
	"os"
	"strings"
)

type KVToJsonCommand struct{}

func (c *KVToJsonCommand) Name() string {
	return "md5"
}

func (c *KVToJsonCommand) Description() Description {
	return Description{
		Short: "MD5 input string",
		Usage: []string{"md5 string"},
	}
}

func (c *KVToJsonCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	var inputString string

	if len(os.Args) > 2 { // has input inputString
		inputString = os.Args[2]
	}

	if len(strings.TrimSpace(inputString)) == 0 {
		inputString = common.RandStringRunes(32)
	}
	md5String := common.MD5(inputString)

	wf.NewItem(md5String).Valid(true).Arg(md5String).Subtitle(inputString)
	wf.NewItem(strings.ToUpper(md5String)).Valid(true).Arg(strings.ToUpper(md5String)).Subtitle(inputString)
	return  wf
}

func init() {
	RegisterCommand(&KVToJsonCommand{})
}


