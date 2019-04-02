package match

import (
	Goroutine "github.com/gfandada/gserver/goroutine"
)

// 创建匹配器
// @params flag:匹配类型
func CreateMatch(flag int) error {
	_, err := Goroutine.Start(&MatchScheduler{Flag: flag})
	if err != nil {
		return err
	}
	return nil
}

// 开始匹配
// @params flag:匹配类型
func Match(userid int32, flag int) {
	Goroutine.Cast(GetMatchName(flag), MATCHING, []interface{}{userid})
}
