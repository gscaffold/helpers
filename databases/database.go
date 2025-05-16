package databases

type OpenConfig struct {
	Address  string `yaml:"address" json:"address"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Name     string `yaml:"name" json:"name"` // 数据库名称
	DB       int    `yaml:"db" json:"db"`     // redis 数据名称
	DSN      string `yaml:"dsn" json:"dsn"`
}
