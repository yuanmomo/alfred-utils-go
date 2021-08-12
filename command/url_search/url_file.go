package command

import (
	aw "github.com/deanishe/awgo"
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
func Add(config *Config, groupName string, url Url) *Config {
	groupNameLower := strings.ToLower(groupName)
	for index, group := range config.GroupList {
		if strings.EqualFold(groupNameLower, group.GroupName) {
			// group exists
			newUrls := append(group.URLList, url)
			group.URLList = newUrls
			config.GroupList[index] = group
			return config
		}
	}

	// add new group
	newGroup := Group{
		GroupName: groupName,
		URLList:   []Url{url},
	}

	newGroups := append(config.GroupList, newGroup)
	config.GroupList = newGroups
	return config
}


func Read(wf *aw.Workflow, fileName string) *Config {
	var config Config
	wf.Data.LoadJSON(fileName, &config)
	return &config
}
