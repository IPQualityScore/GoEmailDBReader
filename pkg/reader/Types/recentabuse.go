package Types

type RecentAbuse struct {
	RecentAbuse bool
}

func (f *RecentAbuse) ToString() string {
	if f.RecentAbuse {
		return "recent abuse"
	}
	return "No recent abuse"

}

func (f *RecentAbuse) GetID() int {
	return 3
}

func (f *RecentAbuse) GetSize() int64 {
	return 1 //1 byte
}
func (f *RecentAbuse) Deserialize(d []byte) {
	if d[0] == 0x01 {
		f.RecentAbuse = true
	}
	f.RecentAbuse = false

}
