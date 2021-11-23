// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// CreateTodoRequestBody defines model for CreateTodoRequestBody.
type CreateTodoRequestBody struct {
	Description *string `json:"description,omitempty"`
	Title       string  `json:"title"`
}

// CreateTodoResult defines model for CreateTodoResult.
type CreateTodoResult struct {
	Id int `json:"id"`
}

// Error defines model for Error.
type Error struct {
	Error string `json:"error"`
}

// GetTodoByIdResult defines model for GetTodoByIdResult.
type GetTodoByIdResult struct {
	CreatedAt   time.Time `json:"created_at"`
	Description *string   `json:"description,omitempty"`
	Id          int       `json:"id"`
	IsDone      bool      `json:"is_done"`
	Title       string    `json:"title"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UpdateTodoRequestBody defines model for UpdateTodoRequestBody.
type UpdateTodoRequestBody struct {
	Description *string `json:"description,omitempty"`
	IsDone      *bool   `json:"is_done,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// UpdateTodoResult defines model for UpdateTodoResult.
type UpdateTodoResult struct {
	CreatedAt   time.Time `json:"created_at"`
	Description *string   `json:"description,omitempty"`
	Id          int       `json:"id"`
	IsDone      bool      `json:"is_done"`
	Title       string    `json:"title"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateTodoJSONBody defines parameters for CreateTodo.
type CreateTodoJSONBody CreateTodoRequestBody

// UpdateTodoJSONBody defines parameters for UpdateTodo.
type UpdateTodoJSONBody UpdateTodoRequestBody

// CreateTodoJSONRequestBody defines body for CreateTodo for application/json ContentType.
type CreateTodoJSONRequestBody CreateTodoJSONBody

// UpdateTodoJSONRequestBody defines body for UpdateTodo for application/json ContentType.
type UpdateTodoJSONRequestBody UpdateTodoJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CreateTodo request with any body
	CreateTodoWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateTodo(ctx context.Context, body CreateTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTodoById request
	GetTodoById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateTodo request with any body
	UpdateTodoWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateTodo(ctx context.Context, id int, body UpdateTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CreateTodoWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateTodoRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateTodo(ctx context.Context, body CreateTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateTodoRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTodoById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTodoByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateTodoWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateTodoRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateTodo(ctx context.Context, id int, body UpdateTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateTodoRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreateTodoRequest calls the generic CreateTodo builder with application/json body
func NewCreateTodoRequest(server string, body CreateTodoJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateTodoRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateTodoRequestWithBody generates requests for CreateTodo with any type of body
func NewCreateTodoRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/todo")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetTodoByIdRequest generates requests for GetTodoById
func NewGetTodoByIdRequest(server string, id int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/todo/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateTodoRequest calls the generic UpdateTodo builder with application/json body
func NewUpdateTodoRequest(server string, id int, body UpdateTodoJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateTodoRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateTodoRequestWithBody generates requests for UpdateTodo with any type of body
func NewUpdateTodoRequestWithBody(server string, id int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/todo/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// CreateTodo request with any body
	CreateTodoWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateTodoResponse, error)

	CreateTodoWithResponse(ctx context.Context, body CreateTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateTodoResponse, error)

	// GetTodoById request
	GetTodoByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetTodoByIdResponse, error)

	// UpdateTodo request with any body
	UpdateTodoWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateTodoResponse, error)

	UpdateTodoWithResponse(ctx context.Context, id int, body UpdateTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateTodoResponse, error)
}

type CreateTodoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *CreateTodoResult
	JSON400      *Error
	JSON500      *Error
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r CreateTodoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateTodoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTodoByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetTodoByIdResult
	JSON400      *Error
	JSON404      *Error
	JSON500      *Error
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GetTodoByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTodoByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateTodoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UpdateTodoResult
	JSON400      *Error
	JSON404      *Error
	JSON500      *Error
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r UpdateTodoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateTodoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateTodoWithBodyWithResponse request with arbitrary body returning *CreateTodoResponse
func (c *ClientWithResponses) CreateTodoWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateTodoResponse, error) {
	rsp, err := c.CreateTodoWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateTodoResponse(rsp)
}

func (c *ClientWithResponses) CreateTodoWithResponse(ctx context.Context, body CreateTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateTodoResponse, error) {
	rsp, err := c.CreateTodo(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateTodoResponse(rsp)
}

// GetTodoByIdWithResponse request returning *GetTodoByIdResponse
func (c *ClientWithResponses) GetTodoByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetTodoByIdResponse, error) {
	rsp, err := c.GetTodoById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTodoByIdResponse(rsp)
}

// UpdateTodoWithBodyWithResponse request with arbitrary body returning *UpdateTodoResponse
func (c *ClientWithResponses) UpdateTodoWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateTodoResponse, error) {
	rsp, err := c.UpdateTodoWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateTodoResponse(rsp)
}

func (c *ClientWithResponses) UpdateTodoWithResponse(ctx context.Context, id int, body UpdateTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateTodoResponse, error) {
	rsp, err := c.UpdateTodo(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateTodoResponse(rsp)
}

// ParseCreateTodoResponse parses an HTTP response from a CreateTodoWithResponse call
func ParseCreateTodoResponse(rsp *http.Response) (*CreateTodoResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateTodoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest CreateTodoResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetTodoByIdResponse parses an HTTP response from a GetTodoByIdWithResponse call
func ParseGetTodoByIdResponse(rsp *http.Response) (*GetTodoByIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTodoByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetTodoByIdResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseUpdateTodoResponse parses an HTTP response from a UpdateTodoWithResponse call
func ParseUpdateTodoResponse(rsp *http.Response) (*UpdateTodoResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateTodoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UpdateTodoResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

