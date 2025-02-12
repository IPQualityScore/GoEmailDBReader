package Reader

import (
	"golangemaildbreader/pkg/reader/Types"
	"math/big"
)

type Data struct {
	Hash *big.Int
	Data []Types.TypeInterface
}

func btoi(val []byte) *big.Int {
	var v big.Int
	v.SetBytes(val)
	return &v
}
func (d *Data) GetSize() int64 {
	size := int64(0)
	for _, v := range d.Data {
		size += v.GetSize()
	}
	return 32 + size // 32 bytes for hash + all other data
}

func (d *Data) Deserialize(b []byte) {
	offset := 0
	d.Hash = btoi(b[:32])
	offset += 32
	for _, v := range d.Data {
		v.Deserialize(b[offset:])
		offset += int(v.GetSize())
	}
}
