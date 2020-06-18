package usecase

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ariefsam/go-rest-api-proxy/entity"
)

func (u *Usecase) Request(param entity.Parameter) (resp entity.Response, err error) {
	jsonReq, err := json.Marshal(param.Body)
	if err != nil {
		return
	}
	request, err := http.NewRequest(param.Method, param.URL, bytes.NewReader(jsonReq))
	for _, val := range param.Headers {
		if _, ok := request.Header[val.Key]; !ok {
			request.Header[val.Key] = []string{val.Value}
		} else {
			request.Header[val.Key] = append(request.Header[val.Key], val.Value)
		}
	}
	if err != nil {
		return
	}
	response, err := u.HTTPClient.Do(request)
	if err != nil {
		return
	}
	for key, val := range response.Header {
		for _, v := range val {
			var header entity.Header
			header.Key = key
			header.Value = v
			resp.Headers = append(resp.Headers, header)
		}
	}
	body, err := ioutil.ReadAll(response.Body)

	json.Unmarshal(body, &resp.Body)

	return
}
