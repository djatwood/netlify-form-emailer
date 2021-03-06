package formailer

//go:generate go run generate/main.go

import (
	"reflect"
	"text/template"
)

var templateFuncMap = template.FuncMap{
	"isSlice": isSlice,
}

func isSlice(v interface{}) bool {
	return "slice" == reflect.TypeOf(v).Kind().String()
}
