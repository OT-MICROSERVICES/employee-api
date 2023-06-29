package client

import (
	"employee-api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRedisClient(t *testing.T) {
	expectedAddr := "localhost:6379"
	expectedPassword := ""
	expectedDB := 0

	readConfigAndPropertyMock := func() model.Config {
		return model.Config{
			Redis: model.Redis{
				Host:     expectedAddr,
				Password: expectedPassword,
				Database: expectedDB,
			},
		}
	}

	ReadConfigAndProperty := readConfigAndPropertyMock
	ReadConfigAndProperty()

	client := CreateRedisClient()

	assert.NotNil(t, client)
	assert.Equal(t, expectedAddr, client.Options().Addr)
	assert.Equal(t, expectedPassword, client.Options().Password)
	assert.Equal(t, expectedDB, client.Options().DB)
}
