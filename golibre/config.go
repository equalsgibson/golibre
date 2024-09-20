package golibre

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type configOption func(s *config)

type config struct {
	transport            *http.Transport
	timeout              time.Duration
	userAgent            string
	requestPreProcessors []RequestPreProcessor
}

type RequestPreProcessor interface {
	ProcessRequest(r *http.Request) error
}

type RequestPreProcessorFunc func(*http.Request) error

func (p RequestPreProcessorFunc) ProcessRequest(r *http.Request) error {
	return p(r)
}

func WithTLSInsecureSkipVerify() configOption {
	return func(s *config) {
		s.transport.TLSClientConfig.InsecureSkipVerify = true
	}
}

func SetCustomTimeout(customTimeout time.Duration) configOption {
	return func(s *config) {
		s.timeout = customTimeout
	}
}

func SetCustomUserAgent(customUserAgent string) configOption {
	return func(s *config) {
		s.userAgent = customUserAgent
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