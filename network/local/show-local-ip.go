package main

// Package is called aw
import (
	"github.com/deanishe/awgo"
	"net"
	"strings"
)

// Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

// Your workflow starts here
func run() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		localIP := addr.String()
		if strings.Contains(localIP,"::")  {
			continue
		}
		if strings.Contains(localIP,"127.0.0.1")  {
			continue
		}
		localIP = strings.Split(localIP,"/")[0]
		wf.NewItem(localIP).Valid(true).Arg(localIP)
	}


	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
