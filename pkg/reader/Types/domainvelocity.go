package Types

import (
	"fmt"
)

type DomainVelocity struct {
	DomainVelocity byte
}

func (f *DomainVelocity) ToString() string {
	return fmt.Sprintf("DomainVelocity: %x", f.DomainVelocity)
}
func (f *DomainVelocity) GetID() int {
	return 5
}
func (f *DomainVelocity) GetSize() int64 {
	return 1
}
func (f *DomainVelocity) Deserialize(d []byte) {
	f.DomainVelocity = d[0]
}
