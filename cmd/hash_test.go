package cmd

import (
	"fmt"
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
