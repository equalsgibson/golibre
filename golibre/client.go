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
	apiURL               string
	requestPreProcessors []RequestPreProcessor
	authentication       Authentication
	jwt                  jwtAuth
}

type jwtAuth struct {
	rawToken string
	mutex    *sync.RWMutex
}

func (c *client) do(request *http.Request, target any) error {
	if request.URL.Host == "" {
		request.URL.Host = c.apiURL
	}

	request.URL.Scheme = "https"

	// Required
	request.Header.Set("cache-control", "no-cache")
	request.Header.Set("content-type", "application/json")
	request.Header.Set("product", "llu.android")
	request.Header.Set("version", "4.8.0")

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

	if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("network request error: %d", response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// NOTE: Libreview API returns HTTP Status OK for every request, and changes the "status" field
	// in the response body based on the success / failure of a request.
	responseErr := APIError{
		RawResponse: response,
	}

	if err := json.Unmarshal(bodyBytes, &responseErr); err != nil {
		return err
	}

	switch responseErr.Status {
	case StatusOK:
		if target != nil {
			if err := json.Unmarshal(bodyBytes, target); err != nil {
				return err
			}
		}

		return nil

	case StatusUnauthenticated:
		c.jwt.rawToken = ""

		return &responseErr

	default:
		// Unknown status code, return the response error if possible
		return &responseErr
	}
}

func (c *client) Do(request *http.Request, target any) error {
	if err := c.addAuthentication(request); err != nil {
		return err
	}

	return c.do(request, target)
}

func (c *client) addAuthentication(r *http.Request) error {
	if c.jwt.rawToken != "" {
		r.Header.Set("Authorization", "Bearer "+c.jwt.rawToken)

		return nil
	}

	c.jwt.mutex.RLock()
	defer c.jwt.mutex.RUnlock()

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

	c.jwt.mutex.Lock()
	c.jwt.rawToken = target.Data.AuthTicket.Token
	c.jwt.mutex.Unlock()

	r.Header.Set("Authorization", "Bearer "+c.jwt.rawToken)

	return nil
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BaseResponse[data []ConnectionData | ConnectionGraphData | UserData | AccountDetailData] struct {
	Status uint       `json:"status"`
	Data   data       `json:"data"`
	Ticket AuthTicket `json:"ticket"`
}
