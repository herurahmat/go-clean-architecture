package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewBook(t *testing.T) {

	newBook := CreateNewBook("Tenggelamnya Kapal Van Der Wjick")

	t.Run("failed", func(t *testing.T) {
		assert.NotEqual(t, "Terkilirnya Kapal Van Der Wjick", newBook.GetBookTitle())
	})

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, "Tenggelamnya Kapal Van Der Wjick", newBook.GetBookTitle())
	})

}

func TestGetAttributes(t *testing.T) {
	newBook := CreateNewBook("Tenggelamnya Kapal Van Der Wjick")

	assert.Equal(t, "Tenggelamnya Kapal Van Der Wjick", newBook.GetBookTitle())
}

func TestGetBookId(t *testing.T) {
	newBook := CreateNewBook("Hamka")

	assert.NotNil(t, newBook.GetBookId())
}
