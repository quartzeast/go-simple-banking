package log

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"
)

type Level = slog.Level

const (
	LevelDebug = slog.LevelDebug
	LevelTrace = slog.Level(-2)
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelFatal = slog.Level(10)
)

type Logger struct {
	logger *slog.Logger
	level  *slog.LevelVar
}

func New(level slog.Level) *Logger {
	var lvl slog.LevelVar
	lvl.Set(level)

	h := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     &lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel := level.String()

				switch level {
				case LevelTrace:
					levelLabel = "TRACE"
				}

				a.Value = slog.StringValue(levelLabel)
			}
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}))

	return &Logger{logger: h, level: &lvl}
}

func (l *Logger) SetLevel(level Level) {
	l.level.Set(level)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.log(context.Background(), LevelInfo, msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.log(context.Background(), LevelWarn, msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log(context.Background(), LevelError, msg, args...)
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.log(context.Background(), LevelFatal, msg, args...)
	os.Exit(1)
}

func (l *Logger) Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	l.log(ctx, level, msg, args...)
}

// log is the low-level logging method for methods that take ...any.
// It must always be called directly by an exported logging method
// or function, because it uses a fixed call depth to obtain the pc.
func (l *Logger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if !l.logger.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	// NOTE: 这里修改 skip 为 4，*slog.log 源码中 skip 为 3
	runtime.Callers(3, pcs[:])
	pc = pcs[0]
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = l.logger.Handler().Handle(ctx, r)
}
