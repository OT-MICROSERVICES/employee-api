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
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}
