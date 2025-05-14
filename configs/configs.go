package configs

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/gscaffold/helpers/logger"
	"gopkg.in/yaml.v2"
)

func Get(key string) string {
	realValue := getClient().Get("", key)
	return realValue
}

func GetOrDefault(key string, defaultValue string) string {
	value := Get(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetInt(key string) int {
	value := Get(key)
	if value == "" {
		return 0
	}
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		logger.Errorf(context.TODO(), "get config format int error. key:%s,value:%s err:%s", key, value, err)
	}
	return int(intValue)
}

func GetIntOrDefault(key string, defaultValue int) int {
	value := GetInt(key)
	if value == 0 {
		value = defaultValue
	}
	return value
}

func GetJson(key string, out interface{}) {
	value := Get(key)
	if len(value) == 0 {
		logger.Errorf(context.TODO(), "get config error, value is empty key:%s", key)
		return
	}
	err := json.Unmarshal([]byte(value), out)
	if err != nil {
		logger.Errorf(context.TODO(), "get config error, unmarshal error. key:%s, value:%s, err:%s", key, value, err)
	}
}

func GetYaml(key string, out interface{}) {
	value := Get(key)
	if len(value) == 0 {
		logger.Errorf(context.TODO(), "get config error, value is empty key:%s", key)
		return
	}
	err := yaml.Unmarshal([]byte(value), out)
	if err != nil {
		logger.Errorf(context.TODO(), "get config error, unmarshal error. key:%s, value:%s, err:%s", key, value, err)
	}
}

func GetByKind(kind, key string) string {
	realValue := getClient().Get(kind, key)
	return realValue
}

func GetByKindByDefault(kind, key string, defaultValue string) string {
	value := GetByKind(kind, key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetIntByKind(kind, key string) int {
	value := GetByKind(kind, key)
	if value == "" {
		return 0
	}
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		logger.Errorf(context.TODO(), "get config format int error. kind:%s, key:%s,value:%s err:%s", kind, key, value, err)
	}
	return int(intValue)
}

func GetIntByKindOrDefault(kind, key string, defaultValue int) int {
	value := GetIntByKind(kind, key)
	if value == 0 {
		value = defaultValue
	}
	return value
}

func GetJsonByKind(kind, key string, out interface{}) {
	value := GetByKind(kind, key)
	if len(value) == 0 {
		logger.Errorf(context.TODO(), "get config error, value is empty. kind:%s, key:%s", kind, key)
		return
	}
	err := json.Unmarshal([]byte(value), out)
	if err != nil {
		logger.Errorf(context.TODO(), "get config error, unmarshal error. kind:%s, key:%s, value:%s, err:%s", kind, key, value, err)
	}
}

func GetYamlByKind(kind, key string, out interface{}) {
	value := GetByKind(kind, key)
	if len(value) == 0 {
		logger.Errorf(context.TODO(), "get config error, value is empty. kind:%s, key:%s", kind, key)
		return
	}
	err := yaml.Unmarshal([]byte(value), out)
	if err != nil {
		logger.Errorf(context.TODO(), "get config error, unmarshal error. kind:%s, key:%s, value:%s, err:%s", kind, key, value, err)
	}
}

func MonitorChange(kind, key string, callback ConfigChangeCallback) {
	getClient().MonitorChange(kind, key, callback)
}
