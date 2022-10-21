package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
)

func FormatValidationError(err error) []string {
	var resErrors []string

	var jsErr *json.UnmarshalTypeError
	var sqlErr *mysql.MySQLError
	if errors.As(err, &jsErr) {
		fmt.Println("data is invalid", err.Error())
		resErrors = append(resErrors, "Invalid data type")
	} else if errors.As(err, &sqlErr) {
		if strings.Contains(err.Error(), "1062") {
			resErrors = append(resErrors, "Duplicate Data")
		} else {
			resErrors = append(resErrors, err.Error())
		}
		fmt.Println("the sql error", err.Error())
	} else {
		for _, e := range err.(validator.ValidationErrors) {
			resErrors = append(resErrors, fmt.Sprintf("%s is %s", strcase.ToLowerCamel(e.Field()), e.Tag()))
		}

	}
	return resErrors
}
