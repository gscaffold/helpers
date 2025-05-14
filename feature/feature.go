package feature

import (
	"math/rand"

	"github.com/gscaffold/helpers/configs"
)

// 开关控制, 配置有值时为开, 否则为关
func CheckSwitch(key string) bool {
	toggle := configs.Get(key)
	return toggle != ""
}

// 放量开关. 配置 30 时, 表示放量 30%
func CheckRateLimit(key string, defaultPercent int) bool {
	percent := configs.GetIntOrDefault(key, defaultPercent)
	return percent > rand.Intn(100)
}
