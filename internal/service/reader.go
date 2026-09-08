package service

import (
	"fmt"
	"io"
	"os"
)

func FileReader(path string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic Occured: ", err)
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Cannot Open File.")
	}

	defer func (){
		file.Close()
		fmt.Println("File Closed.")
	}()

	fmt.Println("Opening: ", file.Name())
	content, err := io.ReadAll(file)
	if err != nil {
		panic("Error Reading File.")
	}

	fmt.Println(string(content))

	return nil
}