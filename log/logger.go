package log

import "context"

// Logger 定义日志接口
type Logger interface {
	Infof(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
}

// NoopLogger 为默认空实现
type NoopLogger struct{}

func (NoopLogger) Infof(ctx context.Context, format string, args ...interface{})  {}
func (NoopLogger) Warnf(ctx context.Context, format string, args ...interface{})  {}
func (NoopLogger) Errorf(ctx context.Context, format string, args ...interface{}) {}
func (NoopLogger) Debugf(ctx context.Context, format string, args ...interface{}) {}