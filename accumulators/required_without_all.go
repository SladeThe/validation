package accumulators

import (
	"github.com/SladeThe/yav"
)

type RequiredWithoutAll[T any] struct {
	provideFunc ProvideFunc[T]
	enabled     bool
}

func NewRequiredWithoutAll[T any](provideFunc ProvideFunc[T]) RequiredWithoutAll[T] {
	return RequiredWithoutAll[T]{
		provideFunc: provideFunc,
		enabled:     true,
	}
}

func (r RequiredWithoutAll[T]) String(value string) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == ""
	return r
}

func (r RequiredWithoutAll[T]) Bytes(value []byte) RequiredWithoutAll[T] {
	r.enabled = r.enabled && len(value) == 0
	return r
}

func (r RequiredWithoutAll[T]) Bool(value bool) RequiredWithoutAll[T] {
	r.enabled = r.enabled && !value
	return r
}

func (r RequiredWithoutAll[T]) Int(value int) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == 0
	return r
}

func (r RequiredWithoutAll[T]) Int8(value int8) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == 0
	return r
}

func (r RequiredWithoutAll[T]) Int16(value int16) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == 0
	return r
}

func (r RequiredWithoutAll[T]) Int32(value int32) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == 0
	return r
}

func (r RequiredWithoutAll[T]) Int64(value int64) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == 0
	return r
}

func (r RequiredWithoutAll[T]) Uint(value int) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == 0
	return r
}

func (r RequiredWithoutAll[T]) Uint8(value uint8) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == 0
	return r
}

func (r RequiredWithoutAll[T]) Uint16(value uint16) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == 0
	return r
}

func (r RequiredWithoutAll[T]) Uint32(value uint32) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == 0
	return r
}

func (r RequiredWithoutAll[T]) Uint64(value uint64) RequiredWithoutAll[T] {
	r.enabled = r.enabled && value == 0
	return r
}

func (r RequiredWithoutAll[T]) Fields(fields string) yav.ValidateFunc[T] {
	return r.provideFunc(fields, r.enabled)
}
