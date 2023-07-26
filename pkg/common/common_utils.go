package common

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)


//This function use-case is to prompt with some text
//and wait for user input
func Prompt(promptText string) string {
	var input string
	fmt.Print(promptText)
	fmt.Scanln(&input)
	return input
}

//For appending password entries in password files
func ReadEntriesFromFile(fileName string) ([]PasswordEntry, error) {
	// Read the existing entries from the file, if any
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	// Unmarshal the file data into a slice of PasswordEntry structs
	var entries []PasswordEntry
	err = json.Unmarshal(fileData, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

//This function is to hash the password
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

//This fucntion is to get data from Application name or Username
func GetPasswordEntry(fileName string, username string, applicationName string) ([]PasswordEntry, error){
	// Read the existing entries from the file
	var entries []PasswordEntry
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return entries, err
	}
	// Unmarshal the file data into a slice of PasswordEntry structs
	err = json.Unmarshal(fileData, &entries)
	if err != nil {
		return entries, err
	}
	return entries, nil
}