package urls_test

import (
	"github.com/isaquesb/url-shortener/internal/urls"
	testAssert "github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeCreateEvent(t *testing.T) {
	assert := testAssert.New(t)

	ce, _ := urls.MakeEvent(urls.Created)

	assert.Equal(urls.Created, ce.GetName())
}

func TestMakeVisitEvent(t *testing.T) {
	assert := testAssert.New(t)

	ve, _ := urls.MakeEvent(urls.Visited)

	assert.Equal(urls.Visited, ve.GetName())
}

func TestMakeDeleteEvent(t *testing.T) {
	assert := testAssert.New(t)

	de, _ := urls.MakeEvent(urls.Deleted)

	assert.Equal(urls.Deleted, de.GetName())
}

func TestMakeEventNotFound(t *testing.T) {
	assert := testAssert.New(t)

	_, err := urls.MakeEvent("unknown")
	assert.NotNil(err)
	assert.Equal("event not found: unknown", err.Error())
}

func TestEventParserFor(t *testing.T) {
	assert := testAssert.New(t)

	cb := urls.EventParserFor(urls.Created)

	assert.NotNil(cb)

	ce, _ := cb()
	assert.Equal(urls.Created, ce.GetName())

	vb := urls.EventParserFor(urls.Visited)
	assert.NotNil(vb)

	ve, _ := vb()
	assert.Equal(urls.Visited, ve.GetName())

	db := urls.EventParserFor(urls.Deleted)
	assert.NotNil(db)

	de, _ := db()
	assert.Equal(urls.Deleted, de.GetName())
}
