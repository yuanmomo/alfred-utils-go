package command

// Package is called aw
import (
	. "fmt"
	aw "github.com/deanishe/awgo"
	"log"
	"net/url"
	"strings"
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

var codecUrl = "https://api.yuanmomo.net/api/common/tools/urlcodec/v1"

func (c *UrlCodeCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	if len(args) <= 1 {
		return wf;
	}

	// encode or decode
	var optionType string
	// input string
	var inputString string

	optionType = args[0]
	inputString = args[1]

	log.Printf("type[%v],param:[%v]", optionType, inputString)

	result := ""
	switch optionType {
	case "en":
		result = EncodeURL(inputString)
	case "de":
		result = DecodeURLContainsAnd(inputString)
	default:
		log.Println(Sprintf("Unknown urlcode type : [%v]", optionType))
	}
	wf.NewItem(result).Valid(true).Arg(result).Subtitle(inputString)


	//     Http API
	//reqParams := url.Values{
	//	"optionType" : {optionType},
	//	"value" : {inputString},
	//}
	//
	//var result *Result
	//common.HttpPostFormReturnJson(codecUrl,reqParams,&result)
	//if(result.Code == 0){
	//	log.Printf("%v",result)
	//	wf.NewItem(result.Value).Valid(true).Arg(result.Value).Subtitle(inputString)
	//}else{
	//	log.Printf("%v",result)
	//	wf.NewItem(result.Message).Valid(true).Arg(result.Message).Subtitle(inputString)
	//}

	return wf
}

func EncodeURL(str string) string {
	//return url.PathEscape(str)
	return strings.ReplaceAll(url.QueryEscape(str), "+", "%20")
}

func decodeURL(str string) string {
	//result, err := url.PathUnescape(str)
	result, err := url.QueryUnescape(str)
	if err != nil {
		Printf("ERROR : %v\n", err)
		log.Println(Sprintf("Decode ERROR : [%v] to [%v]", str, err))
		return ""
	}
	return result
}
func DecodeURLContainsAnd(str string) string {
	//result, err := url.PathUnescape(str)
	var strArray []string
	if strings.Index(str, "&") >= 0 {
		// if there is &, split by it
		strArray = strings.Split(str, "&")
	} else {
		// no &
		return decodeURL(str)
	}

	var decodeArray []string
	for _,subStr := range strArray {
		decodeArray = append(decodeArray, decodeURL(subStr))
	}
	return  strings.Join(decodeArray,"&")
}

func init() {
	RegisterCommand(&UrlCodeCommand{})
}
