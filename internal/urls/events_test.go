package urls_test

import (
	"github.com/isaquesb/url-shortener/internal/urls"
	testAssert "github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEvent(t *testing.T) {
	assert := testAssert.New(t)

	ce := urls.NewCreateEvent([]byte("SHORT1"), "https://example.com")
	assert.Equal(urls.Created, ce.GetName())
	assert.Equal([]byte("SHORT1"), ce.GetKey())
	assert.NotEmpty(ce.CreatedAt)
}

func TestVisitEvent(t *testing.T) {
	assert := testAssert.New(t)

	ve := urls.NewVisitEvent([]byte("SHORT1"))
	assert.Equal(urls.Visited, ve.GetName())
	assert.Equal([]byte("SHORT1"), ve.GetKey())
	assert.NotEmpty(ve.Date)
}

func TestDeleteEvent(t *testing.T) {
	assert := testAssert.New(t)

	de := urls.NewDeleteEvent([]byte("SHORT1"))
	assert.Equal(urls.Deleted, de.GetName())
	assert.Equal([]byte("SHORT1"), de.GetKey())
}
