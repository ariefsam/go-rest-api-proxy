package ioc

import (
	"net/http"

	"github.com/ariefsam/go-rest-api-proxy/usecase"
)

func Usecase() (u usecase.Usecase) {
	u.HTTPClient = &http.Client{}
	return
}
