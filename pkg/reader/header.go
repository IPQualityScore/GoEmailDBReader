package Reader

import (
	"errors"
	"fmt"
	"golangemaildbreader/pkg/reader/Types"
	"os"
)

type Header struct {
	Header  [4]byte
	Version byte
	// Type 0x00 for email 0x01 for domain
	Type         byte
	CreationTime uint64
	Headers      []Types.TypeInterface
}

func (h *Header) GetSize() int64 {
	return int64(4 + 1 + 1 + 8 + 1 + len(h.Headers)*2)
}

func (h *Header) Deserialize(file *os.File) error {
	buf := []byte{0, 0, 0, 0}
	file.ReadAt(buf, 0)
	if buf[0] != 'I' || buf[1] != 'P' || buf[2] != 'Q' || buf[3] != 'S' {
		return errors.New("invalid header")
	}
	h.Header = [4]byte{buf[0], buf[1], buf[2], buf[3]}
	buf = []byte{0}
	file.ReadAt(buf, 4)
	if buf[0] != 0x01 {
		return errors.New(fmt.Sprintf("invalid version, expected 0x01, got %x", buf[0]))
	}
	h.Version = buf[0]
	buf = []byte{0}
	file.ReadAt(buf, 5)
	h.Type = buf[0]
	buf = []byte{0, 0, 0, 0, 0, 0, 0, 0}
	file.ReadAt(buf, 6)
	h.CreationTime = uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 | uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56
	buf = []byte{0}
	file.ReadAt(buf, 14)
	count := int(buf[0])
	h.Headers = make([]Types.TypeInterface, 0)
	possibleHeaders := []Types.TypeInterface{&Types.Base{}, &Types.FraudScore{}, &Types.Leaked{}, &Types.RecentAbuse{}, &Types.UserVelocity{}, &Types.DomainVelocity{}, &Types.DomainCommon{}, &Types.DomainDisposable{}}
	offset := 14
	for i := 0; i < count; i++ {
		buf = []byte{0, 0}
		file.ReadAt(buf, int64(1+offset))
		id := int(buf[0])
		size := int64(buf[1])
		found := false
		for _, possibleHeader := range possibleHeaders {
			if possibleHeader.GetID() == id {
				h.Headers = append(h.Headers, possibleHeader)
				found = true
				break
			}
		}
		if !found {
			h.Headers = append(h.Headers, &Types.Unknown{Size: int64(size)})
		}
		offset += 2
	}
	return nil
}
