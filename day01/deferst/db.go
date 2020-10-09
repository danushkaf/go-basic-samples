package deferst

import (
	"errors"
	"fmt"
)

func connectToDB() (bool, error) {
	fmt.Println("ok, connected to db")
	return false, errors.New("error: This is a dummy error")
}

func disconnectFromDB() error {
	fmt.Println("ok, disconnected from db")
	return nil
}

// DoDBOperations is to explain defer statements.
func DoDBOperations() {
	_, err := connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	if err != nil {
		fmt.Println("Oops! some crash or network error ...")
		fmt.Println("Returning from function here!")
	}
	// deferred function executed here just before actually returning, even if there is a return or abnormal termination before
}
