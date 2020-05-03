package goexcel

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateEmptyFile(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	os.Remove("Book1.xlsx")
}

func Test_CreateSheet(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	f, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	actual = CreateSheet(f, "Sheet2")
	assert.NoError(t, actual)
	os.Remove("Book1.xlsx")
}
