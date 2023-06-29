package model

// Config struct is managing the configuration file
type Config struct {
	ScyllaDB ScyllaDB `yaml:"scylladb"`
	Redis    Redis    `yaml:"redis"`
}

type ScyllaDB struct {
	Host     []string `yaml:"host"`
	Keyspace string   `yaml:"keyspace"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
	Enabled  bool   `yaml:"enabled"`
}
