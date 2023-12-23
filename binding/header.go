package binding

import (
	"net/textproto"
	"reflect"

	http "github.com/bogdanfinn/fhttp"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req *http.Request, obj any) error {

	if err := mapHeader(obj, req.Header); err != nil {
		return err
	}

	return validate(obj)
}

func mapHeader(ptr any, h map[string][]string) error {
	return mappingByPtr(ptr, headerSource(h), "header")
}

type headerSource map[string][]string

var _ setter = headerSource(nil)

func (hs headerSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt setOptions) (bool, error) {
	return setByForm(value, field, hs, textproto.CanonicalMIMEHeaderKey(tagValue), opt)
}
