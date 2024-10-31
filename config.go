package golibre

import (
	"fmt"
	"log/slog"
	"net/http"
)

type configOption func(s *config)

type config struct {
	transport            *http.Transport
	existingJWTToken     string
	requestPreProcessors []RequestPreProcessor
}

type RequestPreProcessor interface {
	ProcessRequest(r *http.Request) error
}

type RequestPreProcessorFunc func(*http.Request) error

func (p RequestPreProcessorFunc) ProcessRequest(r *http.Request) error {
	return p(r)
}

func WithExistingJWTToken(existingToken string) configOption {
	return func(s *config) {
		s.existingJWTToken = existingToken
	}
}

func WithTLSInsecureSkipVerify() configOption {
	return func(s *config) {
		s.transport.TLSClientConfig.InsecureSkipVerify = true
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
