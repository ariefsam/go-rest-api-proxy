package usecase_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ariefsam/go-rest-api-proxy/entity"
	"github.com/ariefsam/go-rest-api-proxy/usecase"
	"github.com/ariefsam/go-rest-api-proxy/usecase/dependency/mockdependency"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	var u usecase.Usecase
	parameter := entity.Parameter{
		Headers: []entity.Header{
			entity.Header{
				Key:   "k1",
				Value: "v1",
			},
			entity.Header{
				Key:   "k2",
				Value: "v2",
			},
		},
		Method: "POST",
		URL:    "http://torequest.com",
		Body: map[string]interface{}{
			"keyB1": "valB1",
			"keyB2": "valB2",
		},
	}
	expectedResponse := entity.Response{
		Error: "",
		Headers: []entity.Header{
			entity.Header{
				Key:   "respHeaderKey",
				Value: "respHeaderVal",
			},
		},
		Body: map[string]interface{}{
			"respKey1": "respBody1",
			"respKey2": "respBody2",
		},
	}

	jsonReq, _ := json.Marshal(parameter.Body)
	jsonResp, _ := json.Marshal(expectedResponse.Body)

	httpRequest, err := http.NewRequest(parameter.Method, parameter.URL, bytes.NewReader(jsonReq))
	httpRequest.Header = http.Header{
		parameter.Headers[0].Key: []string{parameter.Headers[0].Value},
		parameter.Headers[1].Key: []string{parameter.Headers[1].Value},
	}
	assert.NoError(t, err)
	httpResponse := http.Response{
		Header: http.Header{
			expectedResponse.Headers[0].Key: []string{expectedResponse.Headers[0].Value},
		},
		Body: ioutil.NopCloser(bytes.NewReader(jsonResp)),
	}

	var mockHTTPClient mockdependency.HTTPClient
	u.HTTPClient = &mockHTTPClient

	mockHTTPClient.CalledDoResponse = []http.Response{httpResponse}

	resp, err := u.Request(parameter)
	assert.Equal(t, httpRequest.Method, mockHTTPClient.CalledDoRequest[0].Method)
	assert.Equal(t, httpRequest.URL, mockHTTPClient.CalledDoRequest[0].URL)
	assert.NoError(t, err)

	assert.Equal(t, expectedResponse, resp)

}
