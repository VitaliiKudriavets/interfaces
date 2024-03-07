package main

import "fmt"

type AppError struct {
	Err    error
	Custom string
	Field  int
}

func (a *AppError) Error() string {
	fmt.Println(a.Custom, a.Field)
	return a.Err.Error()
}

func main() {
	err := m()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("success")
}

func m() error {
	return &AppError{
		Err:    fmt.Errorf("my new error"),
		Custom: "value",
		Field:  5,
	}
}
