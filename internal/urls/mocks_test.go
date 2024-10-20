package urls_test

import (
	"context"
	inputevents "github.com/isaquesb/url-shortener/internal/events"
	"github.com/stretchr/testify/mock"
)

type MockedDispatcher struct {
	mock.Mock
}

func (m *MockedDispatcher) Dispatch(ctx context.Context, msg inputevents.Event) error {
	args := m.Called(ctx, msg)
	return args.Error(0)
}

func (m *MockedDispatcher) Close() {
	m.Called()
}

type MockedUrlRepository struct {
	mock.Mock
}

func (m *MockedUrlRepository) UrlFromShort(ctx context.Context, short string) (string, error) {
	args := m.Called(ctx, short)
	return args.String(0), args.Error(1)
}

func (m *MockedUrlRepository) StatsFromShort(ctx context.Context, short string) (map[string]interface{}, error) {
	args := m.Called(ctx, short)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

type MockedRequest struct {
	mock.Mock
}

func (m *MockedRequest) Ctx() context.Context {
	args := m.Called()
	return args.Get(0).(context.Context)
}

func (m *MockedRequest) IsGet() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedRequest) IsPost() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedRequest) IsDelete() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedRequest) FormValue(field string) string {
	args := m.Called(field)
	return args.String(0)
}

func (m *MockedRequest) PathValue(field string) any {
	args := m.Called(field)
	return args.Get(0)
}

func (m *MockedRequest) Header(field string) string {
	args := m.Called(field)
	return args.String(0)
}
