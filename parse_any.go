package dot

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func MakeTypedVar(targetType reflect.Type, initialValue any) any {
	res := reflect.New(targetType).Elem()
	if initialValue != nil {
		res.Set(reflect.ValueOf(initialValue))
	}

	return res
}

// ParseTypedVar parses a any into a value of the specified reflect.Type and returns it as any.
// Returns nil and an error if parsing fails or the type is unsupported.
// If the type implements Scanner, it uses the Scan method for parsing.
//
//nolint:gocognit,exhaustive,cyclop,funlen
func ParseTypedVar(targetType reflect.Type, input any) (result any, err error) {
	type Scanner interface {
		Scan(src any) error
	}

	// Check if the type implements Scanner
	if reflect.PointerTo(targetType).Implements(reflect.TypeOf((*Scanner)(nil)).Elem()) {
		result, err = func() (any, error) {
			// Create a new instance of the type
			val := reflect.New(targetType).Interface()
			scanner, ok := val.(Scanner)
			if !ok {
				return nil, fmt.Errorf("type %v claims to implement Scanner but does not", targetType)
			}
			if err = scanner.Scan(input); err != nil {
				return nil, fmt.Errorf("scanner failed for type %v: %w", targetType, err)
			}
			return reflect.ValueOf(val).Elem().Interface(), nil
		}()
		if err == nil {
			// return on success
			return result, nil
		}
	}

	inputStr := func() string {
		var inputStr string
		switch v := input.(type) {
		case string:
			inputStr = v
		case []byte:
			inputStr = string(v)
		default:
			inputStr = fmt.Sprintf("%v", input)
		}
		return inputStr
	}
	inputType := reflect.TypeOf(input).Kind()

	// Handle built-in types
	switch targetType.Kind() {
	case reflect.String:
		var val string
		switch inputType {
		case reflect.String:
			val = reflect.ValueOf(input).String()
		default:
			val = inputStr()
		}
		retVal := reflect.New(targetType).Elem()
		retVal.SetString(val)
		return retVal.Interface(), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var val int64
		switch inputType {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val = reflect.ValueOf(input).Int()
		case reflect.String:
			if val, err = strconv.ParseInt(input.(string), 10, targetType.Bits()); err != nil {
				return nil, err
			}
		default:
			if val, err = strconv.ParseInt(inputStr(), 10, targetType.Bits()); err != nil {
				return nil, err
			}
		}
		retVal := reflect.New(targetType).Elem()
		retVal.SetInt(val)
		return retVal.Interface(), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var val uint64
		switch inputType {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val = reflect.ValueOf(input).Uint()
		case reflect.String:
			if val, err = strconv.ParseUint(input.(string), 10, targetType.Bits()); err != nil {
				return nil, err
			}
		default:
			if val, err = strconv.ParseUint(inputStr(), 10, targetType.Bits()); err != nil {
				return nil, err
			}
		}
		retVal := reflect.New(targetType).Elem()
		retVal.SetUint(val)
		return retVal.Interface(), nil

	case reflect.Float32, reflect.Float64:
		var val float64
		switch inputType {
		case reflect.Float32, reflect.Float64:
			val = reflect.ValueOf(input).Float()
		case reflect.String:
			if val, err = strconv.ParseFloat(inputStr(), targetType.Bits()); err != nil {
				return nil, err
			}
		default:
			if val, err = strconv.ParseFloat(inputStr(), targetType.Bits()); err != nil {
				return nil, err
			}
		}
		retVal := reflect.New(targetType).Elem()
		retVal.SetFloat(val)
		return retVal.Interface(), nil

	case reflect.Bool:
		var val bool
		switch inputType {
		case reflect.Bool:
			val = reflect.ValueOf(input).Bool()
		case reflect.String:
			if val, err = strconv.ParseBool(input.(string)); err != nil {
				return nil, err
			}
		default:
			if val, err = strconv.ParseBool(inputStr()); err != nil {
				return nil, err
			}
		}
		retVal := reflect.New(targetType).Elem()
		retVal.SetBool(val)
		return retVal.Interface(), nil
	}

	return nil, errors.Join(err, fmt.Errorf("unsupported type: %v", targetType))
}
