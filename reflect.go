package sqlm

import (
	"reflect"
)

var derefCount int64 = 0
var flatCount int64 = 0

// Deref
func deRef(i interface{}) interface{} {
	if i == nil {
		return i
	}
	derefCount += 1
	typeOfI := reflect.TypeOf(i)
	switch typeOfI.Kind() {
	case reflect.Ptr:
		return deRef(reflect.ValueOf(i).Elem().Interface())
	default:
		return i
	}
}

func flat(i interface{}) []interface{} {
	flatCount += 1

	result := reflect.ValueOf([]interface{}{})

	kindOfI := reflect.TypeOf(i).Kind()
	valueOfI := reflect.ValueOf(i)

	switch kindOfI {
	case reflect.Ptr:
		result = reflect.Append(result, reflect.ValueOf(i))
	case reflect.Slice, reflect.Array:
		// Iterate the slice and flat each of them
		for index := 0; index < valueOfI.Len(); index++ {
			v := valueOfI.Index(index)
			/*
			if v.CanInterface() {
				v = v.Elem()
			}
			*/
			if v.Elem().Kind() == reflect.Slice || v.Elem().Kind() == reflect.Array {
				result = reflect.AppendSlice(result,
					reflect.ValueOf(flat(v.Interface())))
			} else {
				result = reflect.Append(result, v)
			}
		}
	default:
		result = reflect.Append(result, reflect.ValueOf(i))
	}

	back := result.Interface()
	return back.([]interface{})
}

func assign(target interface{}, value interface{}) error {
	switch t := target.(type) {
	case *string:
		*t = deRef(value).(string)
	case *int:
		*t = deRef(value).(int)
	case *int8:
		*t = deRef(value).(int8)
	case *int16:
		*t = deRef(value).(int16)
	case *int32:
		*t = deRef(value).(int32)
	case *int64:
		*t = deRef(value).(int64)
	case *float32:
		*t = deRef(value).(float32)
	case *float64:
		*t = deRef(value).(float64)
	case *interface{}:
		*t = deRef(value)
	}

	return nil
}
