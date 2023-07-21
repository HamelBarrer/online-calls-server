package utils

import (
	"fmt"
	"reflect"
)

func ValidationEmptyValues[T comparable](data T) (string, bool) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		required := field.Tag.Get("required")

		if required == "true" && reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface()) {
			if field.Anonymous {
				innerValues := reflect.ValueOf(value.Interface())

				for j := 0; j < innerValues.NumField(); j++ {
					innerV := innerValues.Field(j)
					innerT := innerValues.Type().Field(j)

					required := innerT.Tag.Get("required")

					if required == "true" && reflect.DeepEqual(innerV.Interface(), reflect.Zero(innerV.Type()).Interface()) {
						return fmt.Sprintf("%s is required", innerT.Name), true
					}
				}
			}

			return fmt.Sprintf("%s is required", field.Name), true
		}
	}

	return "", false
}
