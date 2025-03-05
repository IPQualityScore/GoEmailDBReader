package Types

import (
	"time"
)

type DomainAge struct {
	DomainAge time.Time
}

func (f *DomainAge) ToString() string {
	return "DomainAge: " + f.DomainAge.String()
}
func (f *DomainAge) GetID() int {
	return 9
}
func (f *DomainAge) GetSize() int64 {
	return 8
}
func (f *DomainAge) Deserialize(d []byte) {
	var r int64
	for i := int64(0); i < 8; i++ {
		r |= int64(d[i]) << (i * 8)
	}
	f.DomainAge = time.Unix(r, 0).UTC()
}
