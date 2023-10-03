package vduration

import (
	"time"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	excludedWithAnyFuncs    map[string]yav.ValidateFunc[time.Duration]
	excludedWithoutAnyFuncs map[string]yav.ValidateFunc[time.Duration]
	excludedWithAllFuncs    map[string]yav.ValidateFunc[time.Duration]
	excludedWithoutAllFuncs map[string]yav.ValidateFunc[time.Duration]
)

func ExcludedIf(conditionString string, condition bool) yav.ValidateFunc[time.Duration] {
	if !condition {
		return yav.Next[time.Duration]
	}

	return excludedIf(conditionString).validate
}

func ExcludedUnless(conditionString string, condition bool) yav.ValidateFunc[time.Duration] {
	if condition {
		return yav.Next[time.Duration]
	}

	return excludedUnless(conditionString).validate
}

func ExcludedWithAny() accumulators.ExcludedWithAny[time.Duration] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAny)
}

func ExcludedWithoutAny() accumulators.ExcludedWithoutAny[time.Duration] {
	return accumulators.NewExcludedWithoutAny(provideExcludedWithoutAny)
}

func ExcludedWithAll() accumulators.ExcludedWithAll[time.Duration] {
	return accumulators.NewExcludedWithAll(provideExcludedWithAll)
}

func ExcludedWithoutAll() accumulators.ExcludedWithoutAll[time.Duration] {
	return accumulators.NewExcludedWithoutAll(provideExcludedWithoutAll)
}

type excludedIf string

func (r excludedIf) validate(name string, value time.Duration) (stop bool, err error) {
	if value != 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameExcludedIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

type excludedUnless string

func (r excludedUnless) validate(name string, value time.Duration) (stop bool, err error) {
	if value != 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameExcludedUnless,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func provideExcludedWithAny(names string, excluded bool) yav.ValidateFunc[time.Duration] {
	if !excluded {
		return yav.Next[time.Duration]
	}

	if validateFunc, ok := excludedWithAnyFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyFuncs, names, excludedWithAny(names))
}

func excludedWithAny(names string) yav.ValidateFunc[time.Duration] {
	return func(name string, value time.Duration) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithoutAny(names string, excluded bool) yav.ValidateFunc[time.Duration] {
	if !excluded {
		return yav.Next[time.Duration]
	}

	if validateFunc, ok := excludedWithoutAnyFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithoutAnyFuncs, names, excludedWithoutAny(names))
}

func excludedWithoutAny(names string) yav.ValidateFunc[time.Duration] {
	return func(name string, value time.Duration) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithAll(names string, excluded bool) yav.ValidateFunc[time.Duration] {
	if !excluded {
		return yav.Next[time.Duration]
	}

	if validateFunc, ok := excludedWithAllFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAllFuncs, names, excludedWithAll(names))
}

func excludedWithAll(names string) yav.ValidateFunc[time.Duration] {
	return func(name string, value time.Duration) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithoutAll(names string, excluded bool) yav.ValidateFunc[time.Duration] {
	if !excluded {
		return yav.Next[time.Duration]
	}

	if validateFunc, ok := excludedWithoutAllFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithoutAllFuncs, names, excludedWithoutAll(names))
}

func excludedWithoutAll(names string) yav.ValidateFunc[time.Duration] {
	return func(name string, value time.Duration) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
