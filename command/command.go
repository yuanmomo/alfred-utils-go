package command

import (
	aw "github.com/deanishe/awgo"
	"strings"
)

type Description struct {
	Short string
	Usage []string
}

type Command interface {
	Name() string
	Description() Description
	//CheckArgs(wf *aw.Workflow, args []string) *aw.Workflow
	Execute(wf *aw.Workflow, args []string) *aw.Workflow
}

var (
	commandRegistry = make(map[string]Command)
)

func RegisterCommand(cmd Command) string {
	entry := strings.ToLower(cmd.Name())
	if entry == "" {
		return "empty command name"
	}
	commandRegistry[entry] = cmd
	return ""
}

func GetCommand(name string) Command {
	cmd, found := commandRegistry[name]
	if !found {
		return nil
	}
	return cmd
}

type hiddenCommand interface {
	Hidden() bool
}

func PrintUsage(wf *aw.Workflow) *aw.Workflow {
	for name, cmd := range commandRegistry {
		if _, ok := cmd.(hiddenCommand); ok {
			continue
		}

		for _,line := range cmd.Description().Usage {
			wf.NewItem(name).Valid(true).Autocomplete(name).Subtitle(line)
		}
	}
	return  wf
}
