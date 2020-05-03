package goexcel

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrintValue(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	f, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	actual = SetValue(f, "Sheet1", "A1", "TestValue")
	assert.NoError(t, actual)
	actual = PrintValue(f, "Sheet1", "A1")
	assert.NoError(t, actual)
	os.Remove("Book1.xlsx")
}

func Test_Dump(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	f, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	actual = SetValue(f, "Sheet1", "A1", "TestValue")
	assert.NoError(t, actual)
	actual = Dump(f, "Sheet1", "\t")
	assert.NoError(t, actual)
	os.Remove("Book1.xlsx")
}

func Test_PrintSheetName(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	f, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	actual = PrintSheetName(f)
	assert.NoError(t, actual)
	os.Remove("Book1.xlsx")
}
