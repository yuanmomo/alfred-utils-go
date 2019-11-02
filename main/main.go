package main

import (
	"os"

	"alfred-utils-go/command"
	_ "alfred-utils-go/command/go_json_to_struct"
	"github.com/deanishe/awgo"
)

// Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

func getCommandName() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return ""
}

func run() {
	name := getCommandName()
	cmd := command.GetCommand(name)
	if cmd == nil {
		//title := fmt.Sprintf("Unknown command: %v",name);
		//wf.NewItem(title).Arg(title);
		wf = command.PrintUsage(wf)
	}else {
		wf = cmd.Execute(wf, os.Args[2:])
	}

	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
