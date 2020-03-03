/**
 * @auther:  zu1k
 * @date:    2020/3/3
 */
package common

import "encoding/json"

type RuleMap map[string]string

var (
	AllGroupRules = make(map[int64]RuleMap)
	AllGroupId    []int64
)

func (rm RuleMap) String() string {
	s, err := json.Marshal(rm)
	if err != nil {
		return ""
	}
	return string(s)
}

func Json2kvs(rulesJson string) RuleMap {
	tkvs := make(RuleMap)
	_ = json.Unmarshal([]byte(rulesJson), &tkvs)
	return tkvs
}
