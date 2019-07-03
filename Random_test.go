package main

import (
	"testing"
)

const (
	BASE = 32
	MASK = BASE - 1
	MAX  = 0xFFFFFFFF
)

type BitSet struct {
	Bits []uint32
}

func NewBitSet() *BitSet {
	return &BitSet{
		Bits: make([]uint32, MAX/BASE+1),
	}
}

func (b *BitSet) Set(t int64) {
	b.Bits[t/BASE] |= (1 << uint32(t&MASK))
}

func (b *BitSet) IsExist(t int64) bool {
	if b.Bits[t/BASE]&(1<<(uint32(t&MASK))) != 0 {
		return true
	}
	return false
}

type CheckRep func(*BitSet, []int64) bool

func ElemRepeted(b *BitSet, a []int64) bool {
	for _, item := range a {
		if b.IsExist(item) {
			return true
		} else {
			b.Set(item)
		}
	}
	return false
}

var tests = []struct {
	RandNum int64
	IsExsit CheckRep
}{
	{10000000, ElemRepeted},
	{1, ElemRepeted},
	{0, ElemRepeted},
	{-1, ElemRepeted},
}

func TestRandomIndex(t *testing.T) {
	var (
		Bitmap = &BitSet{}
	)
	for _, test := range tests {
		Bitmap = NewBitSet()
		if test.IsExsit(Bitmap, RandomIndex(test.RandNum)) == false {
			t.Logf("SUCCESS! non-replicated items! \r\n")
		} else {
			t.Errorf("FAILED! has replicated items!\r\n")
		}
	}
}
