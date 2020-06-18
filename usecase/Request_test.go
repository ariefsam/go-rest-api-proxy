package usecase_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/ariefsam/go-rest-api-proxy/entity"
	"github.com/ariefsam/go-rest-api-proxy/usecase"
	"github.com/ariefsam/go-rest-api-proxy/usecase/dependency/mockdependency"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestRequest(t *testing.T) {
	var u usecase.Usecase
	parameter := entity.Parameter{
		Headers: []entity.Header{
			entity.Header{
				"k1": "v1",
			},
			entity.Header{
				"k2": "v2",
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
				"respHeaderKey": "respHeaderVal",
			},
		},
		Body: map[string]interface{}{
			"respKey1": "respBody1",
			"respKey2": "respBody2",
		},
		StatusCode: http.StatusOK,
	}

	jsonReq, _ := json.Marshal(parameter.Body)
	jsonResp, _ := json.Marshal(expectedResponse.Body)

	httpRequest, err := http.NewRequest(parameter.Method, parameter.URL, bytes.NewReader(jsonReq))
	httpRequest.Header = http.Header{
		parameter.Headers[0]["k1"]: []string{parameter.Headers[0]["v1"]},
		parameter.Headers[1]["k2"]: []string{parameter.Headers[1]["v2"]},
	}
	assert.NoError(t, err)
	httpResponse := http.Response{
		Header: http.Header{
			"respHeaderKey": []string{"respHeaderVal"},
		},
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader(jsonResp)),
	}

	var mockHTTPClient mockdependency.HTTPClient
	u.HTTPClient = &mockHTTPClient

	mockHTTPClient.CalledDoResponse = []http.Response{httpResponse}

	resp, err := u.Request(parameter)
	assert.Equal(t, httpRequest.Method, mockHTTPClient.CalledDoRequest[0].Method)
	assert.Equal(t, httpRequest.URL, mockHTTPClient.CalledDoRequest[0].URL)
	assert.Equal(t, httpRequest.Body, mockHTTPClient.CalledDoRequest[0].Body)
	assert.NoError(t, err)

	assert.Equal(t, expectedResponse, resp)

}
