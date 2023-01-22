// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package vnumber

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestRequiredWithAnyInt(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      int
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyInt()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyInt8(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      int8
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyInt8()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyInt16(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      int16
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyInt16()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyInt32(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      int32
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyInt32()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyInt64(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      int64
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyInt64()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyUint(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      uint
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyUint()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyUint8(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      uint8
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyUint8()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyUint16(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      uint16
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyUint16()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyUint32(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      uint32
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyUint32()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyUint64(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      uint64
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyUint64()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyFloat32(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      float32
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyFloat32()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestRequiredWithAnyFloat64(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      float64
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAnyFloat64()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
