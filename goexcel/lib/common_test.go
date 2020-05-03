package goexcel

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Load(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	_, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	os.Remove("Book1.xlsx")
}
