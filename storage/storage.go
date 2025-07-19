package storage

import (
	"encoding/json"
	"fmt"
	"mineAPI/bins"
	"os"
)

func Read(name string) (*bins.Bin, error) {
	var box *bins.Bin
	file, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &box)
	if err != nil {
		return nil, err
	}
	return box, nil
}

func Write(name string, bin *bins.Bin) {
	acc, err := json.Marshal(bin)
	if err != nil {
		fmt.Println(err.Error())
	}
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	_, err = file.Write(acc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
