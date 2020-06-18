package mockdependency

import (
	"net/http"
)

type HTTPClient struct {
	CalledDoRequest  []http.Request
	CalledDoResponse []http.Response
}

func (m *HTTPClient) Do(req *http.Request) (response *http.Response, err error) {
	m.CalledDoRequest = append(m.CalledDoRequest, *req)
	response = &m.CalledDoResponse[len(m.CalledDoResponse)-1]
	return
}
