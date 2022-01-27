package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testConfigFile = "testdata/test-config.json"
)

func TestParseConfigFile(t *testing.T) {
	config, err := ParseConfigFile(testConfigFile)
	assert.Nil(t, err)
	assert.NotNil(t, config)
	assert.NotEmpty(t, config.Style.Mode)
}
