package storage

import (
	"fmt"
	"os"
)



func NewStorage(){}


func Read(name string) ([]byte, error){
	file, err := os.ReadFile(name)
	if err != nil{
		return nil, err
	}
	return file, nil
}


func Write(name string, content []byte){
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
}