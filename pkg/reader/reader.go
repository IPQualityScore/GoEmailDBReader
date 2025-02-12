package Reader

import (
	"fmt"
	"golangemaildbreader/pkg/reader/Types"
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
	leaf := bLeaf[0] == 0x01 //is node a leaf node
	bN := make([]byte, 8)
	r.F.ReadAt(bN, offset)
	offset += int64(len(bN))
	N := int(btoi64(bN)) ///number of keys in the node
	for i := 0; i < N; i++ {
		Key := &Data{}
		Key.Data = []Types.TypeInterface{}
		for _, v := range r.Header.Headers {
			Key.Data = append(Key.Data, reflect.New(reflect.TypeOf(v).Elem()).Interface().(Types.TypeInterface))
		}
		b1 := make([]byte, Key.GetSize())
		r.F.ReadAt(b1, offset)
		Key.Deserialize(b1) //deserialize the key
		offset += Key.GetSize()
		if hash.Cmp(Key.Hash) == 0 { //= found matching key
			return Key
		}
		if hash.Cmp(Key.Hash) == -1 { //< node does not contain the key
			if leaf { //skip further processing if leaf node
				return nil
			}
			offset += int64(8 * (i - 1))
			pos := make([]byte, 8)
			r.F.ReadAt(pos, offset)
			offset = btoi64(pos) //move to the next node
			if offset == 0 {
				return nil
			}
			return r.ContainsOnOffset(hash, offset) //recursive call to search the next node
		}
	}
	return nil
}

// cleanup function to close the file
func (r *Reader) Close() {
	err := r.F.Close()
	if err != nil {
		fmt.Println(err)
	}
}
