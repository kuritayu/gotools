package goaddress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	err := Update()
	assert.NoError(t, err)
}
