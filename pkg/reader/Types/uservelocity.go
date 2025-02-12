package Types

import (
	"fmt"
)

type UserVelocity struct {
	UserVelocity byte
}

func (f *UserVelocity) ToString() string {
	return fmt.Sprintf("UserVelocity: %x", f.UserVelocity)
}
func (f *UserVelocity) GetID() int {
	return 4
}
func (f *UserVelocity) GetSize() int64 {
	return 1
}
func (f *UserVelocity) Deserialize(d []byte) {
	f.UserVelocity = d[0]
}
