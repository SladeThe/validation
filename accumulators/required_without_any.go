package accumulators

import (
	"github.com/SladeThe/yav"
)

type RequiredWithoutAny[T any] struct {
	provideFunc ProvideFunc[T]
	enabled     bool
}

func NewRequiredWithoutAny[T any](provideFunc ProvideFunc[T]) RequiredWithoutAny[T] {
	return RequiredWithoutAny[T]{
		provideFunc: provideFunc,
		enabled:     false,
	}
}

func (r RequiredWithoutAny[T]) String(value string) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == ""
	return r
}

func (r RequiredWithoutAny[T]) Bytes(value []byte) RequiredWithoutAny[T] {
	r.enabled = r.enabled || len(value) == 0
	return r
}

func (r RequiredWithoutAny[T]) Bool(value bool) RequiredWithoutAny[T] {
	r.enabled = r.enabled || !value
	return r
}

func (r RequiredWithoutAny[T]) Int(value int) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Int8(value int8) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Int16(value int16) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Int32(value int32) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Int64(value int64) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Uint(value int) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Uint8(value uint8) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Uint16(value uint16) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Uint32(value uint32) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Uint64(value uint64) RequiredWithoutAny[T] {
	r.enabled = r.enabled || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Fields(fields string) yav.ValidateFunc[T] {
	return r.provideFunc(fields, r.enabled)
}
