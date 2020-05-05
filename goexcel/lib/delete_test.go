package goexcel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteBook(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	actual = DeleteBook("Book1.xlsx")
	assert.NoError(t, actual)
}

func TestDeleteSheet(t *testing.T) {
	actual := CreateEmptyFile("Book1.xlsx")
	assert.NoError(t, actual)
	f, err := Load("Book1.xlsx")
	assert.NoError(t, err)
	actual = CreateSheet(f, "Sheet2")
	assert.NoError(t, actual)
	actual = DeleteSheet(f, "Sheet2")
	assert.NoError(t, actual)
	actual = DeleteBook("Book1.xlsx")
	assert.NoError(t, actual)
}

func TestDeleteSheet_error(t *testing.T) {
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
