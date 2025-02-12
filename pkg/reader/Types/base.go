package Types

import "fmt"

type Base struct {
	Valid          bool
	Disposable     bool
	Suspect        bool
	CatchAll       bool
	SmtpScore      byte
	Deliverability byte
}

func (b *Base) GetID() int {
	return 0
}
func (b *Base) GetSize() int64 {
	return 1 + 1 + 1 //1 byte for valid,diposable,suspect,catchall,1 byte for smtpscore, 1 byte for deliverability
}

func (b *Base) Deserialize(d []byte) {
	if ((d[0] & (1 << 0)) >> 0) == 1 {
		b.Valid = true
	} else {
		b.Valid = false
	}
	if ((d[0] & (1 << 1)) >> 1) == 1 {
		b.Disposable = true
	} else {
		b.Disposable = false
	}

	if ((d[0] & (1 << 2)) >> 2) == 1 {
		b.Suspect = true
	} else {
		b.Suspect = false
	}
	if ((d[0] & (1 << 3)) >> 3) == 1 {
		b.CatchAll = true
	} else {
		b.CatchAll = false
	}
	b.SmtpScore = d[1]
	b.Deliverability = d[2]
}
func (b *Base) ToString() string {
	return fmt.Sprintf("Valid: %t, Disposable: %t, Suspect: %t, CatchAll: %t, SmtpScore: %d, Deliverability: %d", b.Valid, b.Disposable, b.Suspect, b.CatchAll, b.SmtpScore, b.Deliverability)
}
