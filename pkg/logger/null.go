package logger

import (
	"context"
	"log/slog"
)

type NullLogger struct {
	stdOut *slog.Logger
	stdErr *slog.Logger
}

func NewNullLogger() Logger {
	return &NullLogger{}
}

func (l *NullLogger) With(args ...any) Logger {
	return &NullLogger{}
}

func (l *NullLogger) Info(_ string, _ ...any) {
}

func (l *NullLogger) Warn(_ string, _ ...any) {
}

func (l *NullLogger) Error(_ string, _ ...any) {
}

func (l *NullLogger) Debug(_ string, _ ...any) {
}

func (l *NullLogger) InfoContext(_ context.Context, _ string, _ ...any) {
}

func (l *NullLogger) WarnContext(_ context.Context, _ string, _ ...any) {
}

func (l *NullLogger) ErrorContext(_ context.Context, _ string, _ ...any) {
}

func (l *NullLogger) DebugContext(_ context.Context, _ string, _ ...any) {
}
