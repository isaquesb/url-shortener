package urls_test

import (
	"context"
	"fmt"
	"github.com/isaquesb/url-shortener/internal/app"
	"github.com/isaquesb/url-shortener/internal/ports/input/http"
	"github.com/isaquesb/url-shortener/internal/ports/output"
	"github.com/isaquesb/url-shortener/internal/urls"
	testAssert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateShortUrl(t *testing.T) {
	assert := testAssert.New(t)

	ctx := context.Background()
	dispatcher := new(MockedDispatcher)

	app.SetApp(&app.App{
		Ctx: ctx,
		Api: &app.Api{
			Dispatcher: app.Lazy[output.Dispatcher]{
				Create: func() output.Dispatcher {
					return dispatcher
				},
			},
		},
	})

	req := new(MockedRequest)
	req.On("Ctx").Return(ctx)
	req.On("FormValue", "url").Return("https://example.com")
	req.On("Header", "Accept").Return("application/json")

	dispatcher.On("Dispatch", ctx, mock.Anything).Return(nil)

	resp, err := urls.CreateShortUrl(req)

	assert.Nil(err)
	assert.Equal(http.Created, resp.GetStatusCode())
	assert.Equal("application/json", resp.GetHeaders()["Content-Type"])

	dispatcher.AssertExpectations(t)
	req.AssertExpectations(t)
}

func TestCreateWithDispatchError(t *testing.T) {
	assert := testAssert.New(t)

	ctx := context.Background()
	dispatcher := new(MockedDispatcher)

	app.SetApp(&app.App{
		Ctx: ctx,
		Api: &app.Api{
			Dispatcher: app.Lazy[output.Dispatcher]{
				Create: func() output.Dispatcher {
					return dispatcher
				},
			},
		},
	})

	req := new(MockedRequest)
	req.On("Ctx").Return(ctx)
	req.On("FormValue", "url").Return("https://example.com")

	dispatcher.On("Dispatch", ctx, mock.Anything).Return(fmt.Errorf("dispatch error"))

	resp, err := urls.CreateShortUrl(req)

	assert.NotNil(err)
	assert.Nil(resp)

	assert.Equal("dispatch error", err.Error())

	dispatcher.AssertExpectations(t)
	req.AssertExpectations(t)
}

func TestCreateWithEmptyUrl(t *testing.T) {
	assert := testAssert.New(t)

	req := new(MockedRequest)
	req.On("FormValue", "url").Return("")

	resp, err := urls.CreateShortUrl(req)

	assert.Nil(err)
	assert.Equal(http.BadRequest, resp.GetStatusCode())

	req.AssertExpectations(t)
}
