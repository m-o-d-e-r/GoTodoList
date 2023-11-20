package utils

import (
	"os"
	"path"
)

var PROGRAM_HOME string

func CreateProgramFolder() {
	userHome, _ := os.UserHomeDir()
	PROGRAM_HOME = path.Join(userHome, ".go-todo-list")
	os.MkdirAll(PROGRAM_HOME, 0777)
}
