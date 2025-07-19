package main

import (
	"fmt"
	"mineAPI/bins"
	"mineAPI/storage"
)

func main() {
	bin := bins.NewBin("1", "new", false)

	storage.Write("file.json", bin)
	file, err := storage.Read("file.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(*file)
}
