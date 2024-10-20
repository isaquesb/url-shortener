package urls_test

import (
	"context"
	"fmt"
	"github.com/isaquesb/url-shortener/internal/app"
	"github.com/isaquesb/url-shortener/internal/ports/input/http"
	"github.com/isaquesb/url-shortener/internal/ports/output"
	"github.com/isaquesb/url-shortener/internal/urls"
	"github.com/isaquesb/url-shortener/pkg/logger"
	testAssert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRedirectShortUrl(t *testing.T) {
	assert := testAssert.New(t)

	ctx := context.Background()
	dispatcher := new(MockedDispatcher)
	repository := new(MockedUrlRepository)

	app.SetApp(&app.App{
		Ctx: ctx,
		Api: &app.Api{
			Repository: app.Lazy[output.UrlRepository]{
				Create: func() output.UrlRepository {
					return repository
				},
			},
			Dispatcher: app.Lazy[output.Dispatcher]{
				Create: func() output.Dispatcher {
					return dispatcher
				},
			},
		},
	})

	req := new(MockedRequest)
	req.On("Ctx").Return(ctx)
	req.On("PathValue", "short").Return("xyzABC")

	repository.On("UrlFromShort", ctx, "xyzABC").Return("https://example.com", nil)
	dispatcher.On("Dispatch", ctx, mock.Anything).Return(nil)

	resp, err := urls.RedirectShort(req)

	assert.Nil(err)
	assert.Equal(http.Redirect, resp.GetStatusCode())
	assert.Equal("https://example.com", resp.GetBody())

	req.AssertExpectations(t)
	repository.AssertExpectations(t)
	dispatcher.AssertExpectations(t)
}

func TestRedirectWithNotFoundUrl(t *testing.T) {
	assert := testAssert.New(t)

	ctx := context.Background()
	dispatcher := new(MockedDispatcher)
	repository := new(MockedUrlRepository)

	app.SetApp(&app.App{
		Ctx: ctx,
		Api: &app.Api{
			Repository: app.Lazy[output.UrlRepository]{
				Create: func() output.UrlRepository {
					return repository
				},
			},
			Dispatcher: app.Lazy[output.Dispatcher]{
				Create: func() output.Dispatcher {
					return dispatcher
				},
			},
		},
	})

	req := new(MockedRequest)
	req.On("Ctx").Return(ctx)
	req.On("PathValue", "short").Return("xyzABC")

	repository.On("UrlFromShort", ctx, "xyzABC").Return("", nil)

	resp, err := urls.RedirectShort(req)

	assert.Nil(err)
	assert.Equal(http.NotFound, resp.GetStatusCode())
	assert.Equal("Not Found URL for xyzABC", resp.GetBody())

	req.AssertExpectations(t)
	repository.AssertExpectations(t)
	dispatcher.AssertExpectations(t)
}

func TestRedirectWithDispatchError(t *testing.T) {
	assert := testAssert.New(t)

	ctx := context.Background()
	dispatcher := new(MockedDispatcher)
	repository := new(MockedUrlRepository)

	logger.Setup(logger.NewNullLogger())

	app.SetApp(&app.App{
		Ctx: ctx,
		Api: &app.Api{
			Repository: app.Lazy[output.UrlRepository]{
				Create: func() output.UrlRepository {
					return repository
				},
			},
			Dispatcher: app.Lazy[output.Dispatcher]{
				Create: func() output.Dispatcher {
					return dispatcher
				},
			},
		},
	})

	req := new(MockedRequest)
	req.On("Ctx").Return(ctx)
	req.On("PathValue", "short").Return("xyzABC")

	repository.On("UrlFromShort", ctx, "xyzABC").Return("https://example.com", nil)
	dispatcher.On("Dispatch", ctx, mock.Anything).Return(fmt.Errorf("dispatch error"))

	resp, err := urls.RedirectShort(req)

	assert.Nil(err)
	assert.Equal(http.Redirect, resp.GetStatusCode())
	assert.Equal("https://example.com", resp.GetBody())

	req.AssertExpectations(t)
	repository.AssertExpectations(t)
	dispatcher.AssertExpectations(t)
}

func TestRedirectWithRepositoryError(t *testing.T) {
	assert := testAssert.New(t)

	ctx := context.Background()
	dispatcher := new(MockedDispatcher)
	repository := new(MockedUrlRepository)

	app.SetApp(&app.App{
		Ctx: ctx,
		Api: &app.Api{
			Repository: app.Lazy[output.UrlRepository]{
				Create: func() output.UrlRepository {
					return repository
				},
			},
			Dispatcher: app.Lazy[output.Dispatcher]{
				Create: func() output.Dispatcher {
					return dispatcher
				},
			},
		},
	})

	req := new(MockedRequest)
	req.On("Ctx").Return(ctx)
	req.On("PathValue", "short").Return("xyzABC")

	repository.On("UrlFromShort", ctx, "xyzABC").Return("", fmt.Errorf("repo error"))

	resp, err := urls.RedirectShort(req)

	assert.NotNil(err)
	assert.Nil(resp)
	assert.Equal("repo error", err.Error())

	req.AssertExpectations(t)
	repository.AssertExpectations(t)
	dispatcher.AssertExpectations(t)
}

func TestRedirectWithEmptyUrl(t *testing.T) {
	assert := testAssert.New(t)

	req := new(MockedRequest)
	req.On("PathValue", "short").Return("")

	resp, err := urls.RedirectShort(req)

	assert.Nil(err)
	assert.Equal(http.BadRequest, resp.GetStatusCode())
	assert.Equal("Missing 'short' field", resp.GetBody())

	req.AssertExpectations(t)
}
