package Reader

import (
	"fmt"
	"github.com/IPQualityScore/GoEmailDBReader/pkg/reader/Types"
	"math/big"
	"os"
	"reflect"
)

const (
	t = 64
)

// F should be a pointer to a tree file, Reader.Header.Deserialize will init the header variable
type Reader struct {
	Header Header
	F      *os.File
}

// Helper function to convert byte array to int64
func btoi64(val []byte) int64 {
	var r int64
	for i := int64(0); i < 8; i++ {
		r |= int64(val[i]) << (i * 8)
	}
	return r
}

// Recursive function to search for a hash in the tree
// the first call should have offset at the end of the header
func (r *Reader) ContainsOnOffset(hash *big.Int, offset int64) *Data {
	bLeaf := make([]byte, 1)
	r.F.ReadAt(bLeaf, offset)
	offset++

	leaf := bLeaf[0] == 0x01
	bN := make([]byte, 8)
	r.F.ReadAt(bN, offset)
	offset += 8
	N := int(btoi64(bN))
	i := 0
	for i = 0; i < N; i++ {
		Key := &Data{}
		Key.Data = []Types.TypeInterface{}
		for _, v := range r.Header.Headers {
			Key.Data = append(Key.Data, reflect.New(reflect.TypeOf(v).Elem()).Interface().(Types.TypeInterface))
		}
		b1 := make([]byte, Key.GetSize())
		r.F.ReadAt(b1, offset)
		Key.Deserialize(b1)
		offset += Key.GetSize()
		compare := hash.Cmp(Key.Hash)
		if compare == 1 { //>
			continue
		}
		if compare == 0 { //=
			return Key
		}
		offset += Key.GetSize() * int64(N-i-1)
		break

	}
	if leaf {
		return nil
	}
	offset += int64(8 * i)
	pos := make([]byte, 8)
	r.F.ReadAt(pos, offset)
	offset = btoi64(pos)
	if offset == 0 {
		return nil
	}
	return r.ContainsOnOffset(hash, offset)

}

// cleanup function to close the file
func (r *Reader) Close() {
	err := r.F.Close()
	if err != nil {
		fmt.Println(err)
	}
}
