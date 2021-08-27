package command

import (
	"alfred-utils-go/common"
	aw "github.com/deanishe/awgo"
)

func Init(wf *aw.Workflow, jsonFilePath string, dbFilePath string) {
	config := Read(wf, jsonFilePath)
	db := common.NewBoltDB(dbFilePath);

	for _, group := range config.GroupList {
		for _, url := range group.URLList {



		}
	}

}
