package utils

import "fmt"

// example arg "createDocumentFromFilePath:42"
func ReturnError(callerNameAndLineNumber string, err error) error {
	fmt.Println("uh oh very virus!")
	fmt.Println(callerNameAndLineNumber, err)

	return err
}
