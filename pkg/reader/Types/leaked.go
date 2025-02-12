package Types

type Leaked struct {
	Leaked bool
}

func (f *Leaked) ToString() string {
	if f.Leaked {
		return "leaked"
	}
	return "not leaked"

}

func (f *Leaked) GetID() int {
	return 2
}

func (f *Leaked) GetSize() int64 {
	return 1 //1 byte
}
func (f *Leaked) Deserialize(d []byte) {
	if d[0] == 0x01 {
		f.Leaked = true
	}
	f.Leaked = false

}
