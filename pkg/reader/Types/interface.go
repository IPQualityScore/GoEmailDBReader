package Types

type TypeInterface interface {
	GetID() int
	GetSize() int64
	Deserialize([]byte)
	ToString() string
}
