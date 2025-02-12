package Types

import "strconv"

type FraudScore struct {
	FraudScore int
}

func (f *FraudScore) ToString() string {
	return "FraudScore: " + strconv.Itoa(f.FraudScore)
}
func (f *FraudScore) GetID() int {
	return 1
}
func (f *FraudScore) GetSize() int64 {
	return 1 //1 byte for fraudscore
}
func (f *FraudScore) Deserialize(d []byte) {
	f.FraudScore = int(d[0])
}
