package command

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"regexp"
	"strings"
)

type Config struct {
	GroupList []Group `json:"allGroups"`
}

type Group struct {
	GroupName string `json:"groupName"`
	URLList   []Url  `json:"urlList"`
}

type Url struct {
	Password string `json:"password"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Extra    string `json:"extra"`
}

func (config Config) Write(wf *aw.Workflow, fileName string) {
	wf.Data.StoreJSON(fileName, config)
}

// add a new url to group
func Add(config *Config, groupName string, url Url) string {
	groupNameLower := strings.ToLower(groupName)
	for index, group := range config.GroupList {
		if strings.EqualFold(groupNameLower, group.GroupName) {
			// group exists
			for _, existsUrl := range group.URLList {
				if strings.EqualFold(url.URL, existsUrl.URL) {
					// exists
					return "exists"
				}
			}

			newUrls := append(group.URLList, url)
			group.URLList = newUrls
			config.GroupList[index] = group
			return fmt.Sprintf("Add to group:[%s]", group.GroupName)
		}
	}

	// add new group
	newGroup := Group{
		GroupName: groupName,
		URLList:   []Url{url},
	}

	newGroups := append(config.GroupList, newGroup)
	config.GroupList = newGroups
	return fmt.Sprintf("Add new group:[%s]", groupName)
}

func ListGroup(config *Config) []string {
	var groupList []string
	for _, group := range config.GroupList {
		groupList = append(groupList, group.GroupName)
	}
	return groupList
}

func SearchUrl(config *Config, groupNameRegex string, urlRegex string) map[string][]Url {
	urlListMap := make(map[string][]Url)

	for _, group := range config.GroupList {
		groupMatch, _ := regexp.MatchString(groupNameRegex, group.GroupName)
		if groupMatch {
			var urlList []Url
			for _, url := range group.URLList {
				urlMatch, _ := regexp.MatchString(urlRegex, url.URL)
				if urlMatch {
					urlList = append(urlList, url)
				}
			}
			urlListMap[group.GroupName] = urlList
		}
	}
	return urlListMap
}

func Read(wf *aw.Workflow, fileName string) *Config {
	var config Config
	wf.Data.LoadJSON(fileName, &config)
	return &config
}
