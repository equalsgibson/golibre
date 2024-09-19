package golibre

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type configOption func(s *config)

type config struct {
	roundTripper         http.RoundTripper
	userAgent            string
	timeout              time.Duration
	requestPreProcessors []RequestPreProcessor
}

type RequestPreProcessor interface {
	ProcessRequest(r *http.Request) error
}

type RequestPreProcessorFunc func(*http.Request) error

func (p RequestPreProcessorFunc) ProcessRequest(r *http.Request) error {
	return p(r)
}

func WithRoundTripper(roundTripper http.RoundTripper) configOption {
	return func(s *config) {
		s.roundTripper = roundTripper
	}
}

func SetTimeout(timeout time.Duration) configOption {
	return func(s *config) {
		s.timeout = timeout
	}
}

func WithUserAgent(userAgent string) configOption {
	return func(s *config) {
		s.userAgent = userAgent
	}
}

func WithRequestPreProcessor(requestPreProcessor RequestPreProcessor) configOption {
	return func(s *config) {
		s.requestPreProcessors = append(s.requestPreProcessors, requestPreProcessor)
	}
}

func WithSlogger(logger *slog.Logger) configOption {
	return WithRequestPreProcessor(&sloggerWrapper{
		logger: logger,
	})
}

type sloggerWrapper struct {
	logger *slog.Logger
}

func (s *sloggerWrapper) ProcessRequest(r *http.Request) error {
	s.logger.Debug(fmt.Sprintf("Request: %s %s", r.Method, r.URL.String()))

	return nil
}
