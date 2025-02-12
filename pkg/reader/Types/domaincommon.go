package Types

type DomainCommon struct {
	DomainCommon bool
}

func (f *DomainCommon) ToString() string {
	if f.DomainCommon {
		return "Domain is common"
	}
	return "Domain is not commonm"
}
func (f *DomainCommon) GetID() int {
	return 6
}
func (f *DomainCommon) GetSize() int64 {
	return 1
}
func (f *DomainCommon) Deserialize(d []byte) {
	if d[0] == 0x01 {
		f.DomainCommon = true
		return
	}
	f.DomainCommon = false

}
