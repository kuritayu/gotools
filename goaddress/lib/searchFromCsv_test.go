package goaddress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeToAdd(t *testing.T) {
	err := CodeToAdd("1810002")
	assert.NoError(t, err)
}

func TestCodeToAdd_error(t *testing.T) {
	err := CodeToAdd("9999999")
	assert.NoError(t, err)
}

func TestAddToCode(t *testing.T) {
	err := AddToCode("東京都三鷹市牟礼")
	assert.NoError(t, err)
}

func TestAddToCode2(t *testing.T) {
	err := AddToCode("東京都三鷹市牟礼.*")
	assert.NoError(t, err)
}

func TestAddToCode_error(t *testing.T) {
	err := AddToCode("a")
	assert.NoError(t, err)
}
