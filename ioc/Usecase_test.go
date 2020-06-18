package ioc_test

import (
	"testing"

	"github.com/ariefsam/go-rest-api-proxy/ioc"
	"github.com/ariefsam/go-rest-api-proxy/usecase"
	"github.com/stretchr/testify/assert"
)

func TestUsecase(t *testing.T) {
	var u usecase.Usecase
	u = ioc.UseCase()
	assert.NotNil(t, u.HTTPClient)
}
