package main

import (
	"encoding/json"
	"github.com/Ytsssss/go_accumulation/framwork/gorm"
)

func main() {
	//router.Init()
	//algorithm.GetStringAllAloud("abc")
	//conn.DumpSchema()
	//conn.DumpData()
	gorm.Init()
	var tool = &tool{
		Name: "hahhaha",
		Code: 1245,
	}
	bytes, err := json.Marshal(tool)
	if err != nil {
		return
	}
	gorm.GetRuleDaoIns().InsertWithDuplicate([]*gorm.Rule{
		{
			Response:   string(bytes),
			Hit:        41,
			Name:       "1233",
			UriId:      4,
			RuleStatus: 10,
			Identity:   "12334534",
			Parameters: "1pattttttrms",
		}, {
			Response:   "res2",
			Hit:        1,
			Name:       "123",
			UriId:      53,
			RuleStatus: 0,
			Identity:   "inde",
			Parameters: "parms",
		},
	})
}

type tool struct {
	Name string
	Code int
}
