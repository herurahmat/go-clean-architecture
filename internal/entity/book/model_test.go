package book

import (
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewBook(t *testing.T) {

	newAuthor := author.CreateNewAuthor("Hamka")

	action := CreateNewBook("Terkilirnya Kapal Van Der Wjick", *newAuthor)

	t.Run("failed", func(t *testing.T) {
		assert.NotEqual(t, "Tenggelamnya Kapal Van Der Wjick", action.GetBookTitle())
		assert.NotEqual(t, "Bang Toyib", action.Author.GetAuthorName())
	})

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, "Terkilirnya Kapal Van Der Wjick", action.GetBookTitle())
		assert.Equal(t, "Hamka", action.Author.GetAuthorName())
	})

}

func TestGetAttributes(t *testing.T) {
	newAuthor := author.CreateNewAuthor("Hamka")
	action := CreateNewBook("Terkilirnya Kapal Van Der Wjick", *newAuthor)

	assert.Equal(t, "Terkilirnya Kapal Van Der Wjick", action.GetBookTitle())
}

func TestGetAuthorId(t *testing.T) {
	newAuthor := author.CreateNewAuthor("Hamka")
	action := CreateNewBook("New Real Title", *newAuthor)

	assert.NotNil(t, action.GetBookId())
}
