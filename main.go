package main

import "time"

type Bin struct{
	id string
	privat bool
	createdAt time.Time
	name string
}

type BinList struct{
	bins []Bin
}

func NewBin(id, name string, privat bool) *Bin{
	return &Bin{
		id: id,
		privat: privat,
		createdAt: time.Now(),
		name: name,
	}
}

func NewBinList() *BinList{
	return &BinList{
		bins: []Bin{},
	}
}