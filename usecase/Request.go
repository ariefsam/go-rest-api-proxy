package usecase

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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
		for k, v := range val {
			if _, ok := request.Header[k]; !ok {
				request.Header[k] = []string{v}
			} else {
				request.Header[k] = append(request.Header[k], v)
			}
		}
	}
	if err != nil {
		return
	}
	response, err := u.HTTPClient.Do(request)

	if err != nil {
		log.Println(err)
		return
	}
	resp.StatusCode = response.StatusCode
	// log.Printf("%+v", response.Header)
	for key, val := range response.Header {
		for _, v := range val {
			header := entity.Header{key: v}
			resp.Headers = append(resp.Headers, header)
		}
	}
	body, err := ioutil.ReadAll(response.Body)

	json.Unmarshal(body, &resp.Body)

	return
}
