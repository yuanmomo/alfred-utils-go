package command

// Package is called aw
import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"strconv"
	"strings"
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

	//
	_, err := strconv.Atoi(inputString)
	if err == nil {
		// string to long
		if len(inputString) < 13 {
			inputString += "000"
		}

		// long to string
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
	} else {
		// string to second and mill-seconds
		var t time.Time
		var e error
		if strings.Index(inputString, "T") > 0 {
			// 转换 yyyy-MM-ddTHH:mm:ss 格式，做自动补齐
			dateArray := strings.Split(inputString, "T")
			if len(dateArray) == 1 {
				// 输入 2022-06-21T 的格式
				inputString = fmt.Sprintf("%sT00:00:00", dateArray[0])
			} else {
				// 输入 2022-06-21T 的格式
				timeStringArray := strings.Split(dateArray[1], ":")
				timeArray := [3]string{"00", "00", "00"}
				for i := range timeStringArray {
					if len(timeStringArray[i]) == 0 {
						timeStringArray[i] = "0"
					}

					timeNumber, _ := strconv.Atoi(timeStringArray[i])
					timeArray[i] = fmt.Sprintf("%02d", timeNumber)[0:2]
				}

				inputString = fmt.Sprintf("%sT%s:%s:%s", dateArray[0], timeArray[0], timeArray[1], timeArray[2])
			}

			// contains space
			t, e = time.Parse(`2006-01-02T15:04:05`, inputString)
		} else {
			t, e = time.Parse(`2006-01-02`, inputString)
		}
		if e != nil {
			result := "Convert failed!!!"
			wf.NewItem(result).Valid(false).Arg(result).Subtitle(inputString)
			return wf
		}

		secondsStr := strconv.Itoa(int(t.Unix()))
		millSecondsStr := strconv.Itoa(int(t.UnixMilli()))

		wf.NewItem(secondsStr).Valid(true).Arg(secondsStr).Subtitle(inputString)
		wf.NewItem(millSecondsStr).Valid(true).Arg(millSecondsStr).Subtitle(inputString)
	}

	return wf
}

func init() {
	RegisterCommand(&TimestampCommand{})
}
