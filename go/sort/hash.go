package sort

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
func countHash(a interface{}) uint64 {
	var hc uint64 = 0

	switch a.(type) {
	case int:
		hc = uint64(a.(int))
	case string:
		data := []byte(a.(string))
		for i := 0; i < len(data); i++ {
			hc += uint64(data[i] - '0')
		}
	}
	return hc
}
