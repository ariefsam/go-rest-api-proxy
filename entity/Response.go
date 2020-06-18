package entity

type Response struct {
	Error      string
	Body       interface{}
	Headers    []Header
	StatusCode int
}
