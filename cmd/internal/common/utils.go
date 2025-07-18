package common

import (
	"log"
	"os"
)

func Exist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		log.Panic(err)
		return false
	}
}
