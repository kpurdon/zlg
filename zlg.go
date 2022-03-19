// Package zlg provides an opinionated structured logger backed by zerolog.
package zlg

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once

// Logger provides limited structured logging capabilities backed by a well configured zerolog.Logger.
type Logger struct {
	pretty bool
	logger zerolog.Logger
}

// LoggerOption configures the Logger.
type LoggerOption func(l *Logger)

// Level sets the zerolog.Level of the Logger.
func Level(lvl zerolog.Level) LoggerOption {
	return func(l *Logger) {
		l.logger = l.logger.Level(lvl)
	}
}

// Pretty configures a logger to write pretty (console-text) based output instead of JSON.
func Pretty() LoggerOption {
	return func(l *Logger) {
		l.pretty = true // will replace the stack trace formatter w/ a text-based version
		w := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
		l.logger = l.logger.Output(w)
	}
}

// New initializes a new *Logger.
func New(opts ...LoggerOption) *Logger {
	once.Do(func() {
		zerolog.TimeFieldFormat = time.RFC3339
		zerolog.CallerSkipFrameCount = 3 // skip this package
		zerolog.DurationFieldInteger = true
		zerolog.ErrorStackMarshaler = marshalStack
		zerolog.TimestampFunc = func() time.Time {
			return time.Now().UTC()
		}
	})

	l := &Logger{
		logger: zerolog.
			New(os.Stderr).
			Level(zerolog.InfoLevel). // default level == info
			With().
			Timestamp(). // always include a timestamp
			Caller().    // always include caller information
			Logger(),
	}
	for _, opt := range opts {
		opt(l)
	}

	if l.pretty {
		zerolog.ErrorStackMarshaler = marshalStackPretty
	}

	return l
}

func (l *Logger) WithLevel(lvl zerolog.Level) *Logger {
	l.logger = l.logger.Level(lvl)
	return l
}

// With embeds the given key/value in the Logger returning the new Logger.
func (l *Logger) With(k string, v interface{}) *Logger {
	if v == nil {
		return l
	}

	// TODO: might need to support a lot more options here (generation?)
	switch vv := v.(type) {
	case bool:
		l.logger = l.logger.With().Bool(k, vv).Logger()
	case string:
		l.logger = l.logger.With().Str(k, vv).Logger()
	case int:
		l.logger = l.logger.With().Int(k, vv).Logger()
	case uint:
		l.logger = l.logger.With().Uint(k, vv).Logger()
	case int64:
		l.logger = l.logger.With().Int64(k, vv).Logger()
	case uint64:
		l.logger = l.logger.With().Uint64(k, vv).Logger()
	case float64:
		l.logger = l.logger.With().Float64(k, vv).Logger()
	case time.Duration:
		l.logger = l.logger.With().Dur(k, vv).Logger()
	case time.Time:
		l.logger = l.logger.With().Str(k, vv.Format(time.RFC3339Nano)).Logger()
	default:
		l.logger = l.logger.With().Interface(k, vv).Logger()
	}

	return l
}

// WithError embeds a given error in the Logger.
// Prefer to use the *Logger.Error method when logging at the error level.
func (l *Logger) WithError(err error) *Logger {
	l.logger = l.logger.With().Stack().Err(err).Logger()
	return l
}

// Debug writes the given message at the debug level.
func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

// Info writes the given message at the info level.
func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

// Error writes the given error and stack at the error level.
func (l *Logger) Error(err error) {
	l.logger.Error().Stack().Err(err).Send()
}

// Panic writes the given message at error level and then calls panic.
// Generally this will be used along with an error Logger.WithError(err).Panic("something").
func (l *Logger) Panic(msg string) {
	l.logger.Panic().Msg(msg)
}

func marshalStack(err error) interface{} {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	var e stackTracer
	if !errors.As(err, &e) {
		return nil
	}

	return pkgerrors.MarshalStack(err)
}

func marshalStackPretty(err error) interface{} {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	var e stackTracer
	if !errors.As(err, &e) {
		return nil
	}

	for _, frame := range e.StackTrace() {
		fmt.Printf("%+s:%d\r\n", frame, frame)
	}

	return nil // excludes the stack from the log since it's printed
}
