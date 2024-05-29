package utils

import (
	"github.com/leaanthony/spinner"
)

var Spinner *spinner.Spinner

func GenerateSpinner() {
	Spinner = spinner.New("")
}

func StartSpinner(message string) {
	Spinner.Start(message)
}

func ErrorSpinner(message string) {
	Spinner.Error(message)
}

func ErrorfSpinner(format string, args ...interface{}) {
	Spinner.Errorf(format, args)
}

func SuccessSpinner(message string) {
	Spinner.Success(message)
}

func SuccessfSpinner(format string, args ...interface{}) {
	Spinner.Successf(format, args)
}
