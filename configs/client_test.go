package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// use defaultClient:nacos
type ClientTestSuite struct {
	suite.Suite
}

func (suite *ClientTestSuite) SetupTest() {}

func (suite *ClientTestSuite) TestGet() {
	actualValue := Get("test_string")
	assert.Equal(suite.T(), "test_string_0", actualValue, "Get error")

	actualValue = GetByKind("test", "test_string")
	assert.Equal(suite.T(), "test_string_1", actualValue, "GetByKind error")
}

func (suite *ClientTestSuite) TestGetInt() {
	actualValue := GetInt("test_int")
	assert.Equal(suite.T(), 99, actualValue, "GetInt error")
}

func (suite *ClientTestSuite) TestGetJson() {
	type JsonConfig struct {
		Addr string `json:"addr"`
	}
	actualValue := JsonConfig{}
	GetJson("test_json", &actualValue)
	assert.Equal(suite.T(), JsonConfig{
		Addr: "localhost:9999",
	}, actualValue, "GetJson error")
}

func (suite *ClientTestSuite) TestGetYaml() {
	type YamlConfig struct {
		Key1 []string `yaml:"key1"`
		Key2 int      `yaml:"key2"`
	}
	actualValue := YamlConfig{}
	GetYaml("test_yaml", &actualValue)
	assert.Equal(suite.T(), YamlConfig{
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
