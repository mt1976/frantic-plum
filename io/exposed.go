package io

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mt1976/frantic-plum/commonErrors"
	"github.com/mt1976/frantic-plum/logger"
)

// The Copy function copies a file from one path to another and returns true if the copy was
// successful.
func Copy(fileName string, fromPath string, toPath string) bool {

	logger.WarningLogger.Println("Copying " + fileName + " from " + fromPath + " to " + toPath)

	content, err := Read(fileName, fromPath)
	if err != nil {
		logger.ErrorLogger.Fatalf("File Read Error %v", err)
	}

	ok, err2 := Write(fileName, toPath, content)
	if err2 != nil {
		logger.ErrorLogger.Fatalf("File Write Error %v", err2)
	}

	if !ok {
		logger.ErrorLogger.Fatalf("Unable to Copy "+fileName+" from "+fromPath+" to "+toPath, nil)
	}

	return true
}

// The Read function reads the content of a file given its name and path, and returns the content as a
// string.
func Read(fileName string, path string) (string, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}

	// Check it exists - If not create it
	if !Touch(filePath) {
		WriteData(fileName, path, "")
	}

	//log.Println("Read          :", filePath)
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.ErrorLogger.Fatal("Read Error : [", err, "]")
	}
	// Convert []byte to string and print to screen
	return string(content), commonErrors.ReadError(err)
}

// The Write function writes content to a file specified by fileName and path, and returns a boolean
// indicating success and an error if any.
func Write(fileName string, path string, content string) (bool, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Write         :", filePath)

	message := []byte(content)
	err := ioutil.WriteFile(filePath, message, 0644)
	if err != nil {
		logger.ErrorLogger.Fatalf("Write Error : [%v]", err)
		return false, commonErrors.WriteError(err)
	}
	return false, nil
}

// The function `WriteData` writes the given content to a file with the specified name and path.
func WriteData(fileName string, path string, content string) int {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Write         :", filePath)

	message := []byte(content)
	err := os.WriteFile(filePath, message, 0644)
	if err != nil {
		logger.ErrorLogger.Fatalf("Write Error %v", err)
		return -1
	}

	//	log.Println("File Write : " + fileName + " in " + path + "[" + filePath + "]")
	logger.InfoLogger.Panicln(fileName, filePath)
	return 1
}

// Touch returns true if the specified file existing on the filesystem
// The Touch function takes a filename as input and returns a boolean value indicating whether the file
// was successfully touched.
func Touch(filename string) bool {
	return touch(filename)
}

// Empty clears the contents of a specified directory
// The function "Empty" deletes all files in a given directory.
func Empty(dir string) error {
	logger.InfoLogger.Println("TRASH", dir)
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		logger.InfoLogger.Println(err)
		return commonErrors.EmptyError(err)
	}
	//	fmt.Println("do Clear", files)
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			logger.InfoLogger.Println(err)
			return commonErrors.ClearError(err)
		}
	}
	return nil
}
