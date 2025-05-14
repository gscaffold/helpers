package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

// use defaultClient:nacos
type ClientTestSuite struct {
	suite.Suite
}

func (suite *ClientTestSuite) SetupTest() {}

func (suite *ClientTestSuite) TestGet() {
	actualValue := Get("test_string")
	suite.Equal("test_string_0", actualValue, "Get error")

	actualValue = GetByKind("test", "test_string")
	suite.Equal("test_string_1", actualValue, "GetByKind error")
}

func (suite *ClientTestSuite) TestGetInt() {
	actualValue := GetInt("test_int")
	suite.Equal(99, actualValue, "GetInt error")
}

func (suite *ClientTestSuite) TestGetJSON() {
	type JSONConfig struct {
		Addr string `json:"addr"`
	}
	actualValue := JSONConfig{} //nolint:exhaustruct
	GetJSON("test_json", &actualValue)
	suite.Equal(JSONConfig{
		Addr: "localhost:9999",
	}, actualValue, "GetJSON error")
}

func (suite *ClientTestSuite) TestGetYaml() {
	type YamlConfig struct {
		Key1 []string `yaml:"key1"`
		Key2 int      `yaml:"key2"`
	}
	actualValue := YamlConfig{} //nolint:exhaustruct
	GetYaml("test_yaml", &actualValue)
	suite.Equal(YamlConfig{
		Key1: []string{"value11", "value12"},
		Key2: 2,
	}, actualValue, "GetYaml error")
}

func TestDefaultClient(t *testing.T) {
	// 没有 nacos 变量时不进行测试
	if os.Getenv(EnvNacosAddr) == "" {
		return
	}
	suite.Run(t, new(ClientTestSuite))
}
