package utils

type ContentType interface {
	String() string
}

type JSON struct{}

func (json JSON) String() string {
	return "application/json"
}

type PlainText struct{}

func (plainText PlainText) String() string {
	return "text/plain"
}
