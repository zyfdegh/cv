package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCVGenCmd(t *testing.T) {
	c, err := NewCVGenCmd(testConfigFile)
	assert.Nil(t, err)
	assert.NotNil(t, c)
}

func TestCVGenCmd_Run(t *testing.T) {

}
