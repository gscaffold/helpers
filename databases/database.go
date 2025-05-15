package databases

type OpenConfig struct {
	Address  string `yaml:"address" json:"address"`
	Name     string `yaml:"name" json:"name"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	DSN      string `yaml:"dsn" json:"dsn"`
}
