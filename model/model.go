package model

const (
	Ver = "v1.0.0"
)

type Action struct {
	Key  Value
	Type Value
	From Value
}

type Value struct {
	ID      int
	Message string
	Value   string
}
