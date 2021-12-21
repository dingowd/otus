package hw09structvalidator

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	tagName = "validate"
)

type Analized int

const (
	aUncheck Analized = iota
	aInt
	aIntslice
	aString
	aStringslice
)

var (
	ValidatingFieldError  = errors.New("field validation error")
	IncorrectTag          = errors.New("incorrect tag")
	ValidatingStructError = errors.New("structure validation error")
)

type Phone string

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

var ValidErr ValidationErrors

func (v ValidationErrors) Error() string {
	for _, val := range v {
		if val.Err != nil {
			return "error validating structure"
		}
	}
	return "success"
}

func AnalizeField(elem reflect.Value) Analized {
	switch elem.Kind() {
	case reflect.Slice:
		if elem.IsNil() {
			return aUncheck
		}
		if reflect.ValueOf(elem.Index(0).Interface()).Kind() == reflect.Int {
			return aIntslice
		} else if reflect.ValueOf(elem.Index(0).Interface()).Kind() == reflect.String {
			return aStringslice
		}
		return aUncheck
	case reflect.Int:
		return aInt
	case reflect.String:
		return aString
	default:
		return aUncheck
	}
}

func ValidateIntSLice(e reflect.Value, tags []string) error {
	var arr []int
	switch e.Kind() {
	case reflect.Int:
		arr = append(arr, int(e.Int()))
	case reflect.Slice:
		for i := 0; i < e.Len(); i++ {
			h := e.Index(i).Interface()
			hr := reflect.ValueOf(h)
			arr = append(arr, int(hr.Int()))
		}
	default:
	}
	for _, val := range arr {
		for _, tag := range tags {
			tt := strings.Split(tag, ":")
			switch tt[0] {
			case "min":
				min, err := strconv.Atoi(tt[1])
				if err != nil {
					return IncorrectTag
				}
				if val < min {
					return ValidatingFieldError
				}
			case "max":
				max, err := strconv.Atoi(tt[1])
				if err != nil {
					return IncorrectTag
				}
				if val > max {
					return ValidatingFieldError
				}

			case "in":
				strIn := strings.Split(tt[1], ",")
				err := ValidatingFieldError
				for _, tagStr := range strIn {
					intStr, e := strconv.Atoi(tagStr)
					if e != nil {
						return IncorrectTag
					}
					if val == intStr {
						err = nil
						return err
					}
				}
				return err
			}
		}
	}
	return nil
}

func ValidateStringSLice(e reflect.Value, tags []string) (ve error) {
	ve = nil
	var arr []string
	switch e.Kind() {
	case reflect.String:
		arr = append(arr, e.String())
	case reflect.Slice:
		for i := 0; i < e.Len(); i++ {
			h := e.Index(i).Interface()
			hr := reflect.ValueOf(h)
			arr = append(arr, hr.String())
		}
	}
	for _, val := range arr {
		for _, tag := range tags {
			tt := strings.Split(tag, ":")
			switch tt[0] {
			case "len":
				l, err := strconv.Atoi(tt[1])
				if err != nil {
					return IncorrectTag
				}
				if len(val) != l {
					return ValidatingFieldError
				}
			case "in":
				strIn := strings.Split(tt[1], ",")
				err := ValidatingFieldError
				for _, tagStr := range strIn {
					if val == tagStr {
						err = nil
						return err
					}
				}
				return err
			case "regexp":
				re := regexp.MustCompile(tt[1])
				if !re.MatchString(val) {
					return ValidatingFieldError
				}
			}
		}
	}
	return ve
}

func Validate(user interface{}) error {
	var vh ValidationError
	var ve error
	ve = nil
	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)
	for i := 0; i < v.NumField(); i++ {
		e := v.Field(i)
		field := t.Field(i)
		tags := strings.Split(field.Tag.Get(tagName), "|")
		AF := AnalizeField(e)
		if AF == aInt || AF == aIntslice {
			AF = aInt
		}
		if AF == aString || AF == aStringslice {
			AF = aString
		}
		switch AF {
		case aInt:
			vh.Field = field.Name
			vh.Err = ValidateIntSLice(e, tags)
			if vh.Err != nil {
				ve = ValidatingStructError
			}
			ValidErr = append(ValidErr, vh)
		case aString:
			vh.Field = field.Name
			vh.Err = ValidateStringSLice(e, tags)
			ValidErr = append(ValidErr, vh)
			if vh.Err != nil {
				ve = ValidatingStructError
			}
		}
	}
	return ve
}
