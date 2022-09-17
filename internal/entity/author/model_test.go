package author

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAuthor(t *testing.T) {

	action := CreateNewAuthor("Heru Rahmat")

	t.Run("failed", func(t *testing.T) {
		assert.NotEqual(t, "Heru", action.GetAuthorName())
	})

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, "Heru Rahmat", action.GetAuthorName())
	})

}

func TestGetAttributes(t *testing.T) {
	action := CreateNewAuthor("Albert Einstein")

	assert.Equal(t, "Albert Einstein", action.GetAuthorName())
}

func TestGetAuthorId(t *testing.T) {
	action := CreateNewAuthor("New Real Name")

	assert.NotNil(t, action.GetAuthorId())
}
