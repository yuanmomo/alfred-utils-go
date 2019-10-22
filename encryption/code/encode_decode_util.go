package main

// Package is called aw
import (
	"alfred-utils-go/common"
	"encoding/base64"
	"fmt"
	"github.com/deanishe/awgo"
	"net/url"
	"os"
	"strings"
)

// Workflow is the main API
var wf *aw.Workflow

func init() {
	wf = aw.New()
}

// Your workflow starts here
func run() {
	if len(os.Args) < 1 {
		// at least tow args:
		// 	first : optionType,
		// 	second : input string, cloud be null
		return
	}

	// encode or decode
	var optionType string
	// input string
	var inputString string

	optionType = os.Args[1]

	if len(os.Args) > 2 { // has input inputString
		inputString = os.Args[2]
	}

	result := make([]string,0 )

	switch strings.TrimSpace(optionType) {
	case "en":
		result = append(result, base64.StdEncoding.EncodeToString([]byte(inputString)))
	case "de":
		bytes, e := base64.StdEncoding.DecodeString(inputString)
		if e == nil {
			result = append(result, string(bytes))
		} else {
			result = append(result, "BASE64 decode failed!!!");
		}
	case "urlen":
		result = append(result, url.PathEscape(inputString))
	case "urlde":
		decodeUrl, e := url.QueryUnescape(inputString)
		if e == nil {
			result = append(result, decodeUrl)
		} else {
			result = append(result, "URL decode failed!!!");
		}
	case "md5":
		if len(strings.TrimSpace(inputString)) == 0 {
			inputString = common.RandStringRunes(32)
		}
		md5String := common.MD5(inputString)
		result = append(result, md5String )
		result = append(result, strings.ToUpper(md5String))
	default:
		result = append(result, fmt.Sprintf("Unknown optionType for %v", optionType));
	}

	for _, line := range result {
		wf.NewItem(line).Valid(true).Arg(line).Subtitle(inputString)
	}

	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
