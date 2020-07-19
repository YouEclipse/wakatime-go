package wakatime

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

// Client defines the wakatime Client.
type Client struct {
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Commits   *CommitService
	Durations *DurationService
	Stats     *StatService
}

const (
	// APIBase is the URL prefix for the wakatime API
	APIBase = "https://wakatime.com/api/v1/"
)

// NewClient returns a new Client with the given API key.
func NewClient(apikey string, httpClient *http.Client) *Client {
	c := &Client{}
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c.common = service{
		apikey: apikey,
		client: httpClient,
	}

	c.Durations = (*DurationService)(&c.common)
	c.Commits = (*CommitService)(&c.common)
	c.Stats = (*StatService)(&c.common)
	return c
}

type service struct {
	apikey string
	client *http.Client
}

func (s *service) get(ctx context.Context, endpoint string, query url.Values) ([]byte, error) {
	url := APIBase + endpoint
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request err :%w", err)
	}
	request.URL.RawQuery = query.Encode()

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(s.apikey)))

	response, err := s.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("http client do err :%w", err)
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read body err :%w", err)
	}

	if !(response.StatusCode >= 200 && response.StatusCode <= 299) {
		return nil, fmt.Errorf("response code expects 2xx, but got %d, response body %s",
			response.StatusCode, bodyBytes)
	}

	return bodyBytes, nil
}

func (s *service) post(ctx context.Context, endpoint string, query url.Values, params interface{}) ([]byte, error) {

	url := APIBase + endpoint

	var reqBody io.Reader

	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		paramsBytes, err := json.Marshal(params)
		if err != nil {
			return nil, fmt.Errorf("marshal params err :%w", err)
		}
		reqBody = bytes.NewBuffer(paramsBytes)
	} else {
		reqBody = nil
	}

	request, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("new request err :%w", err)
	}

	request.URL.RawQuery = query.Encode()

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(s.apikey)))

	response, err := s.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("http client do err :%w", err)
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read body err :%w", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("response code expects 200, but got %d, response body %s",
			response.StatusCode, bodyBytes)
	}

	return bodyBytes, nil
}

// Bool  allocates a new bool value  to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int64  allocates a new int64 value to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String  allocates a new string value to store v and returns a pointer to it.
func String(v string) *string { return &v }
