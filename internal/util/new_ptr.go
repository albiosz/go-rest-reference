package util

func NewPtr[V any](value V) *V {
	return &value
}
