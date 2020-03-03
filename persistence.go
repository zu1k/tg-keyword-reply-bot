package main

import (
	"regexp"
	"strconv"
	"strings"
	"tg-keyword-reply-bot/common"
	"tg-keyword-reply-bot/db"
)

const addText = "格式要求:\r\n" +
	"`/add 关键字===回复内容`\r\n\r\n" +
	"例如:\r\n" +
	"`/add 机场===https://jiji.cool`\r\n" +
	"就会添加一条规则, 关键词是机场, 回复内容是网址"
const delText = "格式要求:\r\n" +
	"`/del 关键字`\r\n\r\n" +
	"例如:\r\n" +
	"`/del 机场`\r\n" +
	"就会删除一条规则,机器人不再回复机场关键词"

/**
 * 添加规则
 */
func addRule(gid int64, rule string) {
	rules := common.AllGroupRules[gid]
	r := strings.Split(rule, "===")
	keys, value := r[0], r[1]
	if strings.Contains(keys, "||") {
		ks := strings.Split(keys, "||")
		for _, key := range ks {
			_addOneRule(key, value, rules)
		}
	} else {
		_addOneRule(keys, value, rules)
	}
	db.UpdateGroupRule(gid, rules.String())
}

/**
 * 给addRule使用的辅助方法
 */
func _addOneRule(key string, value string, rules common.RuleMap) {
	key = strings.Replace(key, " ", "", -1)
	rules[key] = value
}

/**
 * 删除规则
 */
func delRule(gid int64, key string) {
	rules := common.AllGroupRules[gid]
	delete(rules, key)
	db.UpdateGroupRule(gid, rules.String())
}

/**
 * 获取一个群组所有规则的列表
 */
func getRuleList(gid int64) []string {
	kvs := common.AllGroupRules[gid]
	str := ""
	var strs []string
	num := 1
	group := 0
	for k, v := range kvs {
		str += "\r\n\r\n规则" + strconv.Itoa(num) + ":\r\n`" + k + "` => `" + v + "` "
		num++
		group++
		if group == 10 {
			group = 0
			strs = append(strs, str)
			str = ""
		}
	}
	strs = append(strs, str)
	return strs
}

/**
 * 查询是否包含相应的自动回复规则
 */
func findKey(gid int64, input string) string {
	kvs := common.AllGroupRules[gid]
	for keyword, reply := range kvs {
		if strings.HasPrefix(keyword, "re:") {
			keyword = keyword[3:]
			match, _ := regexp.MatchString(keyword, input)
			if match {
				return reply
			}
		} else if strings.Contains(input, keyword) {
			return reply
		}
	}
	return ""
}
