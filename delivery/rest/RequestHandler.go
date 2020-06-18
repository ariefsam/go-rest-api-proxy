package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ariefsam/go-rest-api-proxy/entity"
	"github.com/ariefsam/go-rest-api-proxy/ioc"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	var param entity.Parameter
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v", param)

	u := ioc.Usecase()
	resp, _ := u.Request(param)
	JSONView(w, resp, http.StatusOK)
}
