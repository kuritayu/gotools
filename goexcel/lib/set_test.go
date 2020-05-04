package goexcel

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetValue(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	f, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	actual = SetValue(f, "Sheet1", "A1", "TestValue")
	assert.NoError(t, actual)
	os.Remove("Book1.xlsx")
}

func Test_SetValue_error(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	f, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	actual = SetValue(f, "Sheet1", "A", "TestValue")
	assert.Error(t, actual)
	os.Remove("Book1.xlsx")
}

func Test_RenameSheet(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	f, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	actual = RenameSheet(f, "Sheet1", "Sheet2")
	assert.NoError(t, actual)
	os.Remove("Book1.xlsx")
}

func Test_RenameSheet_error(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	f, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	actual = RenameSheet(f, "a", "Sheet2")
	assert.Error(t, actual)
	os.Remove("Book1.xlsx")
}
