// This file was auto-generated by Fern from our API Definition.

package body

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	vitalgo "github.com/tryVital/vital-go"
	core "github.com/tryVital/vital-go/core"
	io "io"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

// Get Daily Body data for user_id
func (c *Client) Get(ctx context.Context, userId string, request *vitalgo.BodyGetRequest) (*vitalgo.ClientBodyResponse, error) {
	baseURL := "https://api.tryvital.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v2/summary/body/%v", userId)

	queryParams := make(url.Values)
	if request.Provider != nil {
		queryParams.Add("provider", fmt.Sprintf("%v", *request.Provider))
	}
	queryParams.Add("start_date", fmt.Sprintf("%v", request.StartDate))
	if request.EndDate != nil {
		queryParams.Add("end_date", fmt.Sprintf("%v", *request.EndDate))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 422:
			value := new(vitalgo.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *vitalgo.ClientBodyResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get Daily Body data for user_id
func (c *Client) GetRaw(ctx context.Context, userId string, request *vitalgo.BodyGetRawRequest) (*vitalgo.RawBody, error) {
	baseURL := "https://api.tryvital.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v2/summary/body/%v/raw", userId)

	queryParams := make(url.Values)
	if request.Provider != nil {
		queryParams.Add("provider", fmt.Sprintf("%v", *request.Provider))
	}
	queryParams.Add("start_date", fmt.Sprintf("%v", request.StartDate))
	if request.EndDate != nil {
		queryParams.Add("end_date", fmt.Sprintf("%v", *request.EndDate))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 422:
			value := new(vitalgo.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *vitalgo.RawBody
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return nil, err
	}
	return response, nil
}