package handler

import (
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"reflect"
	"regexp"
	"strings"
)

const tagName = "validate"

var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

type Validator interface {
	Validate(interface{}) error
}

type DefaultValidator struct {
}

func (v DefaultValidator) Validate(val interface{}) error {
	return nil
}

type EmailValidator struct {
}

func (v EmailValidator) Validate(val interface{}) error {
	if !mailRe.MatchString(val.(string)) {
		return fmt.Errorf("emailValidator: it is not a valid email address")
	}
	return nil
}

type PasswordValidator struct {
}

//should be 8 characters long
func (v PasswordValidator) Validate(val interface{}) error {
	if len(val.(string)) != 8 && len(val.(string)) != 0 {
		return fmt.Errorf("passwordValidator: %d", len(val.(string)))
	}
	if len(val.(string)) == 0 {
		return nil
	}
	for _, i := range val.(string) {
		var overlap = false
		for _, j := range model.PasswordComposition {
			if i == j {
				overlap = true
				break
			}
		}
		if overlap == false {
			return fmt.Errorf("passwordValidator: the password must consist of uppercase and lowercase letters and numbers")
		} else {
			overlap = false
		}
	}
	return nil
}

func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")
	switch args[0] {
	case "email":
		return EmailValidator{}
	case "password":
		return PasswordValidator{}
	}
	return DefaultValidator{}
}

func validateStruct(s interface{}) map[string]string {
	var errs = make(map[string]string)
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}
		validator := getValidatorFromTag(tag)
		err := validator.Validate(v.Field(i).Interface())
		if err != nil {
			errs[v.Type().Field(i).Name] = err.Error()
		}
	}
	return errs
}
