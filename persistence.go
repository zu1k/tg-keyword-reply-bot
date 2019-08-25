package main

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
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

const isnotAdminWarn = "你不是管理员,请勿使用此命令"

type Kvs map[string]string

var allkvs = make(map[int64]Kvs)
var token string

var groups []int64

func jsonify(kvsin Kvs) string {
	s, err := json.Marshal(kvsin)
	checkErr(err)
	return string(s)
}

func json2kvs(jsonin string) Kvs {
	tkvs := make(Kvs)
	_ = json.Unmarshal([]byte(jsonin), &tkvs)
	return tkvs
}

func loadinfo() {
	dbread()
}

/**
 * 添加规则
 */
func addRule(gid int64, rule string) {
	kvs := allkvs[gid]
	r := strings.Split(rule, "===")
	keys, value := r[0], r[1]
	if strings.Contains(keys, "||") {
		ks := strings.Split(keys, "||")
		for _, key := range ks {
			_addOneRule(key, value, kvs)
		}
	} else {
		_addOneRule(keys, value, kvs)
	}
	dbupdategroup(gid, jsonify(kvs))
}

/**
 * 给addRule使用的辅助方法
 */
func _addOneRule(key string, value string, kvs Kvs) {
	key = strings.Replace(key, " ", "", -1)
	kvs[key] = value
}

/**
 * 删除规则
 */
func delRule(gid int64, key string) {
	kvs := allkvs[gid]
	delete(kvs, key)
	dbupdategroup(gid, jsonify(kvs))
}

/**
 * 获取一个群组所有规则的列表
 */
func getRuleList(gid int64) []string {
	kvs := allkvs[gid]
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
	kvs := allkvs[gid]
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
