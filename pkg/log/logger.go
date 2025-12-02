package log

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	z *zap.Logger
}

func New(env string, service string) (*Logger, error) {
	var cfg zap.Config

	if env == "prod" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.InitialFields = map[string]interface{}{
		"service": service,
	}

	cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{z: l}, nil
}

func (l *Logger) Sync() {
	_ = l.z.Sync()
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, traceFields(ctx)...)
	l.z.Info(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, traceFields(ctx)...)
	l.z.Error(msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, traceFields(ctx)...)
	l.z.Fatal(msg, fields...)
}

func traceFields(ctx context.Context) []zap.Field {
	if ctx == nil {
		return nil
	}

	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()
	if !sc.IsValid() {
		return nil
	}

	return []zap.Field{
		zap.String("trace_id", sc.TraceID().String()),
		zap.String("span_id", sc.SpanID().String()),
	}
}
