package model

// Config struct is managing the configuration file
type Config struct {
	ScyllaDB struct {
		Host     []string `yaml:"host"`
		Keyspace string   `yaml:"keyspace"`
		Username string   `yaml:"username"`
		Password string   `yaml:"password"`
	} `yaml:"scylladb"`
	Redis struct {
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Database int    `yaml:"database"`
		Enabled  bool   `yaml:"enabled"`
	} `yaml:"redis"`
}
