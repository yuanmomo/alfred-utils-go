package command

// Package is called aw
import (
	aw "github.com/deanishe/awgo"
	"net"
	"strings"
)

type ShowLocalIPCommand struct{}

func (c *ShowLocalIPCommand) Name() string {
	return "ip"
}

func (c *ShowLocalIPCommand) Description() Description {
	return Description{
		Short: "Show local IP list",
		Usage: []string{"ip"},
	}
}

func (c *ShowLocalIPCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, infs := range interfaces {
		interfaceName := infs.Name
		ips,err := infs.Addrs();
		if err != nil {
			panic(err)
		}
		for _, ip := range ips {
			localIP := ip.String()
			if strings.Contains(localIP, "::") {
				continue
			}
			if strings.Contains(localIP, "127.0.0.1") {
				continue
			}
			localIP = strings.Split(localIP, "/")[0]
			wf.NewItem(localIP).Valid(true).Arg(localIP).Subtitle(interfaceName)
		}
	}
	return  wf
}

func init() {
	RegisterCommand(&ShowLocalIPCommand{})
}


