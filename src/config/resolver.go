package config

import (
	"MVC_DI/util"
	"reflect"
	"strconv"
	"strings"
)

func Resolve[T any](v *T) {
	val := reflect.ValueOf(v).Elem()
	processValue(val, val, val.Type(), "")
}

func processValue(root reflect.Value, val reflect.Value, typ reflect.Type, path string) {

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fieldPath := path + "." + fieldType.Name

		switch field.Kind() {
		case reflect.String:
			field.SetString(replacePlaceholders(field.String(), root, val))
		case reflect.Struct:
			processValue(root, field, fieldType.Type, fieldPath)
		}
	}
}

func replacePlaceholders(s string, root reflect.Value, val reflect.Value) string {
	for {
		start := strings.Index(s, "${")
		if start == -1 {
			break
		}
		end := strings.Index(s[start:], "}") + start
		if end == -1 {
			break
		}

		placeholder := s[start+2 : end]
		replacement := resolvePlaceholder(placeholder, root, val)
		s = s[:start] + replacement + s[end+1:]
	}
	return s
}

func resolvePlaceholder(placeholder string, root reflect.Value, val reflect.Value) string {
	if strings.HasPrefix(placeholder, ".") {
		// Relative path
		return resolveRelativePath(placeholder, val)
	}
	// Absolute path
	return resolveAbsolutePath(placeholder, root)
}

func resolveRelativePath(placeholder string, val reflect.Value) string {
	keys := strings.Split(placeholder[1:], ".")
	currentVal := val

	for _, key := range keys {
		key = util.SnakeToPascal(key)
		field, found := findField(currentVal, key)
		if !found {
			return ""
		}
		currentVal = field
	}

	return getValueAsString(currentVal)
}

func resolveAbsolutePath(placeholder string, val reflect.Value) string {
	keys := strings.Split(placeholder, ".")
	currentVal := val

	for _, key := range keys {
		key = util.SnakeToPascal(key)
		field, found := findField(currentVal, key)
		if !found {
			return ""
		}
		currentVal = field
	}

	return getValueAsString(currentVal)
}

func getValueAsString(val reflect.Value) string {
	switch val.Kind() {
	case reflect.String:
		return val.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	default:
		return ""
	}
}

func findField(val reflect.Value, key string) (reflect.Value, bool) {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		if fieldType.Name == key {
			return field, true
		}
	}
	return reflect.Value{}, false
}
