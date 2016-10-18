package deal

import (
	"log"
	"os"
)

func CheckFile() {
	//path := ".i"

	finfo, err := os.Stat(".i")
	if err != nil {
		// no such file or dir
		log.Println("not exist")
		return
	}
	if finfo.IsDir() {
		log.Println("not a file")
		// it's a file
	} else {
		log.Println("is a file")
		// it's a directory
	}
}
