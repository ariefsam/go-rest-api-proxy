package entity

type Parameter struct {
	Headers []Header
	URL     string `json:"URL"`
	Method  string
	Body    interface{}
}
