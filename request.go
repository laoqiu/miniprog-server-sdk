package miniprog

import (
	"io"
	"net/url"
	"reflect"
	"strings"
)

// URLEncode 编码
func URLEncode(req interface{}) string {
	return structToValues(req).Encode()
}

// ToBody 转io.Reader
func ToBody(req interface{}) io.Reader {
	return strings.NewReader(URLEncode(req))
}

// structToValues struct转url.Values
func structToValues(i interface{}) url.Values {
	v := url.Values{}
	if i != nil {
		iVal := reflect.ValueOf(i).Elem()
		tp := iVal.Type()
		for i := 0; i < iVal.NumField(); i++ {
			tag := tp.Field(i).Tag.Get("json")
			if len(tag) > 0 {
				name := strings.Split(tag, ",")[0]
				if name != "-" {
					v.Set(name, iVal.Field(i).String())
				}
			}
		}
	}
	return v
}
