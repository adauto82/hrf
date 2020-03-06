package cmd

//IMT Hash Function
var COEFFICIENTS = [8]int{ 2, 3, 5, 7, 11, 13, 17, 19 }

//Indexed Merkle Tree hash function
func Hash(bytes []byte) []byte {
	hash := []byte{0,0,0,0,0,0,0,0}
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