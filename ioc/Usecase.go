package ioc

import (
	"net/http"

	"github.com/ariefsam/go-rest-api-proxy/usecase"
)

func UseCase() (u usecase.Usecase) {
	u.HTTPClient = &http.Client{}
	return
}
