package cmd

/**
There must be something fishy about this algorithm that I am not understanding and because of lack of direct communication
I can't put this to the test. But as Far as I can understand, it does not matter how large or small is the []byte passed.
It will all be decided by the last byte of the []byte. I will implement this as specified but leave this comment here so
it is noted and in the hopes that someone maybe wants to clarify this for me. Thanks. I try to leave my point coded the tests also.

This is the information I got from Jeff about the hash function:

IMT Hash description:
Length: 8 bytes
coefficients := [8]int{ 2, 3, 5, 7, 11, 13, 17, 19 }
for each incoming byte, ib:
        for each byte of the hash, h
               h[i] = ((h[i-1] + ib) * coefficient[i]) % 255
               // in the case where i-1 == -1, h[i-1] should be 0.



For example, hashing the data:
data := []byte{12}

Should result in a hash of:
[]byte{24, 108, 90, 204, 81, 189, 102, 126}

When converted to hexadecimal for writing to the output file:
186c5acc51bd667e
*/

//IMT Hash Function
var COEFFICIENTS = [8]int{2, 3, 5, 7, 11, 13, 17, 19}

//Indexed Merkle Tree hash function
func Hash(bytes []byte) []byte {
	hash := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	for _, ib := range bytes {
		for i := range hash {
			if i == 0 {
				// in the case where i-1 == -1, h[i-1] should be 0.
				hash[0] = byte((int(ib) * COEFFICIENTS[0]) % 255)
			} else {
				hash[i] = byte(((int(hash[i-1]) + int(ib)) * COEFFICIENTS[i]) % 255)
			}
		}
	}
	return hash
}
