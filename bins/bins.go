package bins

import "time"


type Bin struct{
	Id string
	Private bool
	CreatedAt time.Time
	Name string
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