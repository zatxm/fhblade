package binding

import (
	"bytes"
	"encoding/xml"
	"io"

	http "github.com/bogdanfinn/fhttp"
)

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (xmlBinding) Bind(req *http.Request, obj any) error {
	return decodeXML(req.Body, obj)
}

func (xmlBinding) BindBody(body []byte, obj any) error {
	return decodeXML(bytes.NewReader(body), obj)
}
func decodeXML(r io.Reader, obj any) error {
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}
