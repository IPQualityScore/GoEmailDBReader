package Types

import (
	"time"
)

type FirstSeen struct {
	FirstSeen time.Time
}

func (f *FirstSeen) ToString() string {
	return "FirstSeen: " + f.FirstSeen.String()
}
func (f *FirstSeen) GetID() int {
	return 8
}
func (f *FirstSeen) GetSize() int64 {
	return 8
}
func (f *FirstSeen) Deserialize(d []byte) {
	var r int64
	for i := int64(0); i < 8; i++ {
		r |= int64(d[i]) << (i * 8)
	}
	f.FirstSeen = time.Unix(r, 0).UTC()
}
