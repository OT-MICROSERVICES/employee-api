package client

import (
	"employee-api/config"
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
	"time"
)

// CreateScyllaDBClient is a method to create client connection for ScyllaDB
func CreateScyllaDBClient() (*gocql.Session, error) {
	config := config.ReadConfigAndProperty()
	client := gocql.NewCluster(config.ScyllaDB.Host...)
	fallback := gocql.RoundRobinHostPolicy()
	client.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(fallback)
	client.Timeout = time.Duration(1) * time.Second
	client.Keyspace = config.ScyllaDB.Keyspace
	client.Authenticator = gocql.PasswordAuthenticator{
		Username: config.ScyllaDB.Username,
		Password: config.ScyllaDB.Password,
	}
	session, err := client.CreateSession()
	if err != nil {
		logrus.Errorf("Unable to create session with scylladb: %v", err)
	}
	return session, err
}
