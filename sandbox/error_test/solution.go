package error_test

import "fmt"

type MyError struct {
	message string
}

func (m *MyError) Error() string {
	return m.message
}

func doThis(s string) (string, error) {
	if s == "Hello" {
		return s, nil
	} else {
		return "", &MyError{message: "something went wrong"}
	}
}

func main() {
	something, err := doThis("Hello")

	if err == nil {
		fmt.Println(something)
	} else {
		fmt.Println(err)
	}
}
