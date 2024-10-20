package urls_test

import (
	"context"
	"fmt"
	"github.com/isaquesb/url-shortener/internal/app"
	"github.com/isaquesb/url-shortener/internal/ports/input/http"
	"github.com/isaquesb/url-shortener/internal/ports/output"
	"github.com/isaquesb/url-shortener/internal/urls"
	testAssert "github.com/stretchr/testify/assert"
	"testing"
)

func TestShortUrlStats(t *testing.T) {
	assert := testAssert.New(t)

	ctx := context.Background()
	repository := new(MockedUrlRepository)

	app.SetApp(&app.App{
		Ctx: ctx,
		Api: &app.Api{
			Repository: app.Lazy[output.UrlRepository]{
				Create: func() output.UrlRepository {
					return repository
				},
			},
		},
	})

	req := new(MockedRequest)
	req.On("Ctx").Return(ctx)
	req.On("PathValue", "short").Return("xyzABC")

	repository.On("StatsFromShort", ctx, "xyzABC").Return(map[string]interface{}{"visits": "18"}, nil)

	resp, err := urls.ShowStats(req)

	assert.Nil(err)
	assert.Equal(http.Ok, resp.GetStatusCode())
	assert.Equal("application/json", resp.Header("Content-Type"))
	assert.JSONEq(`{"visits":"18"}`, resp.GetBody())

	req.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestStatsWithNotFoundUrl(t *testing.T) {
	assert := testAssert.New(t)

	ctx := context.Background()
	repository := new(MockedUrlRepository)

	app.SetApp(&app.App{
		Ctx: ctx,
		Api: &app.Api{
			Repository: app.Lazy[output.UrlRepository]{
				Create: func() output.UrlRepository {
					return repository
				},
			},
		},
	})

	req := new(MockedRequest)
	req.On("Ctx").Return(ctx)
	req.On("PathValue", "short").Return("xyzABC")

	repository.On("StatsFromShort", ctx, "xyzABC").Return(nil, nil)

	resp, err := urls.ShowStats(req)

	assert.Nil(err)
	assert.Equal(http.NotFound, resp.GetStatusCode())
	assert.Equal("Not Found URL for xyzABC", resp.GetBody())

	req.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestStatsWithRepositoryError(t *testing.T) {
	assert := testAssert.New(t)

	ctx := context.Background()
	repository := new(MockedUrlRepository)

	app.SetApp(&app.App{
		Ctx: ctx,
		Api: &app.Api{
			Repository: app.Lazy[output.UrlRepository]{
				Create: func() output.UrlRepository {
					return repository
				},
			},
		},
	})

	req := new(MockedRequest)
	req.On("Ctx").Return(ctx)
	req.On("PathValue", "short").Return("xyzABC")

	repository.On("StatsFromShort", ctx, "xyzABC").Return(nil, fmt.Errorf("repo error"))

	resp, err := urls.ShowStats(req)

	assert.NotNil(err)
	assert.Nil(resp)
	assert.Equal("repo error", err.Error())

	req.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestStatsWithEmptyUrl(t *testing.T) {
	assert := testAssert.New(t)

	req := new(MockedRequest)
	req.On("PathValue", "short").Return("")

	resp, err := urls.ShowStats(req)

	assert.Nil(err)
	assert.Equal(http.BadRequest, resp.GetStatusCode())
	assert.Equal("Missing 'short' field", resp.GetBody())

	req.AssertExpectations(t)
}
