package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Setting struct {
	gorm.Model
	Key   string `gorm:"unique;not null"`
	Value string
}

type Rule struct {
	gorm.Model
	GroupId  int64 `gorm:"unique;not null"`
	RuleJson string
}

func dbInit(newToken string) (token string) {
	dbtmp, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Setting{}, &Rule{})
	var setting Setting
	db.Find(&setting, "Key=?", "token")
	token = setting.Value
	if newToken != "" {
		token = newToken
		if setting.ID > 0 {
			setting.Value = newToken
			db.Model(&setting).Update(setting)
		} else {
			db.Create(&Setting{
				Key:   "token",
				Value: newToken,
			})
		}
	}
	dbReadAllGroupRules()
	return
}

func dbAddNewGroup(groupId int64) {
	db.Create(&Rule{
		GroupId:  groupId,
		RuleJson: "",
	})
}

func dbUpdateGroupRule(groupId int64, ruleJson string) {
	db.Model(&Rule{}).Where("group_id=?", groupId).Update("rule_json", ruleJson)
}

func dbReadAllGroupRules() {
	var allGroupRules []Rule
	db.Find(&allGroupRules)
	for _, rule := range allGroupRules {
		kvs := json2kvs(rule.RuleJson)
		allkvs[rule.GroupId] = kvs
		groups = append(groups, rule.GroupId)
	}
}
