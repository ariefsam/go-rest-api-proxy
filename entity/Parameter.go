package entity

type Parameter struct {
	Headers []Header
	URL     string
	Method  string
	Body    interface{}
}
