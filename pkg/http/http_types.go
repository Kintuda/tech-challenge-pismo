package http

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

type NotFound struct {
	Message  string `json:"message"`
	Resource string `json:"resource"`
}

const (
	requireTag     = "required"
	uuidTag        = "uuid"
	minSize        = "min"
	documentNumber = "brazilianDocumentNumber"
)

type HttpUnprocessableEntity struct {
	Message string               `json:"message"`
	Errors  []*ValidationDetails `json:"errors"`
}

type ValidationDetails struct {
	Field       string `json:"field"`
	Value       string `json:"value"`
	Constraint  string `json:"constraint"`
	Description string `json:"description"`
}

func FormatMessage(tag string) string {
	switch tag {
	case documentNumber:
		return "invalid document number, should be a valid CPF or CNPJ"
	case uuidTag:
		return "uuid is invalid"
	case minSize:
		return "array should have at least one item"
	case requireTag:
		return "field is required and should not be null nor empty"
	default:
		return ""
	}
}

func NewHttpUnprocessableEntity() *HttpUnprocessableEntity {
	return &HttpUnprocessableEntity{
		Message: "validation error",
		Errors:  make([]*ValidationDetails, 0),
	}
}

func LowerFirstChar(s string) string {
	if len(s) <= 3 {
		return strings.ToLower(s)
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}

func (u *HttpUnprocessableEntity) AddError(vl *ValidationDetails) {
	u.Errors = append(u.Errors, vl)
}

func (u *HttpUnprocessableEntity) AddErrorFromField(errors validator.ValidationErrors) {
	for _, err := range errors {
		var vl ValidationDetails
		vl.Field = LowerFirstChar(strcase.ToSnake(err.Field()))
		vl.Constraint = err.ActualTag()
		vl.Value = fmt.Sprintf("%v", err.Value())
		vl.Description = FormatMessage(err.ActualTag())
		u.Errors = append(u.Errors, &vl)
	}
}
