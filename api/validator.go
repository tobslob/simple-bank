package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/tobslob/simple-bank/utils"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSurportedCurreny(currency)
	}
	return false
}
