package Types

type DomainDisposable struct {
	DomainDisposable bool
}

func (f *DomainDisposable) ToString() string {
	if f.DomainDisposable {
		return "disposable"
	}
	return "not disposable"
}
func (f *DomainDisposable) GetID() int {
	return 7
}
func (f *DomainDisposable) GetSize() int64 {
	return 1
}
func (f *DomainDisposable) Deserialize(d []byte) {
	if d[0] == 0x01 {
		f.DomainDisposable = true
		return
	}
	f.DomainDisposable = false
}
