package uploadcare

import (
	"io"
	"net/http"
)

type RequestEncoder interface {
	EncodeRequest(*http.Request) error
}

type RespBodyDecoder interface {
	DecodeRespBody(io.ReadCloser) error
}