package main

import (
	"flag"
	"homework_8/errs"
	"homework_8/services"
	"homework_8/storage"
	"io"
	"os"
)

var (
	nameOfFile string
	operation  string
	item       string
	id         string
)

type Arguments map[string]string

func main() {

	flag.StringVar(&nameOfFile, "fileName", "", "file name")
	flag.StringVar(&operation, "operation", "", "operation : add, list, findById, remove")
	flag.StringVar(&item, "item", "", "item")
	flag.StringVar(&id, "id", "", "id")
	flag.Parse()

	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {

	args := Arguments{
		"id":        id,
		"operation": operation,
		"item":      item,
		"fileName":  nameOfFile,
	}

	return args
}

func Perform(args Arguments, writer io.Writer) error {

	err := errs.FlagErrors(args)

	if err != nil {
		return err
	}

	storage := storage.NewStorage(args["fileName"])
	handler := services.NewHandler(storage)

	switch args["operation"] {
	case "add":
		err = handler.Add(args["item"], writer)
	case "list":
		err = handler.GetList(writer)
	case "findById":
		err = handler.GetUserById(args["id"], writer)
	case "remove":
		err = handler.Remove(args["id"], writer)
	default:

		err = &errs.ErrUnallowableFlagValue{Flag: "operation", Value: args["operation"]}
	}

	if err != nil {
		return errs.ErrorHandler(err)
	}

	return nil
}
