package Types

type Unknown struct {
	Size int64
}

func (b *Unknown) ToString() string {
	return string("")
}

func (b *Unknown) GetID() int {
	return 255
}

func (f *Unknown) GetSize() int64 {
	return int64(f.Size)
}
func (f *Unknown) Deserialize(d []byte) {

}
