package files

import (
	"fmt"
	"os"
)

func ReadFile (name string) ([]byte, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil

}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Файл успешно записан")
}
