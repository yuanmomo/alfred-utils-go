package command

// Package is called aw
import (
	"alfred-utils-go/common"
	aw "github.com/deanishe/awgo"
	"log"
	"net/url"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

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
var codecUrl  = "https://api.yuanmomo.net/api/common/tools/urlcodec/v1"

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

	log.Printf("type[%v],param:[%v]",optionType,inputString)

	reqParams := url.Values{
		"optionType" : {optionType},
		"value" : {inputString},
	}

	var result *Result
	common.HttpPostFormReturnJson(codecUrl,reqParams,&result)
	if(result.Code == 0){
		log.Printf("%v",result)
		wf.NewItem(result.Value).Valid(true).Arg(result.Value).Subtitle(inputString)
	}else{
		log.Printf("%v",result)
		wf.NewItem(result.Message).Valid(true).Arg(result.Message).Subtitle(inputString)
	}


	return  wf
}

func init() {
	RegisterCommand(&UrlCodeCommand{})
}


