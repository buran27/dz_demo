package bins

import (
	"encoding/json"
	"time"
)


type Bin struct{
	Id string `json:"Id"`
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name string	`json:"name"`
}


func (bin *Bin) ToBytes() ([]byte, error) {
	acc, err := json.Marshal(bin)
	if err != nil{
		return nil, err
	}
	return acc, nil
}

type BinList struct{
	Bins []Bin
}

func NewBin(id, name string, private bool) *Bin{
	return &Bin{
		Id: id,
		Private: private,
		CreatedAt: time.Now(),
		Name: name,
	}
}

func NewBinList() *BinList{
	return &BinList{
		Bins: []Bin{},
	}
}