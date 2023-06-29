package client

import (
	"employee-api/model"
	"errors"
	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockConfig is a mock implementation of the config.ReadConfigAndProperty function
func MockConfig() model.Config {
	return model.Config{
		ScyllaDB: model.ScyllaDB{
			Host:     []string{"127.0.0.1"},
			Keyspace: "employee_db",
			Username: "scylladb",
			Password: "password",
		},
	}
}

func TestCreateScyllaDBClient(t *testing.T) {
	expectedHosts := []string{"127.0.0.1"}
	ReadConfigAndProperty := MockConfig
	ReadConfigAndProperty()

	gocqlNewClusterMock := func(hosts []string) *gocql.ClusterConfig {
		// Assert that the provided hosts match the expected hosts
		assert.Equal(t, expectedHosts, hosts)
		return &gocql.ClusterConfig{
			Hosts: hosts,
		}
	}

	gocqlCreateSessionMock := func() (*gocql.Session, error) {
		return nil, errors.New("session creation error")
	}
	gocqlNewClusterMock(expectedHosts)
	gocqlCreateSessionMock()

	session, err := CreateScyllaDBClient()

	assert.Nil(t, session)
	assert.EqualError(t, err, "no hosts provided")

	assert.Equal(t, "Unable to create session with scylladb: session creation error", "Unable to create session with scylladb: session creation error")
}
