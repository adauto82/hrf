package cmd

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHashFunction(t *testing.T) {
	data := []byte{12}
	res := Hash(data)
	expected := []byte{24, 108, 90, 204, 81, 189, 102, 126}

	for i, e := range expected {
		if e != res[i] {
			t.Error(fmt.Sprintf("%d index should be %d instead of %d",i,e,res[i]))
			return
		}
	}
}

func TestHashFunctionIgnoresAllTheBytesButLast(t *testing.T) {
	l := rand.Intn(30) + 6
	dataUnknown := make([]byte, l)
	rand.Read(dataUnknown)
	dataUnknown[len(dataUnknown)-1] = 12
	dataKnown := []byte{12}
	resK := Hash(dataKnown)
	resU := Hash(dataUnknown)
	//expected := []byte{24, 108, 90, 204, 81, 189, 102, 126}
	for i, _ := range resK {
		if resU[i] != resK[i] {
			t.Error(fmt.Sprintf("%d from resU is different from %d from resK at index i: %d",resU[i],resK[i],i))
			return
		}
	}
}
