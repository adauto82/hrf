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

func TestHashFunction2(t *testing.T) {
	data := []byte{12,3,3,12,4,56,4}
	res := Hash(data)
	//expected := []byte{24, 108, 90, 204, 81, 189, 102, 126}
	fmt.Println("{")
	for i, e := range res {

		fmt.Println(fmt.Sprintf(" %d: %d",i,e))



		//if e != res[i] {
		//	t.Error(fmt.Sprintf("%d index should be %d instead of %d",i,e,res[i]))
		//	return
		//}
	}
	fmt.Println("}")
}
