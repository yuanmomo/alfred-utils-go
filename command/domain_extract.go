package command

// Package is called aw
import (
	. "fmt"
	aw "github.com/deanishe/awgo"
	"log"
	"net/url"
)

type DomainExtractCommand struct{}

func (d *DomainExtractCommand) Name() string {
	return "domain"
}

func (d *DomainExtractCommand) Description() Description {
	return Description{
		Short: "extract multiple format of domain from url",
		Usage: []string{
			"domain string",
		},
	}
}

func (c *DomainExtractCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	if len(args) < 1 {
		return wf
	}

	// input url
	var inputUrl string

	inputUrl = args[0]

	log.Printf("input url:[%v]", inputUrl)

	u, err := url.Parse(inputUrl)
	if err != nil {
		wf.NewItem("Parse url error").Valid(true).Arg(Sprintf("Parse url error:[%s]", inputUrl)).Subtitle(inputUrl)
		log.Fatal(err)
		return wf
	}

	// Hostname
	wf.NewItem(u.Hostname()).Valid(true).Arg(u.Hostname()).Subtitle(inputUrl)

	// Hostname + path
	hostPath := Sprintf("%s%s", u.Hostname(), u.Path)
	wf.NewItem(hostPath).Valid(true).Arg(hostPath).Subtitle(inputUrl)

	return wf
}

func init() {
	RegisterCommand(&DomainExtractCommand{})
}
