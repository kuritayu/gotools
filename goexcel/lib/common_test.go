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

func Test_IsExistedSheet(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	f, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	actual = CreateSheet(f, "Sheet2")
	assert.NoError(t, actual)
	actual = DeleteSheet(f, "a")
	assert.Error(t, actual)
	actual = DeleteBook("Book1.xlsx")
	assert.NoError(t, actual)
}
