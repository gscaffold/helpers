package middlewares

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gscaffold/helpers/telemetry/metrics"
)

// todo 指标 format 格式有待优化
// 记录路径的访问次数和耗时
func Metrics(prefix string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		path = strings.ReplaceAll(path, "/", ".")
		metric := prefix + "." + path
		defer metrics.TimingSince(metric, time.Now())

		ctx.Next()

		code := ctx.Writer.Status()
		metric += "." + strconv.Itoa(code)
		metrics.Incr(metric)
	}
}
