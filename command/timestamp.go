package command

// Package is called aw
import (
	aw "github.com/deanishe/awgo"
	"strconv"
	"time"
)

type TimestampCommand struct{}

func (c *TimestampCommand) Name() string {
	return "long"
}

func (c *TimestampCommand) Description() Description {
	return Description{
		Short: "Timestamp conversion tool...",
		Usage: []string{
			"long",
		},
	}
}

func (c *TimestampCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	if len(args) <= 0 {
		return wf
	}

	// input string
	inputString := args[0]
	if len(inputString) < 13 {
		inputString += "000"
	}

	millseconds, err := strconv.ParseInt(inputString, 10, 64)

	if err != nil {
		result := "Convert failed!!!"
		wf.NewItem(result).Valid(false).Arg(result).Subtitle(inputString)
		return wf
	}

	timeObj := time.UnixMilli(millseconds)

	localFormat := timeObj.Local().Format("2006-01-02 15:04:05")
	utcFormat := timeObj.UTC().Format("2006-01-02 15:04:05 UTC")

	wf.NewItem(localFormat).Valid(true).Arg(localFormat).Subtitle(inputString)
	wf.NewItem(utcFormat).Valid(true).Arg(utcFormat).Subtitle(inputString)

	return wf
}

func init() {
	RegisterCommand(&TimestampCommand{})
}
