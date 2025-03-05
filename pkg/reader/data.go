package Reader

import (
	"github.com/IPQualityScore/GoEmailDBReader/pkg/reader/Types"
	"math/big"
)

type Data struct {
	Hash *big.Int
	Data []Types.TypeInterface
}

func btoi(val []byte) *big.Int {
	var v big.Int
	v.SetBytes(val)
	return &v
}
func (d *Data) GetSize() int64 {
	size := int64(0)
	for _, v := range d.Data {
		size += v.GetSize()
	}
	return 32 + size // 32 bytes for hash + all other data
}

func (d *Data) Deserialize(b []byte) {
	offset := 0
	d.Hash = btoi(b[:32])
	offset += 32
	for _, v := range d.Data {
		v.Deserialize(b[offset:])
		offset += int(v.GetSize())
	}
}
func (d *Data) Base() *Types.Base {
	for _, v := range d.Data {
		switch v.(type) {
		case *Types.Base:
			return v.(*Types.Base)
		}
	}
	return nil
}
func (d *Data) DomainCommon() *Types.DomainCommon {
	for _, v := range d.Data {
		switch v.(type) {
		case *Types.DomainCommon:
			return v.(*Types.DomainCommon)
		}
	}
	return nil
}

func (d *Data) DomainDisposable() *Types.DomainDisposable {
	for _, v := range d.Data {
		switch v.(type) {
		case *Types.DomainDisposable:
			return v.(*Types.DomainDisposable)
		}
	}
	return nil
}

func (d *Data) DomainVelocity() *Types.DomainVelocity {
	for _, v := range d.Data {
		switch v.(type) {
		case *Types.DomainVelocity:
			return v.(*Types.DomainVelocity)
		}
	}
	return nil
}

func (d *Data) FraudScore() *Types.FraudScore {
	for _, v := range d.Data {
		switch v.(type) {
		case *Types.FraudScore:
			return v.(*Types.FraudScore)
		}
	}
	return nil
}

func (d *Data) Leaked() *Types.Leaked {
	for _, v := range d.Data {
		switch v.(type) {
		case *Types.Leaked:
			return v.(*Types.Leaked)
		}
	}
	return nil
}

func (d *Data) RecentAbuse() *Types.RecentAbuse {
	for _, v := range d.Data {
		switch v.(type) {
		case *Types.RecentAbuse:
			return v.(*Types.RecentAbuse)
		}
	}
	return nil
}

func (d *Data) UserVelocity() *Types.UserVelocity {
	for _, v := range d.Data {
		switch v.(type) {
		case *Types.UserVelocity:
			return v.(*Types.UserVelocity)
		}
	}
	return nil
}
func (d *Data) FirstSeen() *Types.FirstSeen {
	for _, v := range d.Data {
		switch v.(type) {
		case *Types.FirstSeen:
			return v.(*Types.FirstSeen)
		}
	}
	return nil
}

func (d *Data) DomainAge() *Types.DomainAge {
	for _, v := range d.Data {
		switch v.(type) {
		case *Types.DomainAge:
			return v.(*Types.DomainAge)
		}
	}
	return nil
}
