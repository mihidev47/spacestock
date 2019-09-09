package restclient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"../logger"
)

// Errors
var ErrParseUrlValuesBody = errors.New("restclient: failed to convert body to url.Values")

const (
	// Parser Flags
	ParserJSON = ""
	ParserXML  = "xml"
)

type RESTClient struct {
	name string
	log  *logger.AppLogger
	*http.Client
}

// NewRESTClient initiate new RESTClient instance
func NewRESTClient(name string, timeout int64, logger *logger.AppLogger) *RESTClient {
	return &RESTClient{
		name:   name,
		log:    logger,
		Client: &http.Client{Timeout: time.Duration(timeout) * time.Millisecond},
	}
}

func (c *RESTClient) newRequest(method string, url string, query *url.Values, body []byte) (*http.Request, error) {
	// Create buffer body
	b := bytes.NewBuffer(body)
	// Create new request
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		c.log.ErrorCtx(c.name, err)
		return nil, err
	}
	// Add new query
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	// Add accept header
	req.Header.Add("Accept", "application/json")
	return req, nil
}

func (c *RESTClient) do(request *http.Request, destination interface{}, parser string) (httpStatus int, apiResponse string, err error) {
	// Execute request
	raw, err := c.Do(request)
	if err != nil {
		c.log.ErrorCtx(c.name, err)
		return httpStatus, apiResponse, err
	}
	// Close body on return
	defer raw.Body.Close()
	// Get http status
	httpStatus = raw.StatusCode
	// Read request body
	body, err := ioutil.ReadAll(raw.Body)
	if err != nil {
		c.log.ErrorCtx(c.name, err)
		return httpStatus, apiResponse, err
	}
	// Convert resBody to string
	apiResponse = string(body)
	// If destination nil, return api response and
	if destination == nil {
		return httpStatus, apiResponse, nil
	}
	c.log.DebugCtxf(c.name, "Response body: %s", apiResponse)
	// Switch content type
	switch parser {
	case ParserXML:
		err = xml.Unmarshal(body, destination)
		if err != nil {
			c.log.ErrorCtx(c.name, err)
			return httpStatus, apiResponse, err
		}
	default:
		// Parse request body as json
		err = json.Unmarshal(body, destination)
		if err != nil {
			c.log.ErrorCtx(c.name, err)
			return httpStatus, apiResponse, err
		}
	}
	// Return response body
	return httpStatus, apiResponse, nil
}

func (c *RESTClient) POST(endpointUrl string, query *url.Values, body interface{}, dest interface{}, destParser string) (httpStatus int, apiResponse string, err error) {
	// Switch types
	var b []byte
	var contentType string
	switch body.(type) {
	case []byte:
		contentType = "multipart/form-data"
	case url.Values:
		form, ok := body.(url.Values)
		if !ok {
			return httpStatus, apiResponse, ErrParseUrlValuesBody
		}
		// Convert to bytes
		b = bytes.NewBufferString(form.Encode()).Bytes()
		contentType = "application/x-www-form-urlencoded"
	default:
		b, err = json.Marshal(body)
		if err != nil {
			c.log.ErrorCtx(c.name, err)
			return httpStatus, apiResponse, err
		}
		contentType = "application/json"
	}
	c.log.Debugf("Request Body: %s", string(b))
	c.log.Debugf("Content Type: %s", contentType)
	// Create new request
	req, err := c.newRequest("POST", endpointUrl, query, b)
	if err != nil {
		return httpStatus, apiResponse, err
	}
	// Set content type
	req.Header.Add("Content-Type", contentType)
	// Execute request
	return c.do(req, dest, destParser)
}

func (c *RESTClient) GET(url string, query *url.Values, dest interface{}, destParser string) (httpStatus int, apiResponse string, err error) {
	// Create new request
	req, err := c.newRequest("GET", url, query, nil)
	if err != nil {
		return httpStatus, apiResponse, err
	}
	// Execute request
	return c.do(req, dest, destParser)
}
