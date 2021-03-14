package models

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

var (
	r, _ = regexp.Compile("^\\d{4}\\-(0?[1-9]|1[012])\\-(0?[1-9]|[12][0-9]|3[01])$")
)

//easyjson:json
type Request struct {
	Date   string `json:"date" validate:"required"`
	Views  string `json:"views" validate:"required"`
	Clicks string `json:"clicks" validate:"required"`
	Cost   string `json:"cost" validate:"required"`
}

func (req *Request) Validate() error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Date, validation.Match(r)))
}

func ValidateDates(from, to string) error {
	err := validation.Validate(from, validation.Match(r))
	if err != nil {
		return err
	}
	err = validation.Validate(to, validation.Match(r))
	if err != nil {
		return err
	}
	return nil
}

func ValidateSortingParam(sortingParam string) error {
	if sortingParam == "stat_date" || sortingParam == "views" ||
		sortingParam == "clicks" || sortingParam == "cost" ||
		sortingParam == "cpc" || sortingParam == "cpm" ||
		sortingParam == "" {
		return nil
	}
	return errors.New("unexpected sorting param")
}
