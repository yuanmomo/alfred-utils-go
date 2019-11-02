package command

// Package is called aw
import (
	"bytes"
	"encoding/json"
	"fmt"
	aw "github.com/deanishe/awgo"
	"log"
	"strings"
)

type KVToJsonCommand struct{}

func (c *KVToJsonCommand) Name() string {
	return "kv2json"
}

func (c *KVToJsonCommand) Description() Description {
	return Description{
		Short: "convert string between json and k/v",
		Usage: []string{
			"2json string",
			"2kv string",
		},
	}
}

func (c *KVToJsonCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	if len(args) <= 1 { // has input inputString
		return wf
	}

	// encode or decode
	var optionType string
	// input string
	var inputString string

	optionType = args[0]
	inputString = args[1]

	log.Printf("type[%v],param:[%v]", optionType, inputString)

	switch optionType {
	case "2json":
		return kv2json(wf,inputString);
	case "2kv":
		return  json2kv(wf,inputString);
	default:
		title := fmt.Sprintf("unknow option type %s", optionType)
		wf.NewItem(title).Valid(true).Arg(title).Subtitle(inputString)
		return wf
	}
}

func kv2json(wf *aw.Workflow,kvString string) *aw.Workflow {
	mapResult := make(map[string]string)

	kvLine := make([]string,0)
	// split by line
	lineArray := strings.Split(kvString,"\n")

	for _,line := range lineArray {
		line := strings.TrimSpace(line)
		if line == ""  {
			continue
		}

		// split by blank
		kvTemp := strings.Split(line," ")
		for _, kv := range kvTemp {
			kv := strings.TrimSpace(kv)
			if kv == ""  {
				continue
			}
			kvLine = append(kvLine,kv)
		}
	}


	for _,line := range kvLine {
		// delete \n
		line = strings.TrimSpace(line)

		kvArray := strings.Split(line,":")
		if len(kvArray) == 1 {
			mapResult[kvArray[0]] = ""
			continue
		}else if len(kvArray) == 2 {
			mapResult[kvArray[0]] = kvArray[1]
			continue
		}
	}

	jsonByte,_ := json.Marshal(mapResult)

	jsonString := string(jsonByte);

	log.Print(jsonString);
	wf.NewItem(jsonString).Valid(true).Arg(jsonString).Subtitle(kvString)
	return  wf
}

func json2kv(wf *aw.Workflow,jsonString string) *aw.Workflow {
	var mapResult map[string]json.RawMessage

	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(jsonString), &mapResult); err != nil {
		title := "Unsupported format of json."
		wf.NewItem(title).Valid(true).Arg(title).Subtitle(jsonString)
		return  wf
	}
	var buffer bytes.Buffer
	for key, value := range mapResult {
		buffer.WriteString(key)
		buffer.WriteString(":")
		//
		valueString := string(value);
		if strings.Contains(valueString,"{"){
			// contains {
			buffer.WriteString("\n")
			continue
		}

		// delete head and tail double quote
		valueString = strings.TrimPrefix(valueString,"\"")
		valueString = strings.TrimSuffix(valueString,"\"")

		buffer.WriteString(valueString)
		buffer.WriteString("\n")
	}
	result := buffer.String()
	log.Print(result)
	wf.NewItem(result).Valid(true).Arg(result).Subtitle(jsonString)
	return  wf;
}

func init() {
	RegisterCommand(&KVToJsonCommand{})
}
