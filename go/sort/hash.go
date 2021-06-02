package sort

import (
	"math"
)

//	"github.com/cespare/xxhash"

/*
func xxhash64(data Interface) uint64 {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		log.Println("binary.Write failed:", err)
	}
	return xxhash.Sum64(buf.Bytes())
}
*/

// N element to run countHash
var N int = 3

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
Hash function to identify bucket number form element. Customized
to properly encode elements in order within the buckets
*/
func countHash(a interface{}) uint64 {
	var hc uint64 = 0

	switch a.(type) {
	case int:
		hc = uint64(a.(int))
	case string:
		data := []byte(a.(string))
		mask := make([]uint64, N)
		for i := 0; i < N; i++ {
			mask[i] = uint64(math.Pow(26.0, float64(N-1-i)))
		}
		//mask := []int{676, 26, 1}
		for i := 0; i < min(len(data), N); i++ {
			hc += uint64(data[i]-'a') * mask[i]
		}
	}
	return hc
}
