package match

import (
	"strconv"
)

const (
	MATCHFIGHT = "matchfight" // 进程前缀
)

const (
	MATCHING = "matching" // 开始匹配
)

// 获取匹配服务名
func GetMatchName(flag int) string {
	tflag := strconv.Itoa(flag)
	return MATCHFIGHT + tflag + "v" + tflag
}
