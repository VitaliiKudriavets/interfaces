package main

import (
	"fmt"
	"io"
	"strings"
)

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
	r := strings.NewReader("Hello world")
	arr := make([]byte, 8)
	for {
		n, err := r.Read(arr)
		fmt.Printf("n = %d err = %v b = %v\n", n, err, arr)
		fmt.Printf("arr n bytes: %q \n", arr[:n])
		if err == io.EOF {
			break
		}
	}
	// Error
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
