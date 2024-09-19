package golibre

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type client struct {
	httpClient           *http.Client
	userAgent            string
	subDomain            string
	requestPreProcessors []RequestPreProcessor
	authentication       Authentication
	jwt                  jwtAuth
}

type jwtAuth struct {
	rawToken string
	mutex    *sync.Mutex
}

func (c *client) do(request *http.Request, target any) error {
	if request.URL.Host == "" {
		request.URL.Host = fmt.Sprintf("%s.libreview.io", c.subDomain)
	}

	request.URL.Scheme = "https"

	// Optional
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", c.userAgent)

	// Required
	request.Header.Set("Accept-Encoding", "gzip")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "KeepAlive")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Product", "llu.android")
	request.Header.Set("Version", "4.7.0")

	for _, requestPreProcessor := range c.requestPreProcessors {
		if err := requestPreProcessor.ProcessRequest(request); err != nil {
			return err
		}
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode >= http.StatusBadRequest {
		// If the status is Bad Request, wipe out the Cached JWT data and attempt to reauthenticate once.
		c.jwt.rawToken = ""

		if err := c.addAuthentication(request); err != nil {
			return err
		}
	}

	if target != nil {
		if err := json.Unmarshal(bodyBytes, target); err != nil {
			return err
		}
	}

	return nil
}

func (c *client) Do(request *http.Request, target any) error {
	if err := c.addAuthentication(request); err != nil {
		return err
	}

	return c.do(request, target)
}

func (c *client) addAuthentication(r *http.Request) error {
	if c.jwt.rawToken != "" {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.jwt.rawToken))

		return nil
	}

	c.jwt.mutex.Lock()
	defer c.jwt.mutex.Unlock()

	if c.jwt.rawToken != "" {
		return nil
	}

	authenticationRequestBody, err := json.Marshal(c.authentication)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodPost,
		"/llu/auth/login",
		bytes.NewReader(authenticationRequestBody),
	)
	if err != nil {
		return err
	}

	target := LoginResponse{}

	if err := c.do(req, &target); err != nil {
		return err
	}

	c.jwt.rawToken = target.Data.AuthTicket.Token

	r.Header.Set("Bearer", c.jwt.rawToken)

	return nil
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
