package imgflipgo

import (
	"fmt"
	"reflect"

	"github.com/fatih/structtag"
)

func getStructFieldJSONTag(v reflect.Type, fieldName string) (string, error) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	fieldMeta, ok := v.FieldByName(fieldName)
	if !ok {
		return "", fmt.Errorf(`struct %s has no field named "%s"`, v.Name(), fieldName)
	}
	tags, err := structtag.Parse(string(fieldMeta.Tag))
	if err != nil {
		return "", err
	}
	tagData, err := tags.Get("json")
	if err != nil {
		return "", err
	}
	return tagData.Name, nil
}
