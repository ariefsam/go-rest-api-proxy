package usecase

import "github.com/ariefsam/go-rest-api-proxy/usecase/dependency"

type Usecase struct {
	HTTPClient dependency.HTTPClient
}
