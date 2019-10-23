package command

// Package is called aw
import (
	"alfred-utils-go/common"
	aw "github.com/deanishe/awgo"
	"strings"
)

type Md5Command struct{}

func (c *Md5Command) Name() string {
	return "md5"
}

func (c *Md5Command) Description() Description {
	return Description{
		Short: "MD5 input string",
		Usage: []string{"md5 string"},
	}
}

func (c *Md5Command) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	var inputString string

	if len(args) >= 1 { // has input inputString
		inputString = args[0]
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
	RegisterCommand(&Md5Command{})
}


