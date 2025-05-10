package logger

import "log"

// 基础组件自动 init
func init() {
	var err error
	std, err = NewZapLogger()
	if err != nil {
		log.Fatalf("init logger failed: %v", err)
	}
}
