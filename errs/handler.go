package errs

import (
	"fmt"
)

var (
	flagMissing bool = true
)

type ErrWithDescription interface {
	Error() string
	GetDescription() string
}

type ErrMissingFlag struct {
	Flag string
}

func (e *ErrMissingFlag) Error() string {
	return fmt.Sprintf("missing \"-%s\" flag", e.Flag)
}
func (e *ErrMissingFlag) GetDescription() string {
	return fmt.Sprintf("-%s flag has to be specified", e.Flag)
}

type ErrUnallowableFlagValue struct {
	Flag  string
	Value string
}

func (e *ErrUnallowableFlagValue) Error() string {
	return fmt.Sprintf("unallowable flag value: flag \"%s\", value \"%s\"", e.Flag, e.Value)
}

func (e *ErrUnallowableFlagValue) GetDescription() string {
	if e.Flag == "operation" {
		return fmt.Sprintf("Operation %s not allowed!", e.Value)
	} else {
		return fmt.Sprintf("unallowable flag value: flag \"%s\", value \"%s\"", e.Flag, e.Value)
	}
}

type ErrNotFound struct {
	Id string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("not found item with id = \"%s\"", e.Id)
}

func (e *ErrNotFound) GetDescription() string {
	return fmt.Sprintf("Item with id %s not found", e.Id)
}

type ErrAlreadyExists struct {
	Id string
}

func (e *ErrAlreadyExists) Error() string {
	return fmt.Sprintf("item with id = \"%s\" already exists", e.Id)
}

func (e *ErrAlreadyExists) GetDescription() string {
	return fmt.Sprintf("Item with id %s already exists", e.Id)
}

func ErrorHandler(err error) error {

	if difErr, ok := err.(*ErrMissingFlag); ok {
		return fmt.Errorf(difErr.GetDescription())
	}

	if difErr, ok := err.(*ErrUnallowableFlagValue); ok {
		return fmt.Errorf(difErr.GetDescription())
	}

	if difErr, ok := err.(*ErrNotFound); ok {
		return fmt.Errorf(difErr.GetDescription())
	}

	if difErr, ok := err.(*ErrAlreadyExists); ok {
		return fmt.Errorf(difErr.GetDescription())
	}

	return err
}

func FlagErrors(args map[string]string) error {

	mandatoryFlagsArry := [2]string{"operation", "fileName"}

	for i := 0; i <= 1; i++ {
		for k, v := range args {
			if k == mandatoryFlagsArry[i] && v == "" {
				return ErrorHandler(&ErrMissingFlag{Flag: mandatoryFlagsArry[i]})
			}
		}
	}

	if args["operation"] == "add" {
		if args["item"] == "" {
			return ErrorHandler(&ErrMissingFlag{Flag: "item"})
		}
	}

	if args["operation"] == "findById" || args["operation"] == "remove" {
		if args["id"] == "" {
			return ErrorHandler(&ErrMissingFlag{Flag: "id"})
		}
	}

	// if args["fileName"] != fileName {
	// 	return ErrorHandler(&ErrUnallowableFlagValue{Flag: "fileName", Value: args["fileName"]})
	// }

	return nil
}

func ErrHasTypeNotFound(err error) bool {

	_, ok := err.(*ErrNotFound)

	return ok

}

// func ErrHasTypeAlreadyExists(err error) bool {

// 	_, ok := err.(*ErrAlreadyExists)

// 	return ok

// }
