package urls_test

import (
    "github.com/isaquesb/url-shortener/internal/app"
    inputevents "github.com/isaquesb/url-shortener/internal/events"
    "github.com/isaquesb/url-shortener/internal/ports/output"
    "github.com/isaquesb/url-shortener/internal/urls"
    testAssert "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

func TestPersistHandler(t *testing.T) {
	assert := testAssert.New(t)

	dispatcher := new(MockedDispatcher)

	app.SetApp(&app.App{
		Worker: &app.Worker{
			Dispatcher: app.Lazy[output.Dispatcher]{
				Create: func() output.Dispatcher {
					return dispatcher
				},
			},
		},
	})

	msg := &inputevents.Message{
		Uuid:  "1234567890",
		Name:  "MyUrlCreated",
		Event: urls.NewCreateEvent([]byte("SHORT1"), "https://example.com"),
	}

	sub := urls.EventSubscriberFor(urls.Created)
	assert.NotNil(sub)
	assert.Equal("CreatedUrlHandler", sub.Name)

	evt, err := sub.ParseEvent()

	assert.Nil(err)
	assert.Equal(urls.Created, evt.GetName())

	dispatcher.On("Dispatch", app.GetApp().Ctx, mock.Anything).Return(nil)

	err = sub.Handler(msg)
	assert.Nil(err)

	dispatcher.AssertCalled(t, "Dispatch", app.GetApp().Ctx, msg.Event)
}
